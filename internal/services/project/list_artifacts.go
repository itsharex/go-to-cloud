package project

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

var projectArtifacts *cache.Cache

// ListArtifacts 根据制品名称查询
func ListArtifacts(projectId uint, queryString *string) ([]string, error) {
	return nil, errors.New("not implemented")
	//if projectArtifacts.Get(projectId)
	//if repo, err := repositories.QueryImagesByProjectId(projectId); err != nil {
	//	return nil, err
	//} else {
	//
	//}
}

func init() {
	projectArtifacts = cache.New(time.Minute*3, 0)
}
