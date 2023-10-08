package msgpack

import (
	"fmt"
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMsgPack_Pack(t *testing.T) {
	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8099")
	address2, _ := net.ResolveTCPAddr("tcp4", ":8099")

	//server
	listener, _ := net.ListenTCP("tcp4", address2)

	msgChan := make(chan *Imessage, 10)

	go func() {
		for {
			tcpConnServer, _ := listener.AcceptTCP()
			mpServer := NewMsgPack(8, tcpConnServer)
			for {
				msgRev, _ := mpServer.Unpack()
				msgChan <- &msgRev
			}
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

	// Continuously sending short messages
	data1 := "a"
	data2 := "b"
	data3 := "c"
	msgSend1 := NewMessage(0, uint32(len([]byte(data1))), []byte(data1))
	msgSendByte1, _ := mpClient.Pack(msgSend1)
	msgSend2 := NewMessage(0, uint32(len([]byte(data2))), []byte(data2))
	msgSendByte2, _ := mpClient.Pack(msgSend2)
	msgSend3 := NewMessage(0, uint32(len([]byte(data3))), []byte(data3))
	msgSendByte3, _ := mpClient.Pack(msgSend3)
	_, _ = tcpConnClient.Write(msgSendByte1)
	_, _ = tcpConnClient.Write(msgSendByte2)
	_, _ = tcpConnClient.Write(msgSendByte3)
	msgRevOut1 := <-msgChan
	msgRevOut2 := <-msgChan
	msgRevOut3 := <-msgChan
	Convey("Continuously sending short messages", t, func() {
		So(msgSend1, ShouldEqual, *msgRevOut1)
	})
	Convey("Continuously sending short messages", t, func() {
		So(msgSend2, ShouldEqual, *msgRevOut2)
	})
	Convey("Continuously sending short messages", t, func() {
		So(msgSend3, ShouldEqual, *msgRevOut3)
	})

	// Send a message that exceeds the size of the socket buffer
	data4 := make([]byte, 65600)
	msgSend4 := NewMessage(0, uint32(len(data4)), data4)
	msgSendByte4, _ := mpClient.Pack(msgSend4)
	_, _ = tcpConnClient.Write(msgSendByte4)
	msgRevOut4 := <-msgChan
	Convey("Send a message that exceeds the size of the socket buffer", t, func() {
		So(msgSend4, ShouldEqual, *msgRevOut4)
		So(uint32(65600), ShouldEqual, (*msgRevOut4).GetDataLen())
	})
}

func TestMsgPack_Unpack(t *testing.T) {
	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9099")
	address2, _ := net.ResolveTCPAddr("tcp4", ":9099")

	//server
	listener, _ := net.ListenTCP("tcp4", address2)

	msgChan := make(chan *Imessage, 10)

	go func() {
		for {
			tcpConnServer, _ := listener.AcceptTCP()
			mpServer := NewMsgPack(8, tcpConnServer)
			for {
				msgRev, _ := mpServer.Unpack()
				msgChan <- &msgRev
			}
		}
	}()

	// client
	tcpConnClient, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		fmt.Println("create tcpConn failed")
	}

	mpClient := NewMsgPack(8, tcpConnClient)
	data := ""
	msgSend := NewMessage(0, uint32(len([]byte(data))), []byte(data))
	msgSendByte, _ := mpClient.Pack(msgSend)
	_, _ = tcpConnClient.Write(msgSendByte)

	msgRevOut := <-msgChan

	Convey("send an empty msg", t, func() {
		So(msgSend, ShouldEqual, *msgRevOut)
	})
}

func TestMsgPack_GetHeadLen(t *testing.T) {
	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8099")
	tcpConn, _ := net.DialTCP("tcp4", nil, address)
	mp := NewMsgPack(9, tcpConn)
	Convey("get head length", t, func() {
		So(mp.GetHeadLen(), ShouldEqual, 9)
	})
}
