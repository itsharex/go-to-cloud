package agent

import "go-to-cloud/internal/pkg/kube"

// Startup 启动Agent运行流水线
// k8sConfig: k8s登录文件
// namespace: 运行agent的名字空间
// image: agent镜像地址
// nodePort: agent对外服务端口
func Startup(k8sConfig, namespace, image *string, nodePort int) error {
	return kube.Apply(k8sConfig, namespace, &kube.AppDeployConfig{
		Name: "go-to-cloud-agent",
		Ports: []kube.Port{
			{
				ServicePort:   8080,
				ContainerPort: 8080,
				NodePort:      nodePort,
				PortName:      "go-to-cloud-agent",
			},
		},
		PortType: "NodePort",
		Image:    *image,
		Replicas: 1,
	})
}
