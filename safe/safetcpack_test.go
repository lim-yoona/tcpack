package safetcpack

import (
	"fmt"
	"github.com/lim-yoona/tcpack"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"sync"
	"testing"
	"time"
)

func TestNewSafeMsgPack(t *testing.T) {

	address2, _ := net.ResolveTCPAddr("tcp4", ":8099")
	listener, _ := net.ListenTCP("tcp4", address2)
	go func() {
		for {
			_, _ = listener.AcceptTCP()
		}
	}()

	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8099")
	tcpConn, _ := net.DialTCP("tcp4", nil, address)
	smp := NewSafeMsgPack(8, tcpConn)
	smp2 := NewSafeMsgPack(8, tcpConn)
	smp3 := NewSafeMsgPack(6, tcpConn)
	Convey("whether only onne packager for the same connection", t, func() {
		So(smp, ShouldEqual, smp2)
		So(8, ShouldEqual, smp3.GetHeadLen())
		So(smp, ShouldEqual, smp3)
		So(smp2, ShouldEqual, smp3)
	})
}

func TestSafeMsgPack_Pack(t *testing.T) {
	// Mainly testing whether messages can be
	// safely sent and received in concurrent scenarios.
	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9001")
	address2, _ := net.ResolveTCPAddr("tcp4", ":9001")

	//server
	listener, _ := net.ListenTCP("tcp4", address2)

	connChan := make(chan net.Conn, 5)

	revNum1 := 0
	revNum2 := 0
	revNum3 := 0
	go func() {
		go func() {
			tcpConnServer2 := <-connChan
			mpServer2 := NewSafeMsgPack(8, tcpConnServer2)
			for {
				_, _ = mpServer2.Unpack()
				revNum1++
			}
		}()
		go func() {
			tcpConnServer2 := <-connChan
			mpServer2 := NewSafeMsgPack(8, tcpConnServer2)
			for {
				_, _ = mpServer2.Unpack()
				revNum2++
			}
		}()
		for {
			tcpConnServer, _ := listener.AcceptTCP()
			connChan <- tcpConnServer
			connChan <- tcpConnServer
			mpServer := NewSafeMsgPack(8, tcpConnServer)
			for {
				_, _ = mpServer.Unpack()
				revNum3++
			}
		}
	}()

	// client
	tcpConnClient, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		fmt.Println("create tcpConn failed")
	}

	sendNum1 := 0
	sendNum2 := 0
	sendNum3 := 0
	stopChan := make(chan int, 3)
	sChan := make(chan int)
	go func() {
		for {
			data := []byte("helloworld")
			mpClient := NewSafeMsgPack(8, tcpConnClient)
			msgSend := tcpack.NewMessage(0, uint32(len(data)), data)
			_, _ = mpClient.Pack(msgSend)
			sendNum1++
			select {
			case <-stopChan:
				<-sChan
			default:

			}
		}
	}()
	go func() {
		for {
			data := []byte("1")
			mpClient := NewSafeMsgPack(8, tcpConnClient)
			msgSend := tcpack.NewMessage(0, uint32(len(data)), data)
			_, _ = mpClient.Pack(msgSend)
			sendNum2++
			select {
			case <-stopChan:
				<-sChan
			default:

			}
		}
	}()
	go func() {
		for {
			data := make([]byte, 65600)
			mpClient := NewSafeMsgPack(8, tcpConnClient)
			msgSend := tcpack.NewMessage(0, uint32(len(data)), data)
			_, _ = mpClient.Pack(msgSend)
			sendNum3++
			select {
			case <-stopChan:
				<-sChan
			default:

			}
		}
	}()

	var mu sync.Mutex
	for true {
		mu.Lock()
		if sendNum1+sendNum2+sendNum3 == 50000 {
			stopChan <- 1
			stopChan <- 1
			stopChan <- 1
			break
		}
		mu.Unlock()
	}

	time.Sleep(time.Second * 100)
	fmt.Println("send message number:", sendNum1+sendNum2+sendNum3)
	fmt.Println("receive message number:", revNum1+revNum2+revNum3)
	Convey("correct send and receive message number whether equal", t, func() {
		So(sendNum1+sendNum2+sendNum3, ShouldEqual, revNum1+revNum2+revNum3)
	})
}
