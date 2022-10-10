package artifact

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

type Artifact struct {
	Testing
	Name      string    `json:"name" form:"name"`
	Orgs      []uint    `json:"orgs" form:"orgs"`
	OrgLites  []OrgLite `json:"orgLites"`
	Remark    string    `json:"remark"`
	UpdatedAt string    `json:"updatedAt"`
}

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}
