package project

type GitSources struct {
	Id        string `json:"id"`
	Name      string `json:"label"`
	Namespace string `json:"namespace"`
	Url       string `json:"value"`
}
type CodeRepoGroup struct {
	Id   uint         `json:"id"`
	Name string       `json:"label"`
	Git  []GitSources `json:"options"`
}
