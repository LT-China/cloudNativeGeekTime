package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/thinkeridea/go-extend/exnet"
)

func MyHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request => %v\n", req)
}

func HelloHandler(respW http.ResponseWriter, req *http.Request) {
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	reqHeaderContent, err := json.Marshal(req.Header)
	if nil != err {
		fmt.Printf("序列化有误 err=%v\n", err)
	} else {
		// fmt.Printf("序列化后的结果为: %v\n", reqHeaderContent)
		respW.Header().Add("reqHeaderContent", string(reqHeaderContent))
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	//当前PC无法添加VERSION环境变量，便提取了一个可以提取的value->USERNAME
	currentSystemUserName := os.Getenv("USERNAME")
	// fmt.Printf("Local Golang Version is: %v\n", currentSystemUserName)
	respW.Header().Add("currentSystemUserName", currentSystemUserName)

	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	// 3.1 客户端 IP
	remoteIP := exnet.ClientPublicIP(req)
	if remoteIP == "" {
		remoteIP = exnet.ClientIP(req)
	}
	fmt.Printf("Remote IP address: %v\n", remoteIP)
	//HTTP 返回码
	respW.WriteHeader(666)
	fmt.Printf("HTTP response code: %v\n", 666)

	fmt.Fprintf(respW, "<h1>%s<h1>", "Hello World ^_^")
}

//4.当访问 localhost/healthz 时，应返回200
func Healthz(respW http.ResponseWriter, req *http.Request) {
	respW.WriteHeader(200)
}
