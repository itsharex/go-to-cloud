package project

import (
	agent "go-to-cloud/internal/agent_server"
	gotocloud "go-to-cloud/internal/agent_server/proto"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
	"strings"
)

func StartPipeline(userId uint, orgId []uint, projectId, pipelineId int64) error {
	plan, err := repositories.StartPlan(uint(projectId), uint(pipelineId), userId)
	if err == nil {
		// look up available agent; todo: belongs org
		if nodes, err := repositories.GetBuildNodesOnK8sByOrgId(orgId, "", nil); err != nil {
			return err
		} else {
			workIdAndMax := func() map[int]int {
				rlt := make(map[int]int)
				for _, node := range nodes {
					rlt[int(node.ID)] = node.MaxWorkers
				}
				return rlt
			}()
			command := setRequestCommand(workIdAndMax, plan)
			agent.Runner.Execute(command)
		}
	}
	return err
}

// setRequestCommand 组装调用命令
// nodes: 配置的节点, key: 节点ID（workID）; value：最大可运行任务
func setRequestCommand(nodes map[int]int, plan *repositories.Pipeline) *gotocloud.RunRequest {
	for nodeId, max := range nodes {
		_ = max // max用于控制不超任务最大数量
		workId := int64(nodeId)
		nodes := agent.Runner.GetNodeCount(workId)
		if nodes > 0 {
			// TODO: 选择最空闲的节点
			req := &gotocloud.RunRequest{WorkId: workId}

			req.Command = makeGitCloneRequestCommand(plan)

			cmd := req.Command
			for _, step := range plan.PipelineSteps {
				makeShellRequestCommand(cmd, &step)
				cmd = cmd.Next
			}
			return req
		}
	}
	return nil
}

func makeGitCloneRequestCommand(plan *repositories.Pipeline) *gotocloud.RunCommandRequest {
	return &gotocloud.RunCommandRequest{
		Command: "git",
		Args: func() []string {
			args := make([]string, 3)
			args[0] = plan.SourceCode.GitUrl
			args[1] = plan.Branch
			args[2] = utils.Base64AesEny([]byte(plan.SourceCode.CodeRepo.AccessToken))
			return args
		}(),
	}
}

var i int

func makeShellRequestCommand(command *gotocloud.RunCommandRequest, step *repositories.PipelineSteps) {
	// dotnet test --collect:"XPlat Code Coverage" --logger "html;logfilename=testresults.html"

	params := strings.SplitN(step.Script, " ", 2)
	command.Next = &gotocloud.RunCommandRequest{
		Command: params[0],
		Args: func() []string {
			if len(params) > 1 {
				return params[1:]
			} else {
				return nil
			}
		}(),
	}
}
