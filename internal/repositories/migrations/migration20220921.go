package migrations

import (
	"go-to-cloud/internal/middlewares"
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
)

type migration20220921 struct {
}

func (m *migration20220921) Up(db *gorm.DB) {

	if !db.Migrator().HasTable(&repo.CasbinRule{}) {
		err := db.AutoMigrate(&repo.CasbinRule{})
		if err != nil {
			panic(err)
		} else {
			if enforce, err := middlewares.GetCasbinEnforcer(); err == nil {
				enforce.AddGroupingPolicy("root", "*")
				enforce.AddGroupingPolicy("root", "guest")
				enforce.AddPolicies([][]string{{"guest", "/api/monitor/{k8s}/apps/query", "GET"}})
				enforce.AddPolicies([][]string{{"guest", "/api/user/info", "GET"}})
				enforce.AddPolicies([][]string{{"guest", "/org/list", "GET"}})
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
