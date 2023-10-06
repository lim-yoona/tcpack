# msgpack

[![Go Reference](https://pkg.go.dev/badge/github.com/lim-yoona/msgpack.svg)](https://pkg.go.dev/github.com/lim-yoona/msgpack)
![GitHub](https://img.shields.io/github/license/lim-yoona/msgpack)
![GitHub release (with filter)](https://img.shields.io/github/v/release/lim-yoona/msgpack)

English | [简体中文](README-CN.md)  


[msgpack](https://pkg.go.dev/github.com/lim-yoona/msgpack) is an application protocol based on TCP to Pack and Unpack bytes stream in [go](https://go.dev/) (or 'golang' for search engine friendliness) program.  

## What dose msgpack do?  

As we all know, TCP is a transport layer protocol oriented to byte streams. Its data transmission has no clear boundaries, so the data read by the application layer may contain multiple requests and cannot be processed.   

[msgpack](https://pkg.go.dev/github.com/lim-yoona/msgpack) is to solve this problem by encapsulating the request data into a message, packaging it when sending and unpacking it when receiving.  

## What's in the box?  

This library provides a packager which support Pack and Unpack.  

## Installation Guidelines

1. To install the msgpack package, you first need to have [Go](https://go.dev/doc/install) installed, then you can use the command below to add `msgpack` as a dependency in your Go program.  

```sh
go get -u github.com/lim-yoona/msgpack
```

2. Import it in your code:  

```go
import "github.com/lim-yoona/msgpack"
```

## Usage

```go
package main

import "github.com/lim-yoona/msgpack"

func main() {
    // Create a packager
    mp := msgpack.NewMsgPack(8, tcpConn)

    // Pack a message
    msg := msgpack.NewMessage(0, uint32(len([]byte(data))), []byte(data))
    msgByte, err := mp.Pack(msg)
    num, err := tcpConn.Write(msgByte)

    // Unpack a message
    msg, err := mp.Unpack()
}
```

## Examples

Here are some [Examples](https://github.com/lim-yoona/msgpack/tree/main/example).  

