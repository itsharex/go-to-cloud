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

type ClientAgentMap map[string]*gotocloud.Agent_RunningServer // key: node的全局唯一标识；val: 连接对象

type LongRunner struct {
	gotocloud.UnimplementedAgentServer
	locker    sync.Mutex
	nodesPool map[int64]ClientAgentMap // 节点包含的代理端池；key: build_nodes.ID; value: ClientAgentMap
}

// GetNodeCount 获取可运行的节点数量
func (m *LongRunner) GetNodeCount(workId int64) int {
	return len(m.nodesPool[workId])
}

func (m *LongRunner) Execute(request *gotocloud.RunRequest) error {
	for uuid, client := range m.nodesPool[request.GetWorkId()] {
		request.Uuid = uuid
		return (*client).Send(request) // 使劲薅第一台节点；TODO: RoundRobin 寻找可执行任务的节点
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
				deleted := false
				m.locker.Lock()
				for workId, agentMap := range m.nodesPool {
					for uuid, runServer := range agentMap {
						if runServer == &server {
							delete(m.nodesPool[workId], uuid)
							deleted = true
							break
						}
					}
					if deleted {
						break
					}
				}
				m.locker.Unlock()
			} else {
				// 注册代理客户端
				if http.StatusCreated == data.GetCode() {
					uuid := data.GetUuid()
					workId := data.GetWorkId()
					m.locker.Lock()
					if len(m.nodesPool[workId]) == 0 {
						m.nodesPool[workId] = make(ClientAgentMap)
					}
					m.nodesPool[workId][uuid] = &server
					m.locker.Unlock()
				}
			}
		}
	}()

	wg.Wait()
	return nil
}
