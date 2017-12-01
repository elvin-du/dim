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

var (
	newConnectionInterval = time.Millisecond * 5
	sendPacketInterval    = time.Second
	defaultPacket         = []byte("packet:abcdefghijklmnopqrstuvwxyz")
)

// stress client
func main() {
	flag.Parse()
	for i := 0; i <= 1000; i++ {
		time.Sleep(newConnectionInterval)
		launchClient("127.0.0.1:8888", 0)
	}

	for {
		time.Sleep(time.Second * 10)
	}

}

func launchClient(address string, timeout time.Duration) {
	client, err := libnet.NewClient(address, timeout)
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
			gProtoEchoResponse := (iMessage.Body).(*gproto.GProtoEchoResponse)
			// glog.Infof("client id=%v say %v", session.Id(), gProtoEchoResponse.Data)
			if gProtoEchoResponse.Data != "Test_Data" {
				panic(fmt.Errorf("Invalid Logic"))
			}

			// time.Sleep(sendPacketInterval)
			gProtoEchoRequest := &gproto.GProtoEchoRequest{
				Data: "Test_Data",
			}
			gprotoEchoRequestMessage := gproto.NewMessage(1, gproto.MSG_ECHO_REQUEST, gProtoEchoRequest)
			client.SendMessage(gprotoEchoRequestMessage)
		})

		// handle session close
		session.OnClose(func(session *libnet.Session, reason error) {
			glog.Infof("close id=%v, reason=%v", session.Id(), reason)
		})

		session.Start()

		gProtoEchoRequest := &gproto.GProtoEchoRequest{
			Data: "Test_Data",
		}
		gprotoEchoRequestMessage := gproto.NewMessage(1, gproto.MSG_ECHO_REQUEST, gProtoEchoRequest)
		client.SendMessage(gprotoEchoRequestMessage)
	})
}

func checkError(err error) {
	if err != nil {
		glog.Errorf("Fatal error: %v", err.Error())
		os.Exit(1)
	}
}
