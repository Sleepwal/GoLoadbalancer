package main

import "net/http"

// Server 服务器接口
type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}
