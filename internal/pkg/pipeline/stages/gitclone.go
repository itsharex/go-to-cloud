package stages

import (
	"errors"
	"fmt"
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/gitea"
	"github.com/drone/go-scm/scm/driver/gitee"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/driver/gogs"
	"github.com/drone/go-scm/scm/transport"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gohttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"net/http"
	"net/url"
	"strings"
)

type GitScmType int8

const (
	Github GitScmType = iota
	Gitee  GitScmType = iota
	Gitea  GitScmType = iota
	Gitlab GitScmType = iota
	Gogs   GitScmType = iota
)

type GitCloneStage struct {
	Token   string
	GitUrl  string
	Branch  string
	WorkDir string
}

// NewGitClient 创建git客户端对象
func NewGitClient(scmType GitScmType, gitUrl string, token *string) (*scm.Client, error) {
	if strings.HasSuffix(gitUrl, ".git") {
		gitUrl = strings.TrimSuffix(gitUrl, ".git")
	}

	u, err := url.Parse(gitUrl)
	if err != nil {
		return nil, err
	}
	if !strings.EqualFold("http", u.Scheme) && !strings.EqualFold("https", u.Scheme) {
		return nil, errors.New("仅支持http或https的git地址")
	}

	notSupportedScmError := fmt.Errorf("当前仅支持github、gitlab、gitee、gitea、gogs")
	var scmClient *scm.Client
	switch scmType {
	case Gitea, Gitlab, Gogs:
		host := u.Host
		if len(u.Port()) > 0 {
			host = fmt.Sprintf("%s:%s", host, u.Port())
		}
		host = fmt.Sprintf("%s://%s", u.Scheme, host)

		if Gitea == scmType {
			scmClient, err = gitea.New(host)
		} else if Gitlab == scmType {
			scmClient, err = gitlab.New(host)
		} else {
			scmClient, err = gogs.New(host)
		}
	case Github:
		scmClient = github.NewDefault()
	case Gitee:
		scmClient = gitee.NewDefault()
	default:
		err = notSupportedScmError
	}
	if scmClient != nil {
		httpClient := gitHttpClient(scmType, token)
		if httpClient != nil {
			scmClient.Client = httpClient
		} else {
			return nil, notSupportedScmError
		}
	}
	return scmClient, err
}

func gitHttpClient(scmType GitScmType, token *string) *http.Client {
	if token == nil || len(*token) == 0 {
		return &http.Client{}
	}
	switch scmType {
	case Gitlab, Gogs:
		return &http.Client{
			Transport: &transport.PrivateToken{
				Token: *token,
			}}
	case Gitee, Gitea, Github:
		return &http.Client{
			Transport: &transport.BearerToken{
				Token: *token,
			},
		}
	default:
		return nil
	}
}

func (m *GitCloneStage) Stub() error {

	// TODO: master调用
	return errors.New("NOT Implemented")
}

func (m *GitCloneStage) Run() error {
	var auth *gohttp.BasicAuth

	if m == nil || len(m.Token) == 0 {
		auth = nil
	} else {
		auth = &gohttp.BasicAuth{
			Username: "go-to-cloud",
			Password: m.Token,
		}
	}

	opt := &git.CloneOptions{
		URL:  m.GitUrl,
		Auth: auth,
	}
	if len(m.Branch) > 0 {
		opt.ReferenceName = plumbing.NewBranchReferenceName(m.Branch)
	}
	var err error
	if len(m.WorkDir) == 0 {
		_, err = git.Clone(memory.NewStorage(), nil, opt)
	} else {
		_, err = git.PlainClone(m.WorkDir, false, opt)
	}
	return err
}
