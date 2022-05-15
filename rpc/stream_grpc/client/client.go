package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	streampb "rpc/stream_grpc/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 服务端流模式
	c := streampb.NewHelloClient(conn)
	res, _ := c.HelloServerStream(context.Background(), &streampb.HelloRequest{Name: "server_stream"})
	for {
		recv, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv)
	}

	// 客户端流模式
	clientStream, _ := c.HelloClientStream(context.Background())
	i := 0
	for {
		i++
		clientStream.Send(&streampb.HelloRequest{Name: fmt.Sprintf("lzl + %d", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}
