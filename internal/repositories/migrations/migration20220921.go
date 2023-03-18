package migrations

import (
	"github.com/casbin/casbin/v2"
	"go-to-cloud/internal/auth"
	"go-to-cloud/internal/middlewares"
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
	"regexp"
)

type migration20220921 struct {
}

// addGroupPolicy 添加角色继承关系
func addGroupPolicy(enforce *casbin.Enforcer) {
	for _, strings := range auth.GroupPolicies() {
		enforce.AddGroupingPolicy(strings[0], strings[1])
	}
}

// addResourcePolicy 添加权限点
func addResourcePolicy(enforce *casbin.Enforcer) {
	for _, p := range auth.ResourcePolicies() {
		enforce.AddPolicies(p)
	}
}

// addRouterPolicy 添加路由权限
func addRouterPolicy(enforce *casbin.Enforcer) {
	reg := regexp.MustCompile(`:(\w+)`)
	for _, routerMap := range auth.RouterMaps {
		for _, kind := range routerMap.Kinds {
			for _, method := range routerMap.Methods {
				// 需要将路由参数 :params 替换为 {params}来适配keyMatch4匹配算法
				enforce.AddPolicies([][]string{{string(kind), reg.ReplaceAllString(routerMap.Url, "{$1}"), string(method)}})
			}
		}
	}
}

func (m *migration20220921) Up(db *gorm.DB) {

	if !db.Migrator().HasTable(&repo.CasbinRule{}) {
		err := db.AutoMigrate(&repo.CasbinRule{})
		if err != nil {
			panic(err)
		} else {
			if enforce, err := middlewares.GetCasbinEnforcer(); err == nil {
				addGroupPolicy(enforce)
				addResourcePolicy(enforce)
				addRouterPolicy(enforce)
			}
		}
	}
}

func (m *migration20220921) Down(db *gorm.DB) {
	err := db.Migrator().DropTable(
		&repo.CasbinRule{},
	)
	if err != nil {
		panic(err)
	}
}
