package conf

import (
	"os"
	"strings"
)

type Env struct {
	Name     string // 	环境名称
	Debugger bool   // 是否调试模式
}

func (env *Env) IsDevelopment() bool {
	return strings.EqualFold("Development", env.Name)
}

func (env *Env) IsProduction() bool {
	return strings.EqualFold("Production", env.Name)
}

func (env *Env) GetEnvName() *string {
	return &env.Name
}

var Environment *Env

func init() {
	envName := "prod"
	if len(os.Getenv("Env")) > 0 {
		envName = os.Getenv("Env")
	}
	Environment = &Env{
		Name:     envName,
		Debugger: false,
	}
}
