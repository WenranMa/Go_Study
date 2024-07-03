// client/selectsort.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 通过HTTP请求调用其他服务器上的add服务
	url := "http://127.0.0.1:9090/add"
	param := AddParam{
		X: 10,
		Y: 20,
	}
	paramBytes, _ := json.Marshal(param)
	resp, _ := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var respData AddResult
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data) // 30
}
