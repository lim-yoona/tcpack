package main

import (
	"fmt"
	"github.com/lim-yoona/msgpack"
	"net"
	"time"
)

func main() {
	address, err := net.ResolveTCPAddr("tcp4", ":8099")
	if err != nil {
		fmt.Println("create address failed")
	}
	fmt.Println(*address)
	listener, err := net.ListenTCP("tcp4", address)
	if err != nil {
		fmt.Println("create listener failed")
	}
	for {
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("return conn failed")
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
				time.Sleep(time.Second * 5)
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
	}
}
