package main

import (
	"google.golang.org/grpc"
	"net"
	streampb "rpc/stream_grpc/proto"
	"text/template/parse"
)

const PORT = ":50052"

type server struct {
	streampb.UnimplementedHelloServer
}

func (s *server) HelloServerStream(request *streampb.HelloRequest, streamServer streampb.Hello_HelloServerStreamServer) error {

	panic("implement me")
}

func (s *server) HelloClientStream(streamServer streampb.Hello_HelloClientStreamServer) error {
	panic("implement me")
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
}
