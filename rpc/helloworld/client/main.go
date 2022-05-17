package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	helloworldpb "rpc/helloworld/proto-bak"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworldpb.NewHelloClient(conn)
	response, err := client.Hello(context.Background(), &helloworldpb.HelloRequest{Name: "lzl"})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Reply)
}
