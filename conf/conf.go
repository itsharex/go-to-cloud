package conf

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
)

type Conf struct {
	Db struct {
		User     string `yaml:"user"`
		Password string `yaml:"pwd"`
		Host     string `yaml:"host"`
		Schema   string `yaml:"database"`
	}
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
