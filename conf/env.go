package conf

import "strings"

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

var Enviroment *Env

func init() {
	Enviroment = &Env{
		Name:     "Development",
		Debugger: false,
	}
}
