package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := reflection.NewServerReflectionClient(conn)

	// 使用反射查找方法和消息
	listMethodsRequest := &grpc_reflection_v1alpha.ListServiceMethodsRequest{
		ServiceName: "your_service_name", // 替换为您要查找的服务名称
	}

	responseStream, err := client.ServerReflectionInfo(context.Background())
	if err != nil {
		log.Fatalf("Failed to get reflection info: %v", err)
	}

	// 发送 ListServiceMethods 请求
	err = responseStream.Send(&reflection.ServerReflectionRequest{
		MessageRequest: &reflection.ServerReflectionRequest_ListServices{
			ListServices: &reflection.ListServiceRequest{},
		},
	})

	if err != nil {
		log.Fatalf("Failed to send list services request: %v", err)
	}

	// 接收响应
	resp, err := responseStream.Recv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	// 处理响应
	for _, service := range resp.GetListServicesResponse().Service {
		log.Println("Service Name:", service.Name)

		// 发送 ListMethods 请求
		err = responseStream.Send(&grpc_reflection_v1alpha.ServerReflectionRequest{
			MessageRequest: &grpc_reflection_v1alpha.ServerReflectionRequest_ListMethods{
				ListMethods: listMethodsRequest,
			},
		})

		if err != nil {
			log.Fatalf("Failed to send list methods request: %v", err)
		}

		// 接收响应
		resp, err = responseStream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive response: %v", err)
		}

		// 处理响应
		for _, method := range resp.GetListMethodsResponse().Method {
			log.Println("Method Name:", method.GetName())
		}
	}
}
