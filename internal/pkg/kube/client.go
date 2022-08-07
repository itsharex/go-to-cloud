package kube

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	clientSet           *kubernetes.Clientset
	defaultApplyOptions *meta.ApplyOptions
}

// NewClient 创建k8s客户端对象
func NewClient(config *string) (*Client, error) {
	restCfg, err := clientcmd.BuildConfigFromFlags("", *config)

	if err != nil {
		return nil, err
	}

	m := meta.ApplyOptions{
		FieldManager: "application/apply-patch+yaml",
		Force:        true,
	}

	c, e := kubernetes.NewForConfig(restCfg)

	if e != nil {
		return nil, e
	} else {

		client := Client{
			clientSet:           c,
			defaultApplyOptions: &m,
		}
		return &client, nil
	}
}
