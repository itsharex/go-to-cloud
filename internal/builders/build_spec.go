package builders

import (
	"errors"
	lang2 "go-to-cloud/internal/builders/lang"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// BuildPodSpec 创建构建模板 k8s pod spec
func BuildPodSpec(node *repositories.BuilderNode, plan *repositories.Pipeline) error {
	var lang lang2.Tpl
	switch plan.Env {
	case lang2.DotNet3, lang2.DotNet5, lang2.DotNet6, lang2.DotNet7:
		lang = &lang2.DotNet{}
	case lang2.Go120, lang2.Go116, lang2.Go119, lang2.Go118, lang2.Go117:
		lang = &lang2.Golang{}
	}
	spec := &kube.PodSpecConfig{
		Namespace:  node.K8sWorkerSpace,
		TaskName:   plan.Name,
		SourceCode: plan.SourceCode.GitUrl,
		Sdk:        lang.Sdk(plan.Env),
		Steps: func() []kube.Step {
			kvp := lang.Steps(plan.Env, plan.PipelineSteps)
			steps := make([]kube.Step, len(kvp))
			i := 0
			for _, cmd := range kvp {
				steps[i] = kube.Step{Command: cmd}
			}
			return steps
		}(),
	}
	_ = spec

	return errors.New("not implemented")
	// TODO: kube.ApplyPod
}
