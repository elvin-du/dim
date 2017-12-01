package libnet

import (
	"encoding/binary"
	"fmt"
	"goim/common/goproto"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/glog"
)

// 用户Session
// 1. 会启动read & write 2个goroutines
// 2. read goroutine负责从网络读取数据, 解包之后调用下面函数来处理网络包:
//    packetCallback func(*Session, []byte)
// 3. write goroutine负责写入数据, 是一个有buf的chan，可以异步的写入数据
// 4. 任何read错误(不包含timeout)和包解析错误, 都会导致read goroutine退出，并同时关闭Session
// 5. 任何write错误(包含timeout)，都会导致write goroutine退出，并同时关闭Session
// 6. 我们可以通过调用Close方法来主动关闭Session, Session一旦关闭, 会自动从Server关联的SessionManager中删除
//

// session close flag
const (
	sessionFlagClosed = iota
	sessionFlagOpen
)

// read fsm flag
const (
	readFlagStartHead = iota
	readFlagContinueHead
	readFlagStartPacket
	readFlagContinuePacket
)

// 由于我们每个解析出来的packet都会分配一个新的buffer，这个参数的意义是将buffer
// 的cap扩大, 这样我们可以在buffer后面append N个bytes的自定义数据而不需要重新分配
// 一个新的buffer.
const (
	defaultPacketEnlargeSize = 4
)

type PacketWithCallback struct {
	Packet []byte      // 实际发送内容
	Data   interface{} // 自定义数据
}

// Session
type Session struct {
	// manager session
	id     uint64  // session id
	server *Server // session owner

	// about network
	netConn *NetConn

	// about send & recv
	sendPacketChan  chan []byte           // 直接发送
	sendMessageChan chan *goproto.Message // 直接发送

	// about session close
	closeChan   chan int // close write loop
	closeWait   *sync.WaitGroup
	closeFlag   int32
	closeReason error

	// callback function
	sendPacketCallback func(*Session, bool, []byte, interface{}) // 发送Packet的回调
	packetCallback     func(*Session, *goproto.Message)          // 接收Packet的回调
	closeCallback      func(*Session, error)

	// 保存自定义的Session数据
	State   interface{}
	StateId int
}

// create a server Session Instance
func NewSession(id uint64, server *Server, netConn *NetConn, sendChanSize uint) *Session {
	return &Session{
		id:                 id,
		server:             server,
		netConn:            netConn,
		sendPacketChan:     make(chan []byte, sendChanSize),
		sendMessageChan:    make(chan *goproto.Message, sendChanSize),
		closeChan:          make(chan int),
		closeWait:          new(sync.WaitGroup),
		closeFlag:          sessionFlagClosed,
		sendPacketCallback: nil,
		packetCallback:     nil,
		closeCallback:      nil,
		State:              nil,
		StateId:            -1,
	}
}

// 启动一个Session的read & write goroutines.
//
// 注意:
// Session在使用之前, 必须先Start
func (session *Session) Start() {
	if atomic.CompareAndSwapInt32(&session.closeFlag, sessionFlagClosed, sessionFlagOpen) {
		session.closeWait.Add(1)
		go session.writeLoop()

		session.closeWait.Add(1)
		go session.readLoop()

		if session.server == nil {
			glog.Infof("session start id=%v", session.id)
		} else {
			glog.Infof("session start server=%v id=%v", session.server.address, session.id)
		}
	} else {
		panic(SessionDuplicateStartError)
	}
}

