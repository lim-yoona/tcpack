# tcpack

[![Go Reference](https://pkg.go.dev/badge/github.com/lim-yoona/tcpack.svg)](https://pkg.go.dev/github.com/lim-yoona/tcpack)
![GitHub](https://img.shields.io/github/license/lim-yoona/tcpack)
[![Go Report](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/lim-yoona/tcpack)
![GitHub release (with filter)](https://img.shields.io/github/v/release/lim-yoona/tcpack)

[English](README.md) | 简体中文  

[tcpack](https://pkg.go.dev/github.com/lim-yoona/tcpack) 是一个基于 TCP 的应用层协议，用于在 [go](https://go.dev/) 程序中打包和解包字节流。  

## tcpack做了什么？  

众所周知，TCP 是面向字节流的传输层协议，其数据传输没有明确的边界，因此应用层读取的数据可能包含多个请求而导致无法处理业务。  

[tcpack](https://pkg.go.dev/github.com/lim-yoona/tcpack) 就是为了解决这个问题，将请求数据封装成消息，发送时打包，接收时解包。  

## tcpack中有什么?  

`tcpack` 提供了一个支持 Pack 和 Unpack 的打包器。  

## 安装指南

1. 为安装 tcpack 包, 首先你需要安装 [Go](https://go.dev/doc/install) , 然后你可以使用下面的命令将 `tcpack` 作为你Go程序的依赖。    

```sh
go get -u github.com/lim-yoona/tcpack
```

2. 将 tcpack 导入到代码中：  

```go
import "github.com/lim-yoona/tcpack"
```

## 使用

```go
package main

import "github.com/lim-yoona/tcpack"

func main() {
    // 创建一个打包器
    mp := tcpack.NewMsgPack(8, tcpConn)

    // 打包一个消息
    msg := tcpack.NewMessage(0, uint32(len([]byte(data))), []byte(data))
    msgByte, err := mp.Pack(msg)
    num, err := tcpConn.Write(msgByte)

    // 解包一个消息
    msg, err := mp.Unpack()
}
```

### 支持JSON

```go
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 创建一个打包器
mp := tcpack.NewMsgPack(8, tcpConn)

// data JSON Marshal
data := &Person{
	Name: "jack",
	Age:  20,
}
dataJSON, _ := json.Marshal(data)

// 打包一个消息
msgSend := tcpack.NewMessage(0, uint32(len(dataJSON)), dataJSON)
msgSendByte, _ := mpClient.Pack(msgSend)
num, err := tcpConn.Write(msgSendByte)

// 解包一个消息
msgRsv, err := mp.Unpack()

// JSON UnMarshal
var dataRsv Person
json.Unmarshal(msgRsv.GetMsgData(), &dataRsv)
```

## 示例

这有一些 [示例](https://github.com/lim-yoona/tcpack/tree/main/example)。  

