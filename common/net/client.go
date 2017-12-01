package libnet

import (
	"goim/common/goproto"
	"net"
	"sync/atomic"
	"time"
)

var clientSessionId uint64 = 0

type Client struct {
	session *Session
}

// create a Client Instance
func NewClient(address string, timeout time.Duration) (*Client, error) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}

	session := NewSession(
		genSessionId(), nil, NewNetConn(conn, DefaultReadTimeout, DefaultWriteTimeout), DefaultSendChanSize)

	return &Client{
		session: session,
	}, nil
}

// start client
func (client *Client) Start(callback func(*Session)) {
	if callback != nil {
		callback(client.session)
	}
}

// get client session
func (client *Client) Session() *Session {
	return client.session
}

// send binary
func (client *Client) Send(packet []byte) error {
	return client.session.Send(packet)
}

// send message
func (client *Client) SendMessage(iMessage *goproto.Message) error {
	return client.session.SendMessage(iMessage)
}

// close client
func (client *Client) Close(err error) {
	client.session.Close(err)
}

// generate a new session id
func genSessionId() uint64 {
	return atomic.AddUint64(&clientSessionId, 1)
}
