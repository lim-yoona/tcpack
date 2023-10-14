# safetcpack
English | [简体中文](README-CN.md)  

safetcpack provides a thread-safe version of tcpack.

## Difference from tcpack
Unlike tcpack, with safetcpack, you can construct multiple packers for a TCP connection and use them concurrently in multiple goroutines.  

*Notice: Using of multiple packagers based on the same TCP connection to send and receive messages concurrently can result in uncertain order of messages sent and received. If you want to ensure that messages are sent and received in order, please use tcpack and avoid concurrency.*

## Installation Guidelines
1. To install the safetcpack package, you first need to have [Go](https://go.dev/doc/install) installed, then you can use the command below to add `safetcpack` as a dependency in your Go program.

```sh
go get -u github.com/lim-yoona/tcpack
```

2. Import it in your code:

```go
import safetcpack "github.com/lim-yoona/tcpack/safe"
```

## Usage
The usage of safetcpack is consistent with that of tcpack.  

## Examples

Here are some [Examples](https://github.com/lim-yoona/tcpack/blob/main/example/concurrencyPack.go).  