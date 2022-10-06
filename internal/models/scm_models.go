package models

type ScmType int

const (
	Gitlab ScmType = iota
	Github
	Gitee
	Gitea
)

type ScmTesting struct {
	Id       uint    `json:"id"`
	Origin   ScmType `json:"origin"`
	IsPublic bool    `json:"isPublic"`
	Url      string  `json:"url"`
	Token    *string `json:"token"`
}

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type Scm struct {
	ScmTesting
	Name      string    `json:"name" form:"name"`
	Orgs      []uint    `json:"orgs" form:"orgs"`
	OrgLites  []OrgLite `json:"orgLites"`
	Remark    string    `json:"remark"`
	UpdatedAt string    `json:"updatedAt"`
}

type ScmQuery struct {
	Pager
	Scm
}
