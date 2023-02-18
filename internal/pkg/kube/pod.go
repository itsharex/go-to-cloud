package kube

import (
	"bytes"
	"context"
	"go-to-cloud/internal/repositories"
	"io"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"time"
)

const (
	ContainersReady string = "ContainersReady"
	PodInitialized  string = "Initialized"
	PodReady        string = "Ready"
	PodScheduled    string = "PodScheduled"
)

const (
	ConditionTrue    string = "True"
	ConditionFalse   string = "False"
	ConditionUnknown string = "Unknown"
)

type PodDescription struct {
	Name       string
	Status     coreV1.PodPhase
	BuildId    uint
	Containers []string
	GetLog     func(*string) *bytes.Buffer // 获取容器日志

	lastRefresh *time.Time // 最新一次更新时间
}

func (m *PodDescription) Refresh() {
	now := time.Now()
	if m.lastRefresh != nil && now.Sub(*m.lastRefresh) > time.Second*25 {
		if buf := m.GetLog(nil); buf != nil {
			log := buf.String()
			repositories.UpdateBuildLog(m.BuildId, &log)
			m.lastRefresh = &now
		}
	}
}

// GetPods 获取指定名字空间
func (client *Client) GetPods(ctx context.Context, ns, label, labelPipeline string) ([]PodDescription, error) {
	pods, err := client.clientSet.CoreV1().Pods(ns).List(context.TODO(), metaV1.ListOptions{
		LabelSelector: "builder=" + label,
	})

	rlt := make([]PodDescription, len(pods.Items))
	for i, pod := range pods.Items {
		rlt[i] = PodDescription{
			BuildId: func() uint {
				if label, ok := pod.GetObjectMeta().GetLabels()[labelPipeline]; ok {
					idStr := label[len(labelPipeline)+1:]
					if id, err := strconv.ParseUint(idStr, 10, 64); err == nil {
						return uint(id)
					}
				}
				return 0
			}(),
			Name:   pod.Name,
			Status: getPodDescription(&pod),
			Containers: func() []string {
				c := make([]string, len(pod.Spec.Containers))
				for i, container := range pod.Spec.Containers {
					c[i] = container.Name
				}
				return c
			}(),
			GetLog: func(c *string) *bytes.Buffer {
				if podLogs, err := watchContainerLogWithPodNameAndContainerName(ctx, client.clientSet, ns, pod.Name, "", nil, false); err == nil {
					buf := new(bytes.Buffer)
					if _, err := io.Copy(buf, podLogs); err == nil {
						return buf
					}
				}
				return nil
			},
		}
	}

	return rlt, err
}

// getPodDescription 获取Pod状态
func getPodDescription(pod *coreV1.Pod) coreV1.PodPhase {
	for _, cond := range pod.Status.Conditions {
		if string(cond.Type) == ContainersReady {
			if string(cond.Status) != ConditionTrue {
				return "Unavailable"
			}
		} else if string(cond.Type) == PodInitialized && string(cond.Status) != ConditionTrue {
			return "Initializing"
		} else if string(cond.Type) == PodReady {
			if string(cond.Status) != ConditionTrue {
				return "Unavailable"
			}
			for _, containerState := range pod.Status.ContainerStatuses {
				if !containerState.Ready {
					return "Unavailable"
				}
			}
		} else if string(cond.Type) == PodScheduled && string(cond.Status) != ConditionTrue {
			return "Scheduling"
		}
	}
	return pod.Status.Phase
}

// watchContainerLogWithPodNameAndContainerName
// 读取容器日志流
// 入参：
// Container:容器名称
// Follow:跟踪Pod的日志流，默认为false（关闭）对应kubectl logs命令中的 -f 参数
// TailLines:如果设置，则显示从日志末尾开始的行数。如果未指定，则从容器的创建开始或从秒开始或从时间开始显示日志
// Previous:返回以前终止的容器日志。默认为false（关闭）
func watchContainerLogWithPodNameAndContainerName(
	ctx context.Context,
	client *kubernetes.Clientset,
	namespace,
	podName,
	containerName string, tailLine *int64, previous bool) (io.ReadCloser, error) {
	logOpt := &coreV1.PodLogOptions{
		Container: containerName,
		Follow:    true,
		TailLines: tailLine,
		Previous:  previous,
	}
	req := client.CoreV1().Pods(namespace).GetLogs(podName, logOpt)
	return req.Stream(ctx)
}
