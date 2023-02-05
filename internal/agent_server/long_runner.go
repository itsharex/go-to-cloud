package agent_server

import (
	"context"
	"errors"
	gotocloud "go-to-cloud/internal/agent_server/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"net/http"
	"sync"
)

type LongRunner struct {
	gotocloud.UnimplementedAgentServer

	clients map[string]*gotocloud.Agent_RunningServer // 连接的代理端
}

func (m *LongRunner) Execute(request *gotocloud.RunRequest) error {
	// TODO: RoundRobin 寻找可执行任务的节点
	for uuid, client := range m.clients {
		request.Uuid = uuid
		return (*client).Send(request)
	}
	return errors.New("not found build node")
}

func validAccessToken(ctx context.Context) bool {
	//获取元数据信息
	metas, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}
	accessToken := metas["access-token"][0]
	err := bcrypt.CompareHashAndPassword([]byte(accessToken), []byte(Ticket))
	return err == nil
}

func isAllowAnonymous(info *grpc.UnaryServerInfo) bool {
	whitelist := map[string]bool{
		"/gotocloud.LongRunner/Token": true,
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

func (m *LongRunner) Running(server gotocloud.Agent_RunningServer) error {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer func() {
			wg.Done()
		}()

		for {
			data, err := server.Recv()
			if err == io.EOF || data == nil {
				// 注销代理客户端
				for id, runServer := range m.clients {
					if runServer == &server {
						delete(m.clients, id)
					}
				}
				break
			} else {
				// 注册代理客户端
				if http.StatusCreated == data.GetCode() {
					uuid := data.GetUuid()
					m.clients[uuid] = &server
				}
			}
		}
	}()

	wg.Wait()
	return nil
}
