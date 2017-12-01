package main

import (
	"gimcloud/gim_common/gproto"
	"gimcloud/gim_common/libnet"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"time"
)

// echo server demo
func main() {
	flag.Parse()
	client, err := libnet.NewClient("127.0.0.1:9999", 0)
	checkError(err)

	client.Start(func(session *libnet.Session) {
		// handle panic
		defer func() {
			if r := recover(); r != nil {
				glog.Warningf("session start recover: %v", r)
			}
		}()

		glog.Infof("connect to server success id = %v", session.Id())

		// handle incoming packet
		session.OnPacket(func(session *libnet.Session, iMessage *gproto.Message) {
			glog.Infof("id=%v, recv.data=%v", session.Id(), (iMessage.Body.(*gproto.GProtoEchoResponse)).Data)
		})

		// handle session close
		session.OnClose(func(session *libnet.Session, reason error) {
			glog.Infof("close id=%v, reason=%v", session.Id(), reason)
		})

		session.Start()
	})

	// loop input message
	for {
		var input string
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			// do nothing
		} else {
			if input == "exit" {
				break
			}
		}

		gProtoEchoRequest := &gproto.GProtoEchoRequest{
			Data: input,
		}
		gprotoEchoRequestMessage := gproto.NewMessage(1, gproto.MSG_ECHO_REQUEST, gProtoEchoRequest)
		client.SendMessage(gprotoEchoRequestMessage)
	}

	client.Close(nil)
	time.Sleep(time.Second * 2) // wait client to exit

	glog.Infof("client exit")
}

func checkError(err error) {
	if err != nil {
		glog.Errorf("Fatal error: %v", err.Error())
		os.Exit(1)
	}
}
