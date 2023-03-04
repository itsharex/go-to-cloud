package kube

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

var podsCache *cache.Cache

func init() {
	podsCache = cache.New(time.Second*5, 0)
}

type PodDescription struct {
	Name        string
	Status      coreV1.PodPhase
	BuildId     uint // 构建时使用的Pod，对应pipeline id
	Containers  []string
	GetArtifact func(*string) *string // 获取产物
}

// GetPodLogs 读取容器日志
// tailLine: 从末尾开始的行数，nil时从开始显示
// previous: 返回之前已终止的容器日志
func (client *Client) GetPodLogs(ctx context.Context, ns, podName, containerName string, tailLine *int64, previous bool) ([]byte, error) {
	if s, err := client.GetPodStreamLogs(ctx, ns, podName, containerName, tailLine, false, previous); err != nil {
		return nil, err
	} else {
		return io.ReadAll(s)
	}
}

// GetPodStreamLogs 读取容器日志流
// tailLine: 从末尾开始的行数，nil时从开始显示
// previous: 返回之前已终止的容器日志
// follow: 是否跟踪获取最新日志，如果为true
func (client *Client) GetPodStreamLogs(ctx context.Context, ns, podName, containerName string, tailLine *int64, follow, previous bool) (io.ReadCloser, error) {
	logOpt := &coreV1.PodLogOptions{
		Container: containerName,
		Follow:    follow,
		TailLines: tailLine,
		Previous:  previous,
	}
	req := client.clientSet.CoreV1().Pods(ns).GetLogs(podName, logOpt)
	return req.Stream(ctx)
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
func (client *Client) GetPods(ctx context.Context, ns, labelPipeline string, labelSelector func() string, force bool) ([]PodDescription, error) {
	selector := labelSelector()
	key := fmt.Sprintf("%s.%s", ns, selector)
	if v, ok := podsCache.Get(key); !ok || force {
		pods, err := client.clientSet.CoreV1().Pods(ns).List(ctx, metaV1.ListOptions{
			LabelSelector: selector,
		})
		rlt := make([]PodDescription, len(pods.Items))
		for i, pod := range pods.Items {
			rlt[i] = PodDescription{
				BuildId: func() uint {
					if len(labelPipeline) > 0 {
						if label, ok := pod.GetObjectMeta().GetLabels()[labelPipeline]; ok {
							idStr := label[len(labelPipeline)+1:]
							if id, err := strconv.ParseUint(idStr, 10, 64); err == nil {
								return uint(id)
							}
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
			}
		}

		return rlt, err
	} else {
		return v.([]PodDescription), nil
	}
}

// DeletePod 删除指定pod
func (client *Client) DeletePod(ctx context.Context, ns, podName string) error {
	return client.clientSet.CoreV1().Pods(ns).Delete(ctx, podName, metaV1.DeleteOptions{})
}
