package server

import (
	"context"
	"errors"
	gotocloud "go-to-cloud/internal/agent/proto"
	"go-to-cloud/internal/agent/vars"
	"golang.org/x/crypto/bcrypt"
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
	accessToken := metas["access-token"][0]
	err := bcrypt.CompareHashAndPassword([]byte(accessToken), []byte(vars.Ticket))
	return err == nil
}

func isAllowAnonymous(info *grpc.UnaryServerInfo) bool {
	whitelist := map[string]bool{
		"/gotocloud.Agent/Token": true,
	}

	return whitelist[info.FullMethod]
}

// AccessTokenInterceptor 拦截器
func AccessTokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	if !isAllowAnonymous(info) && !validAccessToken(ctx) {
		return nil, errors.New("invalid access token")
	}

	// 继续处理请求
	return handler(ctx, req)
}
