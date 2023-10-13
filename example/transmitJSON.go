package main

import (
	"encoding/json"
	"fmt"
	"github.com/lim-yoona/tcpack"
	"net"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	address, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8099")
	address2, _ := net.ResolveTCPAddr("tcp4", ":8099")

	//server
	listener, _ := net.ListenTCP("tcp4", address2)

	msgChan := make(chan *tcpack.Imessage, 10)

	go func() {
		for {
			tcpConnServer, _ := listener.AcceptTCP()
			mpServer := tcpack.NewMsgPack(8, tcpConnServer)
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

	mpClient := tcpack.NewMsgPack(8, tcpConnClient)

	jack := &Person{
		Name: "jack",
		Age:  20,
	}
	data, _ := json.Marshal(jack)
	fmt.Println("send JSON:", string(data))
	msgSend := tcpack.NewMessage(0, uint32(len(data)), data)
	_, _ = mpClient.Pack(msgSend)

	msgRevOut := <-msgChan

	var recPerson Person
	json.Unmarshal((*msgRevOut).GetMsgData(), &recPerson)
	fmt.Println("receive JSON", string((*msgRevOut).GetMsgData()))
	fmt.Println("receive struct:", recPerson)
}
