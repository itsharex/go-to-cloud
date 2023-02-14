package builders

import (
	"bytes"
	"go-to-cloud/internal/repositories"
	"strings"
	"text/template"
)

const (
	DotNet3 = "dot-net-3.1"
	DotNet5 = "dot-net-5"
	DotNet6 = "dot-net-6"
	DotNet7 = "dot-net-7"

	Go116 = "go-1.16"
	Go117 = "go-1.17"
	Go118 = "go-1.18"
	Go119 = "go-1.19"
	Go120 = "go-1.20"
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
	spec := &PodSpecConfig{
		TaskName:   plan.Name,
		SourceCode: plan.SourceCode.GitUrl,
		Sdk: func() string {
			const dotnet = "mcr.microsoft.com/dotnet/sdk"
			const golang = ""
			switch plan.Env {
			case DotNet3, DotNet7, DotNet6, DotNet5:
				{
					switch plan.Env {
					case DotNet6:
						return dotnet + ":6.0.0"
					case DotNet5:
						return dotnet + ":5.0.0"
					case DotNet7:
						return dotnet + "7.0.0"
					case DotNet3:
						return dotnet + "3.1.0"
					}
				}
			case Go118, Go119, Go117, Go116, Go120:
				{
					switch plan.Env {
					case Go118:
						return golang + ":1.18"
					case Go116:
						return golang + ":1.16"
					case Go117:
						return golang + ":1.17"
					case Go119:
						return golang + ":1.19"
					case Go120:
						return golang + ":1.20"
					}

				}
			}
			return ""
		}(),
		Steps: func(env string) []Step {
			panic("not implemented")
			return nil
		}(plan.Env),
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
