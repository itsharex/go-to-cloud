package artifact

type Type int

const (
	OSS Type = iota
	Docker
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

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}
