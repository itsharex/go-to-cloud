package conf

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"sync"
)

var once sync.Once

type Conf struct {
	Agent struct {
		Image string
	}
	Db struct {
		User     string
		Password string
		Host     string
		Schema   string
	}
	Jwt struct {
		Security string // 私钥
		Realm    string // Realm
		IdKey    string // IdentityKey
	}
	Kind []string // 用户分类
}

var conf *Conf

func getConf() *Conf {
	if conf == nil {
		once.Do(func() {
			if conf == nil {
				filePath := getConfFilePath()
				conf = getConfiguration(filePath)
			}
		})
	}
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
