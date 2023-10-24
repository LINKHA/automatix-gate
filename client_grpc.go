package main

import (
	"context"
	"log"

	pb "github.com/LINKHA/automatix-gate/apigrpc" // 导入你的gRPC协议包
	"google.golang.org/grpc"
)

func main() {
	// 连接到gRPC服务器
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	// 创建gRPC客户端
	grpcClient := pb.NewGateServerClient(grpcConn)

	// gRPC调用
	grpcRequest := &pb.HelloRequest{Name: "YourName"}
	grpcResponse, err := grpcClient.SayHello(context.Background(), grpcRequest)

	if err != nil {
		log.Fatalf("gRPC call failed: %v", err)
	}
	log.Printf("gRPC Response: %s", grpcResponse.Message)
}
