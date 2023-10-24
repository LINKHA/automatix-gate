package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP调用
	httpURL := "http://localhost:8080" // 根据你的HTTP服务地址进行调整
	resp, err := http.Get(httpURL)
	if err != nil {
		log.Fatalf("HTTP call failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("HTTP Request was successful")
		// 处理HTTP响应的逻辑
	} else {
		log.Fatalf("HTTP Request failed with status code: %d", resp.StatusCode)
	}
}
