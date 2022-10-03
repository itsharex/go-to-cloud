package scm

import (
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/gitea"
	"github.com/drone/go-scm/scm/driver/gitee"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/transport"
	"go-to-cloud/internal/models"
	"net/http"
	"strings"
)

// newClient 获取scm客户端
func newClient(origin models.ScmType, isPublic bool, uri, token *string) (client *scm.Client, err error) {
	switch origin {
	case models.Github:
		client = github.NewDefault()
		break
	case models.Gitlab:
		client, err = gitlab.New(*uri)
		break
	case models.Gitee:
		client = gitee.NewDefault()
		break
	case models.Gitea:
		client, err = gitea.New(*uri)
	}

	if client != nil {
		client.Client = scmHttpClient(origin, isPublic, token)
	}
	return
}

func scmHttpClient(origin models.ScmType, isPublic bool, token *string) *http.Client {
	if isPublic || token == nil || len(strings.TrimSpace(*token)) == 0 {
		return &http.Client{}
	}

	switch origin {
	case models.Gitlab:
		return &http.Client{
			Transport: &transport.PrivateToken{Token: *token},
		}
	default:
		return &http.Client{
			Transport: &transport.BearerToken{Token: *token},
		}
	}
}
