package migrations

import (
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
)

type Migration20220831 struct {
}

func (m *Migration20220831) Up(db *gorm.DB) {

	userOrgRelNotExists := false

	if !db.Migrator().HasTable(&repo.User{}) {
		db.AutoMigrate(&repo.User{})
		userOrgRelNotExists = true
	}
	if !db.Migrator().HasTable(&repo.Org{}) {
		db.AutoMigrate(&repo.Org{})
		userOrgRelNotExists = true
	}

	if userOrgRelNotExists {
		org := &repo.Org{
			Name: "ROOT",
		}
		db.Debug().Create(org)

		user := &repo.User{
			Account:  "root",
			RealName: "系统管理员",
			Shortcut: "xtgly",
			Orgs:     []*repo.Org{org},
		}
		initRootPassword := "root"
		user.SetPassword(&initRootPassword)

		db.Debug().Create(user)
		db.Debug().Save(user)
	}
}

func (m *Migration20220831) Down(db *gorm.DB) {
	db.Migrator().DropTable(
		&repo.Org{},
		&repo.User{},
	)

	db.Migrator().DropTable("orgs_users_rel")
}
