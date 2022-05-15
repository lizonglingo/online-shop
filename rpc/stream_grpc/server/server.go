package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	streampb "rpc/stream_grpc/proto"
	"time"
)

const PORT = ":50052"

type server struct {
	streampb.UnimplementedHelloServer
}

// HelloServerStream 服务端流模式 客户端发出请求 服务端以 流消息响应.
func (s *server) HelloServerStream(request *streampb.HelloRequest, streamServer streampb.Hello_HelloServerStreamServer) error {
	i := 0
	for {
		streamServer.SendMsg(&streampb.HelloResponse{
			Reply: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}
	return nil
}

// HelloClientStream 客户端流模式 客户端以流模式发送请求.
func (s *server) HelloClientStream(streamServer streampb.Hello_HelloClientStreamServer) error {
	for {
		if recv, err := streamServer.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(recv.Name)
		}
	}
	return nil
}

func (s *server) HelloStream(streamServer streampb.Hello_HelloStreamServer) error {
	panic("implement me")
}

func main()  {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	streampb.RegisterHelloServer(s, &server{})
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
