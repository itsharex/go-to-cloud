package scm

import (
	"context"
	scm2 "github.com/drone/go-scm/scm"
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
	"strings"
)

// ListBranches 列出代码分支
func ListBranches(sourceCodeId uint) ([]scm.Branch, error) {
	sourceCode, err := repositories.GetProjectSourceCodeById(sourceCodeId)
	if err != nil {
		return nil, err
	}
	if client, err := newClient(scm.Type(sourceCode.CodeRepo.ScmOrigin), false, &sourceCode.CodeRepo.Url, &sourceCode.CodeRepo.AccessToken); err != nil {
		return nil, err
	} else {
		repo := strings.TrimLeft(
			strings.TrimRight(sourceCode.GitUrl, ".git"),
			sourceCode.CodeRepo.Url,
		)
		branches, _, err := client.Git.ListBranches(
			context.Background(),
			repo,
			scm2.ListOptions{
				Page: 0,
				Size: 10000,
			})

		rlt := make([]scm.Branch, len(branches))
		for i, branch := range branches {
			rlt[i] = scm.Branch{
				Reference: branch,
			}
		}
		if err != nil {
			return nil, err
		}

		return rlt, nil
	}
}
