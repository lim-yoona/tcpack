package msgpack

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
)

func TestMsgPack_Pack(t *testing.T) {
	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8099")
	address2, _ := net.ResolveTCPAddr("tcp4", ":8099")

	//server
	listener, _ := net.ListenTCP("tcp4", address2)

	msgChan := make(chan *Imessage)

	go func() {
		for {
			tcpConnServer, _ := listener.AcceptTCP()
			mpServer := NewMsgPack(8, tcpConnServer)
			msgRev, _ := mpServer.Unpack()
			msgChan <- &msgRev
		}
	}()

	// client
	tcpConnClient, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		fmt.Println("create tcpConn failed")
	}

	mpClient := NewMsgPack(8, tcpConnClient)
	data := "Hello,World"
	msgSend := NewMessage(0, uint32(len([]byte(data))), []byte(data))
	msgSendByte, _ := mpClient.Pack(msgSend)
	_, _ = tcpConnClient.Write(msgSendByte)

	msgRevOut := <-msgChan

	Convey("pack msg to send", t, func() {
		So(msgSend, ShouldEqual, *msgRevOut)
	})
	Convey("get head length", t, func() {
		So(mpClient.GetHeadLen(), ShouldEqual, 8)
	})
}
