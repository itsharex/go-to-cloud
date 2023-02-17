package builder

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"strconv"
	"time"
)

var idleNodes *cache.Cache

// GetWorkingNodes 获取正在工作的节点数量
func GetWorkingNodes(workerId uint) (int, error) {
	node, err := repositories.GetBuildNodesById(workerId)
	if err != nil {
		return 0, err
	}

	if node == nil {
		return 0, errors.New("没有找到构建节点配置")
	}

	if node.NodeType == int(builder.K8s) {
		return func() (int, error) {
			if a, e := tryGetPodStatusFromCache(node, getPodStatus); e != nil {
				return 0, e
			} else {
				return len(a), nil
			}
		}()
	}

	return 0, errors.New("不支持的构建节点类型")
}

func getPodStatus(node *repositories.BuilderNode) ([]kube.PodStatus, error) {
	client, err := kube.NewClient(node.DecryptKubeConfig())
	if err != nil {
		return nil, err
	}

	return client.GetPods(node.K8sWorkerSpace, BuilderNodeSelectorLabel)
}

func tryGetPodStatusFromCache(node *repositories.BuilderNode, f func(node *repositories.BuilderNode) ([]kube.PodStatus, error)) ([]kube.PodStatus, error) {
	if v, ok := idleNodes.Get(strconv.Itoa(int(node.ID))); ok {
		return v.([]kube.PodStatus), nil
	} else {
		if n, e := f(node); e != nil {
			return nil, e
		} else {
			idleNodes.Set(strconv.Itoa(int(node.ID)), n, cache.DefaultExpiration)
			return n, nil
		}
	}
}

func init() {
	idleNodes = cache.New(5*time.Minute, 5*time.Minute)
}
