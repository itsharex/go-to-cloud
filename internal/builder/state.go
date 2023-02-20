package builder

import (
	"context"
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"sync"
	"time"
)

type State struct {
}

// getAndSetNodesState 获取并更新所有构建节点信息
func getAndSetNodesState() {
	getAndSetK8sNodesState()
}

var allK8sPipelines map[uint]*kube.PodDescription

var artifactWatcher chan uint // pipeline_history.ID

func init() {
	allK8sPipelines = make(map[uint]*kube.PodDescription)
	artifactWatcher = make(chan uint, 5)
}

// PipelinesWatcher 流水线监测
func PipelinesWatcher() {
	// 定时获取构建节点状态
	go func() {
		for {
			c := time.Tick(time.Second * 15)
			<-c
			getAndSetNodesState()
		}
	}()

	// 制品生成监控
	go func() {
		for builderId := range artifactWatcher {
			go SaveDockImage(builderId)
		}
	}()
}

func getAndSetK8sNodesState() {
	var lock sync.Mutex
	if nodes, err := repositories.GetBuildNodesOnK8sByOrgId(nil, "", nil); err == nil {
		for _, node := range nodes {
			if client, err := kube.NewClient(node.DecryptKubeConfig()); err == nil {
				if pods, err := client.GetPods(context.TODO(), node.K8sWorkerSpace, NodeSelectorLabel, BuildIdSelectorLabel); err == nil {
					for i, pod := range pods {
						lock.Lock()
						allK8sPipelines[pod.BuildId] = &pods[i]
						lock.Unlock()

						rlt := func() pipeline.BuildingResult {
							switch pods[i].Status {
							case kube.Pending, kube.Running:
								return pipeline.UnderBuilding
							case kube.Succeeded:
								// TODO: 根据容器的状态最终确定构建结果
								return pipeline.BuildingSuccess
							case kube.Failed:
								return pipeline.BuildingFailed
							default:
								return pipeline.NeverBuild
							}
						}()
						if err := repositories.UpdatePipeline(pod.BuildId, rlt, &pods[i]); err == nil && pipeline.IsComplete(rlt) {
							// 清理Pod
							client.DeletePod(context.TODO(), node.K8sWorkerSpace, pod.Name)

							// 通知制品监视器流水线构建完成
							artifactWatcher <- pod.BuildId
						}
					}
				}
			}
		}
	}
}
