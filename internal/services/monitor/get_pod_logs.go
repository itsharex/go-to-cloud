package monitor

import (
	"context"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func FollowLogs(ctx context.Context, deploymentId, k8sId uint, podName string, previous bool, logs func([]byte)) error {
	repo, err := repositories.QueryK8sRepoById(k8sId)
	if err != nil {
		return err
	}

	deployment, err := repositories.GetDeploymentById(deploymentId)
	if err != nil {
		return err
	}

	client, err := kube.NewClient(&repo.KubeConfig)
	if err != nil {
		return err
	}

	lines := int64(100)
	stream, err := client.GetPodStreamLogs(ctx, deployment.K8sNamespace, podName, deployment.ArtifactDockerImageRepo.Name, &lines, true, previous)
	if err != nil {
		return err
	}

	defer stream.Close()

	p := make([]byte, 1024)
	for {
		n, err := stream.Read(p)
		if err != nil {
			return err
		}

		logs(p[:n])
	}
}
