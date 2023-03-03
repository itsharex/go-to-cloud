package kube

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	clientSet           *kubernetes.Clientset
	defaultApplyOptions *meta.ApplyOptions
}

func (client *Client) GetClientSet() *kubernetes.Clientset {
	return client.clientSet
}

func NewClientByRestConfig(cfg *rest.Config) (*Client, error) {
	c, e := kubernetes.NewForConfig(cfg)

	m := meta.ApplyOptions{
		FieldManager: "application/apply-patch+yaml",
		Force:        true,
	}

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

// NewClient 创建k8s客户端对象
func NewClient(config *string) (*Client, error) {
	kubeConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*api.Config, error) {
		return clientcmd.Load([]byte(*config))
	})
	if err != nil {
		return nil, err
	}

	m := meta.ApplyOptions{
		FieldManager: "application/apply-patch+yaml",
		Force:        true,
	}

	c, e := kubernetes.NewForConfig(kubeConfig)

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
