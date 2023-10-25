package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/LINKHA/automatix-gate/apigrpc" // 导入你的gRPC协议包
)

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
	listen, err := net.Listen("tcp", ":50051") // 替换成您想要监听的地址和端口
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGateServerServer(s, &server{})

	reflection.Register(s) // 启用服务器反射功能

	log.Println("Server started")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
