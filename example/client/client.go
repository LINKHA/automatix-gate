package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080/your-service/your-method"
	payload := []byte(`{"key": "value"}`) // 替换为实际的请求体数据

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json") // 如果需要设置其他请求头，请在此处添加

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
