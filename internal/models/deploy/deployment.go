package deploy

type Deployment struct {
	ProjectId               uint   `json:"projectId"`
	K8sNamespace            string `json:"k8SNamespace"`
	K8sRepoId               uint   `json:"k8SRepoId"`
	ArtifactDockerImageId   uint   `json:"artifactDockerImageId"`
	Ports                   string `json:"ports"`
	Cpus                    uint   `json:"cpus"`
	Env                     string `json:"env"`
	Replicas                uint   `json:"replicas"`
	Liveness                string `json:"liveness"`
	Readiness               string `json:"readiness"`
	RollingMaxSurge         uint   `json:"rollingMaxSurge"`
	RollingMaxUnavailable   uint   `json:"rollingMaxUnavailable"`
	ResourceLimitCpuRequest uint   `json:"resourceLimitCpuRequest"`
	ResourceLimitCpuLimits  uint   `json:"resourceLimitCpuLimits"`
	ResourceLimitMemRequest uint   `json:"resourceLimitMemRequest"`
	ResourceLimitMemLimits  uint   `json:"resourceLimitMemLimits"`
	NodeSelector            string `json:"nodeSelector"`
}

type Base struct {
	Id uint `json:"id,omitempty"` // deployment id
}

type ScalePods struct {
	Base
	Num uint `json:"num,omitempty" json:"num,omitempty"`
}

type RestartPods struct {
	Base
}

type Redeployment struct {
	Base
}
