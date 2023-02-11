package conf

import (
	"sync"
)

type Server struct {
	Url string
}

var server *Server

var onceGrpcHost sync.Once

// GetServerGrpcHost 获取服务器Grpc地址，用于被Agent访问
func GetServerGrpcHost() *Server {
	if server == nil {
		onceGrpcHost.Do(func() {
			if server == nil {
				server = &getConf().Server
			}
		})
	}
	return server
}
