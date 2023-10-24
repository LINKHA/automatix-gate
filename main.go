package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/LINKHA/automatix-gate/apigrpc" // 导入你的gRPC协议包
	"google.golang.org/grpc"
)

type GatewayServer struct {
	grpcClient pb.GateServerClient // 替换成你的gRPC客户端
}

func (s *GatewayServer) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	// 处理HTTP请求并将其转发给gRPC服务
	// 你可以根据需要在此处编写逻辑
}

func (s *GatewayServer) ForwardToGRPC(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("1--------------")

	// 使用gRPC客户端转发请求给其他gRPC服务
	response, err := s.grpcClient.SayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func MyHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// 处理HTTP请求并将其转发给gRPC服务
	// 你可以根据需要在此处编写逻辑
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGateServerServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 创建一个gRPC服务器
	grpcListener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGateServerServer(grpcServer, &server{})

	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// 创建一个HTTP服务器
	http.HandleFunc("/", MyHTTPHandler)
	go func() {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	log.Println("Gateway Server is running...")
	select {}
}
