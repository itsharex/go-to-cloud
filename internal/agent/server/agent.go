package server

import (
	"context"
	"errors"
	gotocloud "go-to-cloud/internal/agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Agent struct {
	gotocloud.UnimplementedAgentServer
}

func validAccessToken(ctx context.Context) bool {
	//获取元数据信息
	metas, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}
	accessToken := metas["access-token"]
	_ = accessToken[0]
	return true
}

// AccessTokenInterceptor 拦截器
func AccessTokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !validAccessToken(ctx) {
		return nil, errors.New("invalid access token")
	}

	// 继续处理请求
	return handler(ctx, req)
}
