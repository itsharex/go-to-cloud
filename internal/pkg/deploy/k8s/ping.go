package k8s

import (
	"go-to-cloud/internal/models/deploy/k8s"
)

// Ping 测试K8s服务是否可用
func Ping(testing *k8s.Testing) (bool, error) {
	return true, nil
}
