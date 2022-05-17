package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	helloworldpb "rpc/metadata/proto"
)

type Server struct {
	helloworldpb.UnimplementedHelloServer
}

func (s *Server) Hello(ctx context.Context, request *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {

	// 从 context 中取出数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	} else {
		for k, v := range md {
			fmt.Println(k, v)
		}
	}

	return &helloworldpb.HelloResponse{
		Reply: "hello " + request.Name,
	}, nil
}


func main()  {

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
		fmt.Println("get a new request")
		// handler 中是接下来要处理的逻辑 直接将 context 和 request 传下去继续执行即可
		// 也可以直接 return handler(ctx, req)
		res, err := handler(ctx, req)
		fmt.Println("request finished")
		return res, err
	}


	// 生成 grpc option
	opt := grpc.UnaryInterceptor(interceptor)
	// 将 option 加入 Server
	server := grpc.NewServer(opt)
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
