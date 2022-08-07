package pipeline

// NewRunner 创建一个运行器
// 运行器以预置的一组容器作为POD在k8s中运行
// 负责拉取代码、执行脚本并制作镜像
func NewRunner() {
	//
}

// IntegrationPod 用于持续集成的POD
// 通常包括：
//
//  1. go-to-cloud-agent（用于下载源码）
//  2. 语言相关的SDK环境容器
//  3. 用于构建镜像的Kaniko
type IntegrationPod struct {
	Containers []Container // 用于编译环境的基础容器，
	Dockerfile string      // 源码中用于打包镜像的Dockerfile
	Registry   string      // 镜像仓库
	ImageName  string      // 镜像名称
	ImageTag   string      // 镜像Tag
	Workspace  string      // 工作目录，由git下载下来的源码存放的路径
}

// Container 编译环境容器
type Container struct {
	Image string // 镜像
	Tag   string // 版本
	Alias string // 容器别名
}

const integrationPodTemplate string = `
apiVersion: "v1"
kind: "Pod"
metadata:
  name: "go-to-cloud-agent"
  labels:
    go-to-cloud-label: "go-to-cloud-agent"
    go-to-cloud: "agent"
  namespace: "go-to-cloud"
spec:
  containers:
{{- if .Containers}}
{{- range .Containers}}
  - command:
    - "cat"
    image: "{{.Image}}:{{.Tag}}"
    name: "{{.Alias}}"
    imagePullPolicy: IfNotPresent
    tty: true
{{- end}}	
{{- end}}	
  - name: "kaniko"
    image: "fanhousanbu-docker.pkg.coding.net/gotocloud-artifacts/kaniko/kaniko:v1.9.0-debug"
    args:
    - "--dockerfile={{.Dockerfile}}"
    - "--context={{.Workspace}}"
    - "--destination={{.Registry}}/{{.ImageName}}:{{.ImageTag}}"
    tty: true 
`
