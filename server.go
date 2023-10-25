package main

import (
	"context"
	"log"
	"net"

	pb "github.com/LINKHA/automatix-gate/apigrpc" // 导入你的gRPC协议包

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) YourMethod(ctx context.Context, in *pb.YourRequest) (*pb.YourResponse, error) {
	return &pb.YourResponse{Message: "Hello, " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
