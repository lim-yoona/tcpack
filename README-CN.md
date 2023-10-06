# msgpack
[English](README.md) | 简体中文 

[msgpack](https://pkg.go.dev/github.com/lim-yoona/msgpack@v0.1.0) 是一个基于 TCP 的应用层协议，用于在 [go](https://go.dev/) 程序中打包和解包字节流。  

## msgpack做了什么？  

众所周知，TCP 是面向字节流的传输层协议，其数据传输没有明确的边界，因此应用层读取的数据可能包含多个请求而导致无法处理业务。  

[msgpack](https://pkg.go.dev/github.com/lim-yoona/msgpack@v0.1.0) 就是为了解决这个问题，将请求数据封装成消息，发送时打包，接收时解包。  

## msgpack中有什么?  

msgpack 提供了一个支持 Pack 和 Unpack 的打包器。  

## 安装指南

1. 为安装 msgpack 包, 首先你需要安装 [Go](https://go.dev/doc/install) , 然后你可以使用下面的命令将 `msgpack` 作为你Go程序的依赖。    

```sh
go get -u github.com/lim-yoona/msgpack
```

2. 将 msgpack 导入到代码中：  

```go
import "github.com/lim-yoona/msgpack"
```

## 使用

```go
package main

import "github.com/lim-yoona/msgpack"

func main() {
    // 创建一个打包器
    mp := msgpack.NewMsgPack(8, tcpConn)

    // 打包一个消息
    msg := msgpack.NewMessage(0, uint32(len([]byte(data))), []byte(data))
    msgByte, err := mp.Pack(msg)
    num, err := tcpConn.Write(msgByte)

    // 解包一个消息
    msg, err := mp.Unpack()
}
```

## 示例

这有一些 [示例](https://github.com/lim-yoona/msgpack/tree/main/example)。  

