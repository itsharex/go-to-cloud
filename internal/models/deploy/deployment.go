package deploy

import "go-to-cloud/internal/repositories"

type Deployment struct {
	repositories.Deployment
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
