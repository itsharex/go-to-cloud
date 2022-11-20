package server

import (
	"context"
	gotocloud "go-to-cloud/internal/agent/proto"
	"go-to-cloud/internal/pkg/pipeline/stages"
	"go-to-cloud/internal/utils"
	"os"
)

func (s *Agent) GitClone(ctx context.Context, in *gotocloud.CloneRequest) (*gotocloud.CloneResponse, error) {
	if workdir, err := os.MkdirTemp("", "gtc"); err != nil {
		return nil, err
	} else {
		gitCloneStage := stages.GitCloneStage{
			Token:   *decodeToken(in),
			GitUrl:  in.Address,
			Branch:  in.Branch,
			WorkDir: workdir,
		}
		if err = gitCloneStage.Run(); err != nil {
			return nil, err
		} else {
			return &gotocloud.CloneResponse{Workdir: workdir}, nil
		}
	}
}

func decodeToken(req *gotocloud.CloneRequest) *string {
	decodedToken := string(utils.AesEny([]byte(req.EncodedToken)))
	return &decodedToken
}
