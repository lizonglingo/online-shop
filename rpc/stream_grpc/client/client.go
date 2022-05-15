package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	streampb "rpc/stream_grpc/proto"
	"sync"
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
	// 与服务端建立一个 服务端流通道
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
	// 首先与服务端建立一个 客户端流通道
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

	// 双向流模式
	// 首先与服务端建立一个双向流通道
	stream, _ := c.HelloStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			recv, _ := stream.Recv()
			fmt.Println(recv.Reply)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			 stream.Send(&streampb.HelloRequest{Name: "i am client"})
			 time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