// 循环的read并解析网络数据
//
// 注意:
// 1. 每一个收到的Packet都会分配一个新的buffer
// 2. 如果数据读取或者解包失败, 会直接关闭Session
func (session *Session) readLoop() {
	var (
		err error
	)

	defer func() {
		// handle panic
		if r := recover(); r != nil {
			glog.Infof("session readLoop recover: %v, DEBUG.STACK=%v", r, string(debug.Stack()))
		}

		session.closeWait.Done()
		session.Close(err)
	}()

	var (
		read_flag  uint = readFlagStartHead
		headBuf    []byte
		packet     []byte
		headLen    uint = 0
		packetLen  uint = 0
		packetSize uint = 0

		// Message Head
		magicWord     uint32
		version       uint16
		command       uint16
		payloadLength uint32

		n int
	)

	headBuf = make([]byte, 12, 12) // Fix GProto Head Len

	for {
		// 1. start read head
		if read_flag == readFlagStartHead {
			n, err = session.netConn.Read(headBuf[0:])
			if isNetConnError(err) {
				return
			}

			headLen = uint(n)

			if headLen == 12 {
				read_flag = readFlagStartPacket
			} else {
				read_flag = readFlagContinueHead
			}
		}

		// 2. continue read head
		if read_flag == readFlagContinueHead {
			n, err = session.netConn.Read(headBuf[headLen:])
			if isNetConnError(err) {
				return
			}

			headLen += uint(n)

			if headLen == 12 {
				read_flag = readFlagStartPacket
			} else {
				read_flag = readFlagContinueHead
			}
		}

		// 3. start read packet
		if read_flag == readFlagStartPacket {
			magicWord = binary.BigEndian.Uint32(headBuf[0:4])
			version = binary.BigEndian.Uint16(headBuf[4:6])
			command = binary.BigEndian.Uint16(headBuf[6:8])
			payloadLength = binary.BigEndian.Uint32(headBuf[8:]) // PacketSize

			packetSize = uint(payloadLength)

			if packetSize == 0 {
				// Fix Bug:
				// 针对packetSize == 0 特殊处理
				read_flag = readFlagStartHead
				headLen = 0
				packetLen = 0
				packetSize = 0

				// handle packet
				if session.packetCallback != nil {
					iMessage := goproto.NewEmptyMessage()

					iMessage.MagicWord = magicWord
					iMessage.Version = version
					iMessage.Command = command
					iMessage.PayloadLength = payloadLength

					if iMessage.MagicWord != goproto.MAGIC_WORLD {
						return
					}

					if err := iMessage.Decode(nil); err != nil {
						return
					}

					session.packetCallback(session, iMessage)
				}
			} else {
				// check packet size
				if packetSize > DefaultMaxPacketSize {
					err = PacketHeadLenError
					return
				}

				// allociate a new buf for this packet
				//
				// Note:
				// enlarge the packet buffer, give a chance to append data under this buffer,
				// and do not need to allociate a new memory.
				packet = make([]byte, packetSize, packetSize+defaultPacketEnlargeSize)

				n, err = session.netConn.Read(packet[0:])
				if isNetConnError(err) {
					return
				}

				packetLen = uint(n)

				if packetLen == packetSize {
					read_flag = readFlagStartHead
					headLen = 0
					packetLen = 0
					packetSize = 0

					// handle packet
					if session.packetCallback != nil {
						iMessage := goproto.NewEmptyMessage()

						iMessage.MagicWord = magicWord
						iMessage.Version = version
						iMessage.Command = command
						iMessage.PayloadLength = payloadLength

						if iMessage.MagicWord != goproto.MAGIC_WORLD {
							return
						}

						if err := iMessage.Decode(packet); err != nil {
							return
						}

						session.packetCallback(session, iMessage)
					}
				} else {
					read_flag = readFlagContinuePacket
				}
			}
		}

		// 4. continue read packet
		if read_flag == readFlagContinuePacket {
			n, err = session.netConn.Read(packet[packetLen:])
			if isNetConnError(err) {
				return
			}

			packetLen += uint(n)

			if packetLen == packetSize {
				read_flag = readFlagStartHead
				headLen = 0
				packetLen = 0
				packetSize = 0

				// handle packet
				if session.packetCallback != nil {
					iMessage := goproto.NewEmptyMessage()

					iMessage.MagicWord = magicWord
					iMessage.Version = version
					iMessage.Command = command
					iMessage.PayloadLength = payloadLength

					if iMessage.MagicWord != goproto.MAGIC_WORLD {
						return
					}

					if err := iMessage.Decode(packet); err != nil {
						return
					}

					session.packetCallback(session, iMessage)
				}
			} else {
				read_flag = readFlagContinuePacket
			}
		}
	}
}

// 循环的写入数据
func (session *Session) writeLoop() {
	var (
		err error
	)

	defer func() {
		// handle panic
		if r := recover(); r != nil {
			glog.Infof("session writeLoop recover: %v, DEBUG.STACK=%v", r, string(debug.Stack()))
		}

		session.closeWait.Done()
		session.Close(err)
	}()

	var (
		packet  []byte
		message *goproto.Message
	)

	for {
		select {
		case packet = <-session.sendPacketChan:
			if _, err = session.netConn.Write(packet); err != nil {
				return
			}

			// 遇到任何写入错误, 直接关闭socket
			//
			// 注意:
			// 包含timeout错误
			if err != nil {
				return
			}
		case message = <-session.sendMessageChan:
			bodyBuf, err := message.Encode()
			if err != nil {
				return
			}

			if bodyBuf == nil {
				message.PayloadLength = 0
			} else {
				message.PayloadLength = uint32(len(bodyBuf)) // 写入Payload Length
			}

			headBuf, err := message.EncodeHead()
			if err != nil {
				return
			}

			if _, err = session.netConn.Write(headBuf); err != nil {
				return
			}

			// Fix Bug:
			// 有内容从才Send
			if bodyBuf != nil || len(bodyBuf) > 0 {
				if _, err = session.netConn.Write(bodyBuf); err != nil {
					return
				}
			}

			// 遇到任何写入错误, 直接关闭socket
			//
			// 注意:
			// 包含timeout错误
			if err != nil {
				return
			} else {
				glog.Infof("send message successfully %X,%+v", message.Command, string(bodyBuf))
			}
		case <-session.closeChan:
			// Note:
			// if close session without this goroutine
			// deadlock will happen when session close itself.
			go func() {
				// 1. close socket, let read loop exit
				session.netConn.Close()

				// 3. wait for read & write loop exit
				session.closeWait.Wait()

				// 4. remove session from its owner(*Server)
				if session.server != nil {
					session.server.delSession(session)
				}

				// 5. invoke session close callback
				if session.closeCallback != nil {
					session.closeCallback(session, session.closeReason)
				}

				if session.server == nil {
					glog.Infof("session close server id=%v", session.id)
				} else {
					glog.Infof("session close server=%v id=%v", session.server.address, session.id)
				}
			}()
			return
		}
	}
}

