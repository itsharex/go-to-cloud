package deploy

type Deployment struct {
	K8S             uint   `json:"k8s"`
	Namespace       string `json:"namespace"`
	Artifact        uint   `json:"artifact"`
	Version         string `json:"version"`
	Replicate       uint   `json:"replicate"`
	Healthcheck     string `json:"healthcheck"`
	HealthcheckPort uint   `json:"healthcheckPort"`
	EnableLimit     bool   `json:"enableLimit"`
	CpuLimits       uint   `json:"cpuLimits"`
	CpuRequest      uint   `json:"cpuRequest"`
	MemLimits       uint   `json:"memLimits"`
	MemRequest      uint   `json:"memRequest"`
	Ports           []struct {
		ServicePort   string `json:"text"`
		ContainerPort string `json:"value"`
	} `json:"ports"`
	Env []struct {
		VarName  string `json:"varName"`
		VarValue string `json:"varValue"`
	} `json:"env"`
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
