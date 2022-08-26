package kube

import (
	"fmt"
	"golang.org/x/net/context"
	core "k8s.io/api/core/v1"
	applyCore "k8s.io/client-go/applyconfigurations/core/v1"
	meta "k8s.io/client-go/applyconfigurations/meta/v1"
	"strings"
)

// GetOrAddNamespace 获取或创建名字空间
func (client *Client) GetOrAddNamespace(ns *string) (*core.Namespace, error) {

	kind := "Namespace"
	apiVer := "v1"
	namespace := applyCore.NamespaceApplyConfiguration{
		TypeMetaApplyConfiguration: meta.TypeMetaApplyConfiguration{
			Kind:       &kind,
			APIVersion: &apiVer,
		},
		ObjectMetaApplyConfiguration: &meta.ObjectMetaApplyConfiguration{
			Name: ns,
		},
	}
	rlt, err := client.clientSet.CoreV1().Namespaces().Apply(context.TODO(), &namespace, *client.defaultApplyOptions)

	if err != nil {
		fmt.Println(err)
	}
	return rlt, err
}

// getOrCreateNamespace 获取或创建namespace
func (client *Client) getOrCreateNamespace(namespace *string) (*core.Namespace, error) {

	cfg := strings.ReplaceAll(namespace_yml, "{{.Namespace}}", *namespace)
	yml := applyCore.NamespaceApplyConfiguration{}
	DecodeYaml(&cfg, &yml)

	return client.clientSet.CoreV1().Namespaces().Apply(context.TODO(), &yml, *client.defaultApplyOptions)
}
