package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	servers := []Server{
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.baidu.com"),
		newSimpleServer("https://www.hao123.com"),
	}

	lb := NewLoadBalancer("8080", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.ServeProxy(w, r)
	}
	// 注册处理函数
	http.HandleFunc("/", handleRedirect)

	log.Printf("Starting server at 'localhost:%s'\n", lb.port)
	log.Fatal(http.ListenAndServe(":"+lb.port, nil))

}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
