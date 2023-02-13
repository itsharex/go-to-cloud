package project

import (
	"errors"
	"go-to-cloud/internal/pkg/builder"
	"go-to-cloud/internal/repositories"
)

func build(nodeId uint, plan *repositories.Pipeline) error {
	return errors.New("not implemented")
}

func StartPipeline(userId uint, orgId []uint, projectId, pipelineId int64) error {
	plan, err := repositories.StartPlan(uint(projectId), uint(pipelineId), userId)
	if err == nil {
		if sortedIdleNodes, err := builder.ListNodesOnK8sOrderByIdle(orgId); err != nil {
			return err
		} else {
			if len(sortedIdleNodes) > 0 && sortedIdleNodes[0].Idle > 0 {
				node := sortedIdleNodes[0]
				return build(node.NodeId, plan)
			} else {
				return errors.New("没有足够可运行的构建节点，请稍后再试")
			}
		}
	}
	return err
}
