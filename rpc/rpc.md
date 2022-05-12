# RPC带来的三个主要问题

## Call的id映射

### 本地调用

本地调用中，编译器自动帮我们调取函数的指针，因此可以知道调用的是哪个函数。

### 远程调用

而如何知道远程服务器上的函数在什么地方呢，不同的地址空间里，函数指针位置完全不同。所以在RPC中，函数需要有一个自己的ID，这个**ID**在所有进程中唯一确定。

同时，客户端和服务器分别维护一个**函数 <--> Call ID**的映射表，两个表不一定一样，但是同一个函数对应的Call ID一定是一样的，在进行RPC时，客户端找到要调用函数的Call ID，传给服务器，服务器就知道要调用哪个函数，并且需要什么参数等信息。

## 序列化和反序列化

跨服务器调度一定会带来网络数据传输问题，这样就引出下面几个点：

1. 使用Web服务吗？
2. 函数参数如何传递，使用什么编码协议？json？xml？protobuf？msgpack？

### 本地调用

例如，某个函数的参数`Book`为：

```go
type Book struct {
    Name 	string
    Page 	int
    Auth	Author
}

type Author struct {
    Name 	string
    Age  	int
    Country	string
}
```

如果在本地调用，我们可以直接传给函数，函数知道如何接收这个格式的参数。

### 远程调用

但是，显然不能直接将上面的结构体传给远程服务器。需要通过网络，以某种格式（如二进制）将其进行编码传输，到达目标服务器后，再将其反编码（如解码二进制）成函数可以理解的形式。

> JSON的局限性：
>
> - 复杂项目中，使用JSON协议，对数据接口维护极难极繁琐
> - JSON数据传输涉及的序列化和反序列化性能不够好
> - 客户端和服务端不同的实现语言带来额外的序列化、反序列化工作

## 网络传输

大部分RCP框架都使用TCP协议，也有使用UDP协议的，此外还可使用其他协议，如gRPC使用     HTTP 2.0 。

> HTTP 的缺陷：
>
> - 一次性，对长连接支持性差，多次请求响应需要重复建立连接
>
> **在 HTTP 2.0 中已经支持长连接**，gRPC就是使用的 HTTP 2.0 。

# gRPC和Protocol Buffer

## Protocol Buffer

一个高压缩比，高性能，序列化和反序列化高效，轻量级数据存储协议。

- 自动生成序列化和反序列化代码
- 只需要维护proto文件
- 向后兼容
- 加密性好
- 跨平台
- 支持各种主流语言



## gRPC

### proto文件格式规范

- Packege

为生成Go代码，对每个`.proto`文件都需要添加Go package相关的内容。可以通过：

1. 在`.proto`文件中声明（推荐）
2. 在命令行中声明

两种方式指明。

`go_package`需要有完整的`import`路径，如：

```protobuf
option go_package = "example.com/project/protos/fizz";
```

示例：

```protobuf
syntax = "proto3";
package helloworld;
option go_package="rpc/helloword/proto;helloworldpb";

message HelloRequest {
  string name = 1;
  int32 age = 2;
  repeated string courses = 3;
}

message HelloResponse {
  string reply = 1;
}

service Hello {
  rpc Hello(HelloRequest) returns (HelloResponse);
}
```

### 文件生成命令

```shell
protoc -I={proto文件所在的 相对于执行这条命令的 目录路径} --go_out={生成的 .pb.go 存放的目录路径} --go_opt=paths=source_relative --go-grpc_out={生成的 xxx_grpc.pb.go 存放的目录路径} --go-grpc_opt=paths=source_relative {具体的文件名 如 example.proto 该文件应该在 -I= 所指的目录中}
```

示例：

```shell
D:\Coding\WorkPlace\Golang\online-shop\rpc\helloworld\proto>protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto
```

目录结构为：

![image-20220512115147137](https://picgo-lzl.oss-cn-beijing.aliyuncs.com/image-20220512115147137.png)











