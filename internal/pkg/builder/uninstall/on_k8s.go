package uninstall

import (
	"errors"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// OnK8s 卸载K8s节点
func OnK8s(userId, nodeId uint) error {
	if userId <= 0 || nodeId <= 0 {
		return errors.New("not allowed")
	}

	if node, err := repositories.GetBuildNodesById(nodeId); err != nil {
		return err
	} else {
		kubeCfg := node.DecryptKubeConfig()
		if client, err := kube.NewClient(kubeCfg); err != nil {
			return err
		} else {
			err = client.DeleteNamespace(&node.K8sWorkerSpace)
			return repositories.DeleteBuilderNode(userId, nodeId)
		}
	}
}
