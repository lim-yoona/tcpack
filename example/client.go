package main

import (
	"fmt"
	"github.com/lim-yoona/msgpack"
	"net"
)

func main() {
	address, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8099")
	if err != nil {
		fmt.Println("create address failed")
	}
	fmt.Println(*address)
	tcpConn, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		fmt.Println("create tcpconn failed", err)
	}
	mp := msgpack.NewMsgPack(8, tcpConn)
	go func() {
		for {
			msg, err := mp.Unpack()
			if err != nil {
				fmt.Println("read msg failed:", err)
			} else {
				fmt.Printf("读取到了%d个字符\n", msg.GetDataLen())
				fmt.Println("内容为:", (string(msg.GetMsgData())))
			}
		}
	}()
	go func() {
		for {
			var data string
			fmt.Scanln(&data)
			msg := msgpack.NewMessage(0, uint32(len([]byte(data))), []byte(data))
			msgByte, err := mp.Pack(msg)
			if err != nil {
				fmt.Println("msg pack failed:", err)
			}
			num, err := tcpConn.Write(msgByte)
			if err != nil {
				fmt.Println("msg send failed:", err)
			}
			fmt.Printf("发送了%d个字符\n", num)
			fmt.Printf("内容为:%s(%s)\n", []byte(data), data)
		}
	}()
	select {}
}
