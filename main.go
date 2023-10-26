package main

import (
	"fmt"

	dynamic_call "github.com/LINKHA/automatix-common/dynamic_call"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AddRouter() {
	r := gin.Default()

	// 统一处理跨域请求
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	})

	// 定义路由规则和处理函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		// 处理创建用户的逻辑
	})

	r.GET("/users/:id", func(c *gin.Context) {
		// 处理获取特定用户信息的逻辑
	})

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path // 获取请求的路径
		method := c.Request.Method // 获取请求的方法

		// 输出请求信息
		fmt.Printf("Received request: %s %s\n", method, path)

		// 返回自定义的响应
		c.JSON(404, gin.H{
			"message": "Page not found....................",
		})
	})

	// 启动服务
	r.Run(":8080")

}

func AddClientCall() {
	fmt.Println("start...")
	defer fmt.Println("end...")

	dynamic_call.SetProtoSetFiles("dynamic_call.protoset")
	err := dynamic_call.InitDescSource()

	if err != nil {
		panic(err.Error())
	}

	var handler = DefaultEventHandler{}
	var sendBody string

	grpcEnter, err := dynamic_call.New(
		dynamic_call.SetHookHandler(&handler),
	)
	grpcEnter.Init()

	sendBody = `{"name": "golang world"}`
	res, err := grpcEnter.Call("127.0.0.1:50051", "apigrpc.Greeter", "SayHello", sendBody)
	fmt.Printf("%+v \n", res)
	fmt.Println("req/reply return err: ", err)
}

type DefaultEventHandler struct {
	sendChan chan []byte
}

func (h *DefaultEventHandler) OnReceiveData(md metadata.MD, resp string, respErr error) {
}

func (h *DefaultEventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
}

func main() {
	AddRouter()
}
