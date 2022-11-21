package server

import (
	"context"
	"errors"
	gotocloud "go-to-cloud/internal/agent/proto"
	"go-to-cloud/internal/agent/vars"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (m *Agent) Token(ctx context.Context, request *gotocloud.AccessTokenRequest) (*gotocloud.AccessTokenResponse, error) {
	if strings.EqualFold(request.Ticket, vars.Ticket) {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(vars.Ticket), bcrypt.DefaultCost)
		return &gotocloud.AccessTokenResponse{AccessToken: string(hashBytes)}, err
	} else {
		return nil, errors.New("invalid ticket")
	}
}
