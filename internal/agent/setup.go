package agent

import (
	"errors"
	"go-to-cloud/internal/repositories"
)

// Setup 安装agent至指定组织
func Setup(orgID uint) error {
	// 读取配置
	infra, err := repositories.FetchInfrastructures(orgID, repositories.InfraTypeAgent)

	if err != nil {
		return err
	}

	if len(infra) != 1 {
		return errors.New("没有找到agent所属k8s配置")
	}

	return nil
}
