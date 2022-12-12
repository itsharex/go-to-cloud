package builder

import "go-to-cloud/internal/models"

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type NodesOnK8s struct {
	Orgs     []uint    `json:"orgs" form:"orgs"`
	OrgLites []OrgLite `json:"orgLites"`
	Name     string    `json:"name" form:"name"`
}

type Query struct {
	models.Pager
	NodesOnK8s
}
