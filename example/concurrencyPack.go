package main

import (
	"fmt"
	"github.com/lim-yoona/tcpack"
	safetcpack "github.com/lim-yoona/tcpack/safe"
	"net"
	"sync"
	"time"
)

func main() {
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
	//var mu sync.Mutex
	go func() {
		go func() {
			tcpConnServer2 := <-connChan
			mpServer2 := safetcpack.NewSafeMsgPack(8, tcpConnServer2)
			for {
				msgRev2, _ := mpServer2.Unpack()
				fmt.Println("tcpConnServer2 receive:", msgRev2)
				fmt.Println("tcpConnServer2 receive:", string(msgRev2.GetMsgData()))
				revNum1++
			}
		}()
		go func() {
			tcpConnServer2 := <-connChan
			mpServer2 := safetcpack.NewSafeMsgPack(8, tcpConnServer2)
			for {
				msgRev2, _ := mpServer2.Unpack()
				//msgChan <- &msgRev2
				fmt.Println("tcpConnServer2 receive:", msgRev2)
				fmt.Println("tcpConnServer2 receive:", string(msgRev2.GetMsgData()))
				revNum2++
			}
		}()
		for {
			tcpConnServer, _ := listener.AcceptTCP()
			connChan <- tcpConnServer
			connChan <- tcpConnServer
			mpServer := safetcpack.NewSafeMsgPack(8, tcpConnServer)
			for {
				msgRev, _ := mpServer.Unpack()
				//msgChan <- &msgRev
				fmt.Println("tcpConnServer1 receive:", msgRev)
				fmt.Println("tcpConnServer1 receive:", string(msgRev.GetMsgData()))
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
			data := []byte("helloworld!")
			mpClient := safetcpack.NewSafeMsgPack(8, tcpConnClient)
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
			mpClient := safetcpack.NewSafeMsgPack(8, tcpConnClient)
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
			data := []byte("qwertyuiopasdfghjklzxcvbnm1234567890")
			mpClient := safetcpack.NewSafeMsgPack(8, tcpConnClient)
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
	fmt.Println("sendNum:", sendNum1+sendNum2+sendNum3)
	fmt.Println("revNum:", revNum1+revNum2+revNum3)
	fmt.Println("succeed!")
}
