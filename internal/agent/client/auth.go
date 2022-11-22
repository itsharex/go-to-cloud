package client

import (
	"context"
	"go-to-cloud/internal/agent/vars"
	"golang.org/x/crypto/bcrypt"
)

// AccessTokenAuth 自定义Token认证
type AccessTokenAuth struct {
}

// GetRequestMetadata 获取元数据
func (c AccessTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// TODO: 增加动态身份认证
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(vars.Ticket), bcrypt.DefaultCost)
	return map[string]string{
		"access-token": string(hashBytes),
	}, nil
}

// RequireTransportSecurity 是否开启传输安全 TLS
func (c AccessTokenAuth) RequireTransportSecurity() bool {
	return false
}
