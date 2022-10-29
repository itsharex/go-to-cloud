package project

type GitSources struct {
	Name string `json:"label"`
	Url  string `json:"value"`
}
type CodeRepoGroup struct {
	Id   uint         `json:"id"`
	Name string       `json:"label"`
	Git  []GitSources `json:"options"`
}
