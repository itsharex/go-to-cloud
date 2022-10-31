package project

import (
	"go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/scm"
)

func GetCodeRepoGroupsByOrg(orgId []uint) ([]project.CodeRepoGroup, error) {

	coderepo, err := scm.List(orgId, nil)

	if err != nil {
		return nil, err
	}

	rlt := make([]project.CodeRepoGroup, len(coderepo))
	for i, s := range coderepo {
		if models, err := scm.ListCodeProjects(s.Origin, &s.Url, s.Token); err != nil {
			return nil, err
		} else {
			rlt[i].Id = s.Id
			rlt[i].Name = s.Name
			rlt[i].Host = s.Url
			rlt[i].Git = make([]project.GitSources, len(models))
			for j, model := range models {
				rlt[i].Git[j] = project.GitSources{
					Id:        model.Id,
					Name:      model.Name,
					Url:       model.Url,
					Namespace: model.Namespace,
				}
			}
		}
	}
	return rlt, nil
}
