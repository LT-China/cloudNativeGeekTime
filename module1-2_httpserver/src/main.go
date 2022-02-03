package main

import (
	"fmt"
	"module1-2_httpserver/src/handler"
	"net/http"
)

func main() {
	fmt.Println("********* [ Starting my first http service ] *********")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.MyHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/healthz", handler.Healthz)
	// http.ListenAndServe("127.0.0.1:8080", mux)
	http.ListenAndServe(":8080", mux)
}
