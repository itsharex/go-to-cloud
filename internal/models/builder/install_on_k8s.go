package builder

type OnK8sModel struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	MaxWorker  int    `json:"maxWorker"`
	Workspace  string `json:"workspace"`
	KubeConfig string `json:"kubeConfig"`
	Orgs       []uint `json:"orgs"`
}
