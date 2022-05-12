package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	helloworldpb "rpc/helloworld/proto"
)

type Server struct {
	helloworldpb.UnimplementedHelloServer
}

func (s *Server) Hello(ctx context.Context, request *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {
	return &helloworldpb.HelloResponse{
		Reply: "hello " + request.Name,
	}, nil
}

func main()  {
	server := grpc.NewServer()
	helloworldpb.RegisterHelloServer(server, &Server{})
	listener, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("fail to listen: " + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
