package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 实现服务器接口
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleError(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	return true
}

func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
