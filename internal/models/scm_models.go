package models

type ScmType int

const (
	Gitlab ScmType = iota
	Github
	Gitee
	Gitea
)

type ScmTesting struct {
	Origin   ScmType `json:"origin"`
	IsPublic bool    `json:"isPublic"`
	Url      string  `json:"url"`
	Token    *string `json:"token"`
}

type Scm struct {
	ScmTesting
	Name   string `json:"name"`
	Remark string `json:"remark"`
}
