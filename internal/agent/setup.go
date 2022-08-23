package agent

import (
	"fmt"
	"go-to-cloud/internal/pkg/kube"
	v1 "k8s.io/api/core/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// Setup 将Agent以静态Pod方式安装至k8s
func Setup(k8sConfig, namespace *string, pod *kube.PodApplyConfig) (*v1.Pod, error) {
	client, err := kube.NewClient(k8sConfig)

	_, err = client.GetOrAddNamespace(namespace)
	if err != nil {
		panic(err)
	}

	yml, err := kube.GetYamlFromPodTemplate(pod)
	if err != nil {
		panic(err)
	}

	podConfig := corev1.PodApplyConfiguration{}
	if err := kube.DecodeYaml(yml, &podConfig); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return client.ApplyPod(namespace, &podConfig)
}
