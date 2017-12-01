package libnet

import (
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/golang/glog"
)

// Server会在当前goroutine中不断的Accept新连接, 每Accept一个连接, 会异步(在一个go goroutine中)
// 执行Server.Start中的callback.

// default send chan buffer size for sessions.
const (
	DefaultSendChanSize  = 2048
	DefaultMaxPacketSize = 10 * 1024 * 1024  // 10MB
	DefaultReadTimeout   = time.Second * 600 // 10 minutes
	DefaultWriteTimeout  = time.Second * 30  // 30 seconds
)

// session close flag
const (
	serverFlagStop = iota
	serverFlagOpen
)

// Server.
type Server struct {
	// About network
	address    string
	listenAddr net.Addr
	listener   net.Listener

	// About sessions
	sendChanSize   uint
	sessionManager *SessionManager

	// About server start and stop
	stopChan chan int
	stopFlag int32

	// Put your server state here.
	State interface{}
}

// create a Server Instance
func NewServer(address string) (*Server, error) {

	listenAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		glog.Warningf("net.ResolveTCPAddr address=%v, err=%v", address, err)
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		glog.Warningf("net.ListenTCP address=%v, err=%v", address, err)
		return nil, err
	}

	glog.Infof("new server success address=%v", address)

	return &Server{
		address:        address,
		listenAddr:     listenAddr,
		listener:       listener,
		sendChanSize:   DefaultSendChanSize,
		sessionManager: NewSessionManager(),
		stopChan:       make(chan int),
		stopFlag:       serverFlagStop,
	}, nil
}

// get listen net.Addr
func (server *Server) ListenAddr() net.Addr {
	return server.listenAddr
}

// get net.Listener
func (server *Server) Listener() net.Listener {
	return server.listener
}

// set session send channel buffer size setting
//
// Note:
// new setting will effect on new sessions
func (server *Server) SetSendChanSize(size uint) {
	server.sendChanSize = size
}

// get current session send chan buffer size setting.
func (server *Server) SendChanSize() uint {
	return server.sendChanSize
}

// 开始Accept Connections
//
// 注意:
// 每Accept到一个Connect，会异步的执行callback，也就是在一个go goroutine中执行这个callback.
func (server *Server) Start(callback func(*Session)) {
	if atomic.CompareAndSwapInt32(&server.stopFlag, serverFlagStop, serverFlagOpen) {
		server.acceptLoop(callback)
	} else {
		panic(ServerDuplicateStartError)
	}
}

// stop server
//
// 1. stop acceptloop goroutine
// 2. close all session
// 3. wait all session exit
func (server *Server) Stop() {
	if atomic.CompareAndSwapInt32(&server.stopFlag, serverFlagOpen, serverFlagStop) {
		// if stop server without this goroutine
		// deadlock will happen when server closed by session.
		go func() {
			glog.Infof("start stop server address=%v", server.address)
			// stop acceptloop
			server.listener.Close()

			// wait acceptLoop exit
			<-server.stopChan

			// close all sessions
			server.CloseSessions()
			glog.Infof("server stopped address=%v", server.address)
		}()
	}
}

// 不断的Accept连接, 直到Accept遇到io.EOF, 会导致Server停止.
func (server *Server) acceptLoop(callback func(*Session)) {
	defer func() {
		close(server.stopChan)
		server.Stop()
		glog.Infof("stop server acceptLoop address=%v", server.address)
	}()

	for {
		conn, err := server.listener.Accept()
		if err != nil && err == io.EOF {
			// stop the server
			break
		}

		// 接收到新的connection, 异步执行callback
		go server.startSession(conn, callback)
	}
}

// start a session to present the connection.
func (server *Server) startSession(conn net.Conn, callback func(*Session)) {
	session := NewSession(
		server.sessionManager.GenSessionId(),
		server,
		NewNetConn(conn, DefaultReadTimeout, DefaultWriteTimeout),
		server.sendChanSize,
	)

	// init the session state
	if callback != nil {
		callback(session) // 执行callback
	}

	// session maybe closed in start callback
	if !session.IsClosed() {
		server.setSession(session)
	} else {
		conn.Close()
	}
}

// get session count
func (server *Server) SessionCount() int {
	return server.sessionManager.Count()
}

// set a session
func (server *Server) setSession(session *Session) {
	if atomic.LoadInt32(&server.stopFlag) == serverFlagOpen {
		glog.Infof("set session=%v", session)
		server.sessionManager.SetSession(session)
	}
}

// delete a session
func (server *Server) delSession(session *Session) {
	if atomic.LoadInt32(&server.stopFlag) == serverFlagOpen {
		glog.Infof("del session=%v", session)
		server.sessionManager.DelSession(session)
	}
}

// close all sessions
func (server *Server) CloseSessions() {
	server.sessionManager.CloseSessions()
}
