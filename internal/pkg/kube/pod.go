package kube

import (
	"bytes"
	"context"
	"io"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
)

type PodDescription struct {
	Name       string
	Status     coreV1.PodPhase
	BuildId    uint
	Containers []string
	GetLog     func(*string) *string // 获取容器日志
}

/*
Pending（悬决）	Pod 已被 Kubernetes 系统接受，但有一个或者多个容器尚未创建亦未运行。此阶段包括等待 Pod 被调度的时间和通过网络下载镜像的时间。
Running（运行中）	Pod 已经绑定到了某个节点，Pod 中所有的容器都已被创建。至少有一个容器仍在运行，或者正处于启动或重启状态。
Succeeded（成功）	Pod 中的所有容器都已成功终止，并且不会再重启。
Failed（失败）	Pod 中的所有容器都已终止，并且至少有一个容器是因为失败终止。也就是说，容器以非 0 状态退出或者被系统终止。
Unknown（未知）	因为某些原因无法取得 Pod 的状态。这种情况通常是因为与 Pod 所在主机通信失败。
*/
const (
	Pending   coreV1.PodPhase = "Pending"
	Running   coreV1.PodPhase = "Running"
	Succeeded coreV1.PodPhase = "Succeeded"
	Failed    coreV1.PodPhase = "Failed"
	Unknown   coreV1.PodPhase = "Unknown"
)

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
			Status: pod.Status.Phase,
			Containers: func() []string {
				c := make([]string, len(pod.Spec.Containers))
				for i, container := range pod.Spec.Containers {
					c[i] = container.Name
				}
				return c
			}(),
			GetLog: func(c *string) *string {
				logBuilder := strings.Builder{}
				for _, container := range pod.Spec.Containers {
					name := container.Name
					logBuilder.WriteString("tl;dl;" + name + "\n")
					if podLogs, err := watchContainerLogWithPodNameAndContainerName(ctx, client.clientSet, ns, pod.Name, name, nil, false); err == nil {
						buf := new(bytes.Buffer)
						if _, err := io.Copy(buf, podLogs); err == nil {
							logBuilder.WriteString(buf.String())
							logBuilder.WriteString("\n")
						}
					}
				}
				log := logBuilder.String()
				return &log
			},
		}
	}

	return rlt, err
}

// DeletePod 删除指定pod
func (client *Client) DeletePod(ctx context.Context, ns, podName string) error {
	return client.clientSet.CoreV1().Pods(ns).Delete(ctx, podName, metaV1.DeleteOptions{})
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
