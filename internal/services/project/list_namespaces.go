package project

import (
	"errors"
	k8smodel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/deploy/k8s"
)

// ListNamespacesByOrg 用于新建
func ListNamespacesByOrg(orgId []uint) ([]string, error) {
	k8sRepo, err := k8s.List(orgId, &k8smodel.Query{})

	if err != nil {
		return nil, err
	}

	for _, s := range k8sRepo {
		_ = s
	}

	return nil, errors.New("not implemented")
}
