package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	helloworldpb "rpc/metadata/proto"
	"time"
)

func main() {

	// client interceptor
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		// invoker 为实际调用服务器的内容
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("use time: %s\n", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), opt)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworldpb.NewHelloClient(conn)

	// 在发送数据前包装 metadata
	// 放在 context 中进行传递
	md := metadata.New(map[string]string{
		"name":   "lzl",
		"passwd": "123",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, err := client.Hello(ctx, &helloworldpb.HelloRequest{Name: "lzl"})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Reply)
}
