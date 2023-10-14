# tcpack

[![Go Reference](https://pkg.go.dev/badge/github.com/lim-yoona/tcpack.svg)](https://pkg.go.dev/github.com/lim-yoona/tcpack)
![GitHub](https://img.shields.io/github/license/lim-yoona/tcpack)
[![Go Report](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/lim-yoona/tcpack)
![GitHub release (with filter)](https://img.shields.io/github/v/release/lim-yoona/tcpack)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)

[English](README.md) | 简体中文  

[tcpack](https://pkg.go.dev/github.com/lim-yoona/tcpack) 是一个基于 TCP 的应用层协议，用于在 [go](https://go.dev/) 程序中打包和解包字节流。  

## tcpack做了什么？  

众所周知，TCP 是面向字节流的传输层协议，其数据传输没有明确的边界，因此应用层读取的数据可能包含多个请求而导致无法处理业务。  

[tcpack](https://pkg.go.dev/github.com/lim-yoona/tcpack) 就是为了解决这个问题，将请求数据封装成消息，发送时打包，接收时解包。  

*注意： 在同一个连接上使用打包器并发读写消息是不安全的，不要去这样做，会带来不可预知的后果！*  

如果你想要使用在同一个TCP连接上的多个打包器并发地收发消息，请使用 [safetcpack](https://github.com/lim-yoona/tcpack/tree/main/safe/README-CN.md)。  

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

    // 打包一个消息并发送
    msg := tcpack.NewMessage(0, uint32(len([]byte(data))), []byte(data))
    num, err := mp.Pack(msg)

    // 解包一个消息并接收
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

// 打包一个消息并发送
msgSend := tcpack.NewMessage(0, uint32(len(dataJSON)), dataJSON)
num, err := mp.Pack(msgSend)

// 解包一个消息并接收
msgRsv, err := mp.Unpack()

// JSON UnMarshal
var dataRsv Person
json.Unmarshal(msgRsv.GetMsgData(), &dataRsv)
```

## 示例

这有一些 [示例](https://github.com/lim-yoona/tcpack/tree/main/example)。  

