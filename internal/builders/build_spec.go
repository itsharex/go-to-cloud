package builders

import (
	"bytes"
	lang2 "go-to-cloud/internal/builders/lang"
	"go-to-cloud/internal/repositories"
	"strings"
	"text/template"
)

type Step struct {
	Command string
} // 构建步骤

type PodSpecConfig struct {
	TaskName   string // pod name
	SourceCode string // git url
	Sdk        string // sdk 基础镜像
	Steps      []Step
}

// BuildPodSpec 创建构建模板 k8s pod spec
func BuildPodSpec(plan *repositories.Pipeline) {
	var lang lang2.Tpl
	switch plan.Env {
	case lang2.DotNet3, lang2.DotNet5, lang2.DotNet6, lang2.DotNet7:
		lang = &lang2.DotNet{}
	case lang2.Go120, lang2.Go116, lang2.Go119, lang2.Go118, lang2.Go117:
		lang = &lang2.Golang{}
	}
	spec := &PodSpecConfig{
		TaskName:   plan.Name,
		SourceCode: plan.SourceCode.GitUrl,
		Sdk:        lang.Sdk(plan.Env),
		Steps: func() []Step {
			kvp := lang.Steps(plan.Env, plan.PipelineSteps)
			steps := make([]Step, len(kvp))
			i := 0
			for _, cmd := range kvp {
				steps[i] = Step{Command: cmd}
			}
			return steps
		}(),
	}

	_, _ = makeTemplate(spec)
}

func makeTemplate(spec *PodSpecConfig) (*string, error) {
	tpl, err := template.New("k8s").Parse(BuildTemplate)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, *spec)
	if err != nil {
		return nil, err
	}

	yml := strings.TrimSpace(buf.String())

	return &yml, nil
}

const BuildTemplate string = `
apiVersion: v1
kind: Pod
metadata:
  name: {{.TaskName}}
  labels:
    builder: gotocloud-builder
spec:
    initContainers:
    - name: coderepo
      image: alpine/git
      imagePullPolicy: IfNotPresent
      command: 
      - /bin/sh
      - -ec
      - |
        git clone {{.SourceCode}} /git
      volumeMounts:
      - name: workdir
        mountPath: "/git"
    containers:
    - name: compile
      image: {{.Sdk}}
      imagePullPolicy: IfNotPresent
      command:
      - /bin/sh
      - -ec
      - |
        cd /workdir
        dotnet build
{{- range .Steps}}
        {{.Command}}
{{- end}}
      volumeMounts:
      - name: workdir
        mountPath: "/workdir"
    restartPolicy: OnFailure
    volumes:
    - name: workdir
      emptyDir: {}
`
