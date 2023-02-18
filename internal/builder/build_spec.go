package builder

import (
	lang2 "go-to-cloud/internal/builder/lang"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"strconv"
	"strings"
)

func ResetIdle(node *repositories.BuilderNode) {
	idleNodes.Delete(strconv.Itoa(int(node.ID)))
}

// BuildPodSpec 创建构建模板 k8s pod spec
// buildId: pipeline_history.ID
func BuildPodSpec(buildId uint, node *repositories.BuilderNode, plan *repositories.Pipeline) *kube.PodSpecConfig {
	var lang lang2.Tpl
	switch plan.Env {
	case lang2.DotNet3, lang2.DotNet5, lang2.DotNet6, lang2.DotNet7:
		lang = &lang2.DotNet{}
	case lang2.Go120, lang2.Go116, lang2.Go119, lang2.Go118, lang2.Go117:
		lang = &lang2.Golang{}
	}
	buildIdStr := strconv.Itoa(int(buildId))
	return &kube.PodSpecConfig{
		LabelFlag:    NodeSelectorLabel,
		LabelBuildId: BuildIdSelectorLabel,
		BuildId:      buildId,
		Namespace:    node.K8sWorkerSpace,
		TaskName: plan.Name + "-" + func(exceptedLen int) string {
			if len(buildIdStr) >= exceptedLen {
				return buildIdStr
			}
			padding := strings.Repeat("0", exceptedLen-len(buildIdStr))
			return padding + buildIdStr
		}(5),
		SourceCode: plan.SourceCode.GitUrl,
		Sdk:        lang.Sdk(plan.Env),
		Steps: func() []kube.Step {
			kvp := lang.Steps(plan.Env, plan.PipelineSteps)
			steps := make([]kube.Step, len(kvp))
			i := 0
			for t, cmd := range kvp {
				steps[i] = kube.Step{
					Command:     cmd,
					CommandType: (&t).GetTypeName(),
				}
				i++
			}
			return steps
		}(),
	}
}
