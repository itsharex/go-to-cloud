package kube

import (
	"context"
	"io"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

type PodStatus struct {
	Name   string
	Status coreV1.PodPhase
}

// GetPods 获取指定名字空间
func (client *Client) GetPods(ns, label string) ([]PodStatus, error) {
	pods, err := client.clientSet.CoreV1().Pods(ns).List(context.TODO(), metaV1.ListOptions{
		LabelSelector: "builder=" + label,
	})

	rlt := make([]PodStatus, len(pods.Items))
	for i, pod := range pods.Items {
		rlt[i] = PodStatus{
			Name:   pod.Name,
			Status: getPodStatus(&pod),
		}
	}

	return rlt, err
}

// getPodStatus 获取Pod状态
func getPodStatus(pod *coreV1.Pod) coreV1.PodPhase {
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
	return req.Stream(context.TODO())
}
