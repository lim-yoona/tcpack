# safetcpack
简体中文 | [English](README.md)  

safetcpack提供了tcpack的线程安全的版本。

## 与tcpack的区别
与tcpack不同的是，使用safetcpack，你可以为一个TCP连接构造多个打包器，并在多个goroutine中并发地使用他们。  

*注意：在同一个TCP连接上并发地使用多个打包器收发消息会导致发送和收到的消息顺序不确定，如果您想保证消息的按序发送和接收，请使用tcpack并且避免并发。*  

## 安装指南

1. 为安装 safetcpack 包, 首先你需要安装 [Go](https://go.dev/doc/install) , 然后你可以使用下面的命令将 `safetcpack` 作为你Go程序的依赖。

```sh
go get -u github.com/lim-yoona/tcpack
```

2. 在代码中导入 safetcpack ：

```go
import safetcpack "github.com/lim-yoona/tcpack/safe"
```

## 使用
safetcpack 的使用方法与 tcpack 一致。  

## 示例
Here are some [Examples](https://github.com/lim-yoona/tcpack/blob/main/example/concurrencyPack.go).  