package artifact

import "go-to-cloud/internal/utils"

type Type int

const (
	Docker Type = iota
	OSS
	Nuget
	Maven
	Npm
	S3
)

type Testing struct {
	Id         uint   `json:"id"`
	Type       Type   `json:"type"`
	IsSecurity bool   `json:"isSecurity"`
	Url        string `json:"url"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

type Image struct {
	Id             uint           `json:"id"`
	Name           string         `json:"name"`
	LatestVersion  string         `json:"latestVersion"`
	PublishedAt    utils.JsonTime `json:"publishedAt"`
	PublishCounter int            `json:"publishCounter"`
}

type Artifact struct {
	Testing
	Name      string    `json:"name" form:"name"`
	Orgs      []uint    `json:"orgs" form:"orgs"`
	OrgLites  []OrgLite `json:"orgLites"`
	Remark    string    `json:"remark"`
	UpdatedAt string    `json:"updatedAt"`
	Items     []Image   `json:"items"`
}

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type Query struct {
	Artifact
}