// get session id.
func (session *Session) Id() uint64 {
	return session.id
}

// get NetConn
func (session *Session) NetConn() *NetConn {
	return session.netConn
}

// get local addr
func (session *Session) LocalAddr() net.Addr {
	return session.netConn.LocalAddr()
}

// get remote addr
func (session *Session) RemoteAddr() net.Addr {
	return session.netConn.RemoteAddr()
}

// format session
func (session *Session) String() string {
	return fmt.Sprintf("id=%v, conn=%v", session.id, session.netConn)
}

// get session owner
func (session *Session) Server() *Server {
	return session.server
}

// check session is closed or not.
func (session *Session) IsClosed() bool {
	return atomic.LoadInt32(&session.closeFlag) == sessionFlagClosed
}

// 注册packet callback.
func (session *Session) OnPacket(callback func(*Session, *goproto.Message)) {
	session.packetCallback = callback
}

// 注册session close callback.
func (session *Session) OnClose(callback func(*Session, error)) {
	session.closeCallback = callback
}

// 延时close session
//
// 1. 这对于一些需要发送最终的LOGOUT或者KICKOFF包再关闭的Session的场景有用.
// 2. 如果调用Close方法关闭Session会导致session立刻关闭(atomic.CompareAndSwapInt32),
//    在这之前调用的Send方法发送数据不保证一定能发送到客户端.
// 3. 这个方法会立刻返回, Session关闭sleep的逻辑在goroutine中, 不会block住调用进程
//
//    例如下面代码:
//    session.Send(data)
//    session.Close(reason)
//
//    data可能没有发送session就关闭了.
//
//    建议使用下面代码:
//    session.Send(data)
//    session.DelayClose(reason, 2000)
//
func (session *Session) DelayClose(reason error, millisecond time.Duration) {
	if atomic.CompareAndSwapInt32(&session.closeFlag, sessionFlagOpen, sessionFlagClosed) {
		session.closeReason = reason

		// Note:
		// if close session without this goroutine
		// deadlock will happen when session close itself.
		go func() {
			time.Sleep(time.Millisecond * millisecond)

			// 1. close socket, let read loop exit
			session.netConn.Close()

			// 2. notify write loop exit
			close(session.closeChan)

			// 3. wait for read & write loop exit
			// session.closeWait.Wait()

			// 4. remove session from its owner(*Server)
			// if session.server != nil {
			// 	session.server.delSession(session)
			// }

			// 5. invoke session close callback
			// if session.closeCallback != nil {
			//	session.closeCallback(session, reason)
			//}

			//if session.server == nil {
			//	glog.Infof("session close server id=%v", session.id)
			//} else {
			//	glog.Infof("session close server=%v id=%v", session.server.address, session.id)
			// }
		}()
	}
}

// async close session and remove it from owner server if owner is not nil
//
// Logic:
// 1. close read goroutine
// 2. close write goroutine
// 3. wait read & wait exit
// 4. remove session from its owner(*Server)
// 5. invoke session close callback
func (session *Session) Close(reason error) {
	if atomic.CompareAndSwapInt32(&session.closeFlag, sessionFlagOpen, sessionFlagClosed) {
		session.closeReason = reason
		close(session.closeChan)

		// Note:
		// if close session without this goroutine
		// deadlock will happen when session close itself.
		// go func() {
		// 	// 1. close socket, let read loop exit
		// 	session.netConn.Close()

		// 	// 2. notify write loop exit
		// 	close(session.closeChan)

		// 	// 3. wait for read & write loop exit
		// 	session.closeWait.Wait()

		// 	// 4. remove session from its owner(*Server)
		// 	if session.server != nil {
		// 		session.server.delSession(session)
		// 	}

		// 	// 5. invoke session close callback
		// 	if session.closeCallback != nil {
		// 		session.closeCallback(session, reason)
		// 	}

		// 	if session.server == nil {
		// 		glog.Infof("session close server id=%v", session.id)
		// 	} else {
		// 		glog.Infof("session close server=%v id=%v", session.server.address, session.id)
		// 	}
		// }()
	}
}

// send binary
func (session *Session) Send(packet []byte) error {
	if session.IsClosed() {
		return SendToClosedError
	}

	select {
	case session.sendPacketChan <- packet:
		return nil
	default:
		session.Close(BlockingError)
		return BlockingError
	}
}

// send binary
func (session *Session) SendMessage(iMessage *goproto.Message) error {
	if session.IsClosed() {
		return SendToClosedError
	}

	select {
	case session.sendMessageChan <- iMessage:
		return nil
	default:
		session.Close(BlockingError)
		return BlockingError
	}
}
