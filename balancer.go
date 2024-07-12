package main

import (
	"fmt"
	"net/http"
)

// LoadBalancer 负载均衡器
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

// getNextAvailableServer
// @Description: 获取下一个可用的服务器
// @receiver lb
// @return Server
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// ServeProxy
// @Description: 代理请求
// @receiver lb
// @param w
// @param r
func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextAvailableServer()
	fmt.Println("forwarding request to address: ", target.Address())
	target.Serve(w, r)
}
