package client

import (
	"context"
	"go-to-cloud/internal/agent/vars"
)

// AccessTokenAuth 自定义Token认证
type AccessTokenAuth struct {
}

// GetRequestMetadata 获取元数据
func (c AccessTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"access-token": vars.Ticket,
	}, nil
}

// RequireTransportSecurity 是否开启传输安全 TLS
func (c AccessTokenAuth) RequireTransportSecurity() bool {
	return false
}
