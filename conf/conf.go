package conf

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"sync"
)

var once sync.Once

type Conf struct {
	Builder struct {
		Kaniko string
	}
	Db struct {
		User     string
		Password string
		Host     string
		Schema   string
	}
	Jwt  JWT
	Kind []string // 用户分类
}

var conf *Conf

func getConf() *Conf {
	once.Do(func() {
		if conf == nil {
			filePath := getConfFilePath()
			conf = getConfiguration(filePath)
		}
	})
	return conf
}

// getConfiguration 读取配置
func getConfiguration(filePath *string) *Conf {
	if file, err := os.ReadFile(*filePath); err != nil {
		panic(err)
	} else {
		c := Conf{}
		err := yaml.Unmarshal(file, &c)
		if err != nil {
			panic(err)
		}

		return &c
	}
}
