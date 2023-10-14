# tcpack

[![Go Reference](https://pkg.go.dev/badge/github.com/lim-yoona/tcpack.svg)](https://pkg.go.dev/github.com/lim-yoona/tcpack)
![GitHub](https://img.shields.io/github/license/lim-yoona/tcpack)
[![Go Report](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/lim-yoona/tcpack)
![GitHub release (with filter)](https://img.shields.io/github/v/release/lim-yoona/tcpack)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)

English | [简体中文](README-CN.md)  


[tcpack](https://pkg.go.dev/github.com/lim-yoona/tcpack) is an application protocol based on TCP to Pack and Unpack bytes stream in [go](https://go.dev/) (or 'golang' for search engine friendliness) program.  

## What dose tcpack do?  

As we all know, TCP is a transport layer protocol oriented to byte streams. Its data transmission has no clear boundaries, so the data read by the application layer may contain multiple requests and cannot be processed.   

[tcpack](https://pkg.go.dev/github.com/lim-yoona/tcpack) is to solve this problem by encapsulating the request data into a message, packaging it when sending and unpacking it when receiving.  

*notice: It is unsafe to use a packer to read and write messages concurrently on the same connection. Do not do this, as it will have unpredictable consequences!*

If you want to use multiple packagers based on the same TCP connection to send and receive messages concurrently, please use [safetcpack](https://github.com/lim-yoona/tcpack/tree/main/safe/README.md).  

## What's in the box?  

This library provides a packager which support Pack and Unpack.  

## Installation Guidelines

1. To install the tcpack package, you first need to have [Go](https://go.dev/doc/install) installed, then you can use the command below to add `tcpack` as a dependency in your Go program.  

```sh
go get -u github.com/lim-yoona/tcpack
```

2. Import it in your code:  

```go
import "github.com/lim-yoona/tcpack"
```

## Usage

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

### Support JSON

```go
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Create a packager
mp := tcpack.NewMsgPack(8, tcpConn)

// data JSON Marshal
data := &Person{
	Name: "jack",
	Age:  20,
}
dataJSON, _ := json.Marshal(data)

// Pack and send a message
msgSend := tcpack.NewMessage(0, uint32(len(dataJSON)), dataJSON)
num, err := mp.Pack(msgSend)

// Unpack and receive a message
msgRsv, err := mp.Unpack()

// JSON UnMarshal
var dataRsv Person
json.Unmarshal(msgRsv.GetMsgData(), &dataRsv)
```

## Examples

Here are some [Examples](https://github.com/lim-yoona/tcpack/tree/main/example).  

