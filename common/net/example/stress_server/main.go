package main

import (
	"gimcloud/gim_common/gproto"
	"gimcloud/gim_common/libnet"
	"flag"
	"github.com/golang/glog"
	"os"
	"sync/atomic"
	"time"
)

var globalPacketNum uint64

// echo client demo
func main() {
	flag.Parse()
	server, err := libnet.NewServer(":8888")
	checkError(err)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			glog.Infof("session Count=%v, packet=%v", server.SessionCount(), globalPacketNum)
		}
	}()

	server.Start(func(session *libnet.Session) {
		// handle panic
		defer func() {
			if r := recover(); r != nil {
				glog.Warningf("session start recover: %v", r)
			}
		}()

		glog.Infof("session_count=%v, new cient id = %v", server.SessionCount(), session.Id())

		// handle incoming packet
		session.OnPacket(func(session *libnet.Session, iMessage *gproto.Message) {
			atomic.AddUint64(&globalPacketNum, 1)

			gProtoEchoRequest := (iMessage.Body).(*gproto.GProtoEchoRequest)
			switch gProtoEchoRequest.Data {
			case "bye":
				glog.Infof("client id=%v say bye", session.Id())
				session.Close(nil)
			default:
				// glog.Infof("client id=%v say %v", session.Id(), gProtoEchoRequest.Data)

				gProtoEchoResponse := &gproto.GProtoEchoResponse{
					Code: 0,
					Data: gProtoEchoRequest.Data,
				}
				gprotoEchoResponseMessage := gproto.NewMessage(1, gproto.MSG_ECHO_RESPONSE, gProtoEchoResponse)
				session.SendMessage(gprotoEchoResponseMessage)
			}
		})

		// handle session close
		session.OnClose(func(session *libnet.Session, reason error) {
			glog.Infof("session_count=%v, close cient id = %v, reason=%v", server.SessionCount(), session.Id(), reason)
		})

		session.Start()
	})
}

func checkError(err error) {
	if err != nil {
		glog.Errorf("Fatal error: %v", err.Error())
		os.Exit(1)
	}
}
