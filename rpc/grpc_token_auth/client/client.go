package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	authpb "rpc/grpc_token_auth/proto"
	"time"
)

// 使用 grpc 内置拦截器
type customCredential struct {
	
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"name":   "lzl",
		"passwd": "123",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

// 使用内置拦截器
func main() {

	// client interceptor
	opts := []grpc.DialOption{grpc.WithPerRPCCredentials(customCredential{}), grpc.WithInsecure()}

	conn, err := grpc.Dial("127.0.0.1:8088", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := authpb.NewHelloClient(conn)

	// 在发送数据前包装 metadata
	// 放在 context 中进行传递
	md := metadata.New(map[string]string{
		"appid": "101010",
		"appkey": "i am key",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// gRPC超时设置
	ctx_timeout, _ := context.WithTimeout(ctx, time.Second*3)


	response, err := client.Hello(ctx_timeout, &authpb.HelloRequest{Name: "lzl"})
	if err != nil {
		// 取到错误信息处理错误
		st, ok := status.FromError(err)
		if !ok {
			panic("cannot pares error")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	fmt.Println(response.Reply)
}


// 使用自定义拦截器
//func main() {
//
//	// client interceptor
//	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//		start := time.Now()
//
//		// 在客户端拦截器中加入一些信息 带着这些信息发送请求
//		md := metadata.New(map[string]string{
//			"appid": "101010",
//			"appkey": "i am key",
//		})
//		// 将元信息加入ctx  然后这个 md 就会跟着 ctx 传给服务器
//		ctx = metadata.NewOutgoingContext(context.Background(), md)
//
//		// invoker 为实际调用服务器的内容 即 invoker 中是调用服务器的逻辑
//		err := invoker(ctx, method, req, reply, cc, opts...)
//		fmt.Printf("use time: %s\n", time.Since(start))
//		return err
//	}
//	opt := grpc.WithUnaryInterceptor(interceptor)
//
//
//
//	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), opt)
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//
//	client := authpb.NewHelloClient(conn)
//
//	// 在发送数据前包装 metadata
//	// 放在 context 中进行传递
//	md := metadata.New(map[string]string{
//		"name":   "lzl",
//		"passwd": "123",
//	})
//	ctx := metadata.NewOutgoingContext(context.Background(), md)
//
//	response, err := client.Hello(ctx, &authpb.HelloRequest{Name: "lzl"})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(response.Reply)
//}
