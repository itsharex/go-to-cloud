package kube

import (
	"golang.org/x/net/context"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	applyApps "k8s.io/client-go/applyconfigurations/apps/v1"
	applyCore "k8s.io/client-go/applyconfigurations/core/v1"
	"strings"
)

// ApplyDeployment kubectl apply -f yml
func (client *Client) ApplyDeployment(namespace *string, yml *applyApps.DeploymentApplyConfiguration) (*apps.Deployment, error) {
	client.getOrCreateNamespace(namespace)
	return client.clientSet.AppsV1().Deployments(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

// ApplyService kubectl apply -f yml
func (client *Client) ApplyService(namespace *string, yml *applyCore.ServiceApplyConfiguration) (*core.Service, error) {
	client.getOrCreateNamespace(namespace)
	return client.clientSet.CoreV1().Services(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

const namespace_yml = `
apiVersion: v1
kind: Namespace
metadata:
  name: {{.Namespace}}
`

// getOrCreateNamespace 获取或创建namespace
func (client *Client) getOrCreateNamespace(namespace *string) (*core.Namespace, error) {

	cfg := strings.ReplaceAll(namespace_yml, "{{.Namespace}}", *namespace)
	yml := applyCore.NamespaceApplyConfiguration{}
	DecodeYaml(&cfg, &yml)

	return client.clientSet.CoreV1().Namespaces().Apply(context.TODO(), &yml, *client.defaultApplyOptions)
}
