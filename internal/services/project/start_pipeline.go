package project

import (
	"errors"
	"go-to-cloud/internal/builders"
	"go-to-cloud/internal/pkg/builder"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"time"
)

func build(nodeId, buildId uint, plan *repositories.Pipeline) (*kube.PodSpecConfig, error) {
	if node, err := repositories.GetBuildNodesById(nodeId); err != nil {
		return nil, err
	} else {
		return builders.BuildPodSpec(buildId, node, plan)
	}
}

func StartPipeline(userId uint, orgId []uint, projectId, pipelineId int64) error {
	plan, buildId, err := repositories.StartPlan(uint(projectId), uint(pipelineId), userId)
	if err == nil {
		if sortedIdleNodes, err := builder.ListNodesOnK8sOrderByIdle(orgId); err != nil {
			return err
		} else {
			if len(sortedIdleNodes) > 0 && sortedIdleNodes[0].Idle > 0 {
				node := sortedIdleNodes[0]
				podSpec, err := build(node.NodeId, buildId, plan)
				addBuildingPipelines(node.NodeId, buildId, uint(pipelineId), podSpec.TaskName)
				return err
			} else {
				return errors.New("没有足够可运行的构建节点，请稍后再试") // TODO: 未来计划使用构建队列
			}
		}
	}
	return err
}

func addBuildingPipelines(nodeId, buildId, pipelineId uint, taskName string) {
	pipelines := buildings[nodeId]
	if pipelines == nil {
		pipelines = []buildingPipeline{
			{
				NodeId:     nodeId,
				BuildId:    buildId,
				PipelineId: pipelineId,
				TaskName:   taskName,
				StartAt:    time.Now(),
			},
		}
	}
}
