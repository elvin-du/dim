package libnet

import (
	"fmt"
	"net"
	"time"
)

// 对net.Conn做了封装
type NetConn struct {
	conn         net.Conn      // 底层网络连接
	readTimeout  time.Duration // socket read超时, 0没有超时
	writeTimeout time.Duration // socket write超时, 0没有超时
}

// 检测错误是否是网络超时, 如果超时返回true
func isNetConnTimeout(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}

// Version1:
// 检测网络错误(忽略超时, 超时不认为是错误)
// func isNetConnError(err error) bool {
//	return err != nil && !isNetConnTimeout(err)
// }

// Version2:
// Date: 2015-01-06
// 检测网络错误(包含超时, 超时也认为是错误)
func isNetConnError(err error) bool {
	return err != nil
}

// create a NetConn Instance
func NewNetConn(conn net.Conn, readTimeout time.Duration, writeTimeout time.Duration) *NetConn {
	return &NetConn{
		conn:         conn,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}

// get net.Conn
func (nc *NetConn) Conn() net.Conn {
	return nc.conn
}

// set read timeout
func (nc *NetConn) SetReadTimeout(readTimeout time.Duration) {
	nc.readTimeout = readTimeout
}

// get read timeout
func (nc *NetConn) ReadTimeout() time.Duration {
	return nc.readTimeout
}

// set write timeout
func (nc *NetConn) SetWriteTimeout(writeTimeout time.Duration) {
	nc.writeTimeout = writeTimeout
}

// get write timeout
func (nc *NetConn) WriteTimeout() time.Duration {
	return nc.writeTimeout
}

// get local addr
func (nc *NetConn) LocalAddr() net.Addr {
	return nc.conn.LocalAddr()
}

// get remote addr
func (nc *NetConn) RemoteAddr() net.Addr {
	return nc.conn.RemoteAddr()
}

// format connection
func (nc *NetConn) String() string {
	return fmt.Sprintf("%s => %s", nc.LocalAddr(), nc.RemoteAddr())
}

// Write数据, 可以使用isNetConnTimeout(err error)来检测返回的错误是否是网络超时
func (nc *NetConn) Write(b []byte) (n int, err error) {
	if nc.writeTimeout > 0 {
		nc.conn.SetWriteDeadline(time.Now().Add(nc.writeTimeout))
	} else {
		nc.conn.SetWriteDeadline(time.Time{})
	}

	return nc.conn.Write(b)
}

// Read数据, 可以使用isNetConnTimeout(err error)来检测返回的错误是否是网络超时
func (nc *NetConn) Read(b []byte) (n int, err error) {
	if nc.readTimeout > 0 {
		nc.conn.SetReadDeadline(time.Now().Add(nc.readTimeout))
	} else {
		nc.conn.SetReadDeadline(time.Time{})
	}

	// Note:
	// 另外一种读取数据的方式
	//return io.ReadFull(nc.conn, b)

	return nc.conn.Read(b)

}

// close connection
func (nc *NetConn) Close() error {
	return nc.conn.Close()
}
