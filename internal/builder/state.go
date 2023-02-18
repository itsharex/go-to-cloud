package builder

import (
	"context"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"sync"
	"time"
)

type State struct {
}

// getAllNodesState 获取所有构建节点信息
func getAllNodesState() {
	getK8sNodesState()
}

var allK8sPipelines map[uint]*kube.PodDescription

func GetPodDescription(pipelineId uint) *kube.PodDescription {
	return allK8sPipelines[pipelineId]
}

func init() {
	allK8sPipelines = make(map[uint]*kube.PodDescription)
}

// PipelinesWatcher 流水线监测
func PipelinesWatcher() {
	// 定时获取构建节点状态
	go func() {
		for {
			c := time.Tick(time.Second * 15)
			<-c
			getAllNodesState()
		}
	}()
}

func getK8sNodesState() {
	var lock sync.Mutex
	if nodes, err := repositories.GetBuildNodesOnK8sByOrgId(nil, "", nil); err == nil {
		for _, node := range nodes {
			if client, err := kube.NewClient(node.DecryptKubeConfig()); err == nil {
				if pods, err := client.GetPods(context.TODO(), node.K8sWorkerSpace, NodeSelectorLabel, BuildIdSelectorLabel); err == nil {
					for _, pod := range pods {
						lock.Lock()
						allK8sPipelines[pod.BuildId] = &pod
						lock.Unlock()
					}
				}
			}
		}
	}
}
