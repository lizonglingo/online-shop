package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	authpb "rpc/grpc_token_auth/proto"

)


type Server struct {
	authpb.UnimplementedHelloServer
}

func (s *Server) Hello(ctx context.Context, request *authpb.HelloRequest) (*authpb.HelloResponse, error) {

	// 从 context 中取出数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	} else {
		for k, v := range md {
			fmt.Println(k, v)
		}
	}

	return &authpb.HelloResponse{
		Reply: "hello " + request.Name,
	}, nil
}


func main()  {

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
		fmt.Println("get a new request")

		// 从 ctx 取出一些值
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("get metadata error")
			return resp, status.Error(codes.Unauthenticated, "cannot get token")
		}

		var (
			appid string
			appkey string
		)
		// 需要特别注意  这 是 slice 类型
		if aid, ok := md["appid"]; ok {
			appid = aid[0]
		}
		if akey, ok := md["appkey"]; ok {
			appkey = akey[0]
		}

		if appid != "101010" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "cannot get token")
		}


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
	authpb.RegisterHelloServer(server, &Server{})
	listener, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("fail to listen: " + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
