package kube

import (
	"bytes"
	"fmt"
	"go-to-cloud/conf"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"strings"
	"text/template"
)

type Step struct {
	Command string
} // 构建步骤

type PodSpecConfig struct {
	Namespace  string
	TaskName   string // pod name
	SourceCode string // git url
	Sdk        string // sdk 基础镜像
	Steps      []Step
}

// Build 构建任务
func (client *Client) Build(podSpecConfig *PodSpecConfig) error {
	spec, err := makeTemplate(podSpecConfig)

	if err != nil {
		return err
	}

	if conf.Environment.IsDevelopment() {
		fmt.Println(*spec)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	podCfg := corev1.PodApplyConfiguration{}

	if err := DecodeYaml(spec, &podCfg); err != nil {
		fmt.Println(err)
		return err
	}
	if _, e := client.ApplyPod(&podSpecConfig.Namespace, &podCfg); e != nil {
		fmt.Println(e)
		return e
	}

	return nil
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
