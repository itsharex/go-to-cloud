package conf

import "sync"

var kaniko *string

var onceBuild sync.Once

func GetBuildImage() *string {
	if kaniko == nil {
		onceBuild.Do(func() {
			if kaniko == nil {
				j := getConf().Builder
				kaniko = &j.Kaniko
			}
		})
	}
	return kaniko
}
