package migrations

import (
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
)

type Migration20220831 struct {
}

func (m *Migration20220831) Up(db *gorm.DB) {

	userOrgsRel := false
	if !db.Migrator().HasTable(&repo.Infrastructure{}) {
		db.AutoMigrate(&repo.Infrastructure{})
	}
	if !db.Migrator().HasTable(&repo.User{}) {
		db.AutoMigrate(&repo.User{})
		userOrgsRel = true
	}
	if !db.Migrator().HasTable(&repo.Org{}) {
		db.AutoMigrate(&repo.Org{})
		userOrgsRel = true
	}

	if userOrgsRel {
		orgs := []*repo.Org{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
		}
		db.Debug().Create(orgs)

		user := []repo.User{
			{
				Account:     "aaa",
				Password:    "bbb",
				Email:       "cccc",
				Mobile:      "133",
				LastLoginAt: nil,
				Orgs:        orgs,
			},
			{
				Account:     "123",
				Password:    "456",
				Email:       "8888",
				Mobile:      "abb",
				LastLoginAt: nil,
				Orgs:        orgs,
			},
		}

		db.Debug().Create(user)
		db.Debug().Save(user)
	}
}

func (m *Migration20220831) Down(db *gorm.DB) {
	db.Migrator().DropTable(
		&repo.Org{},
		&repo.User{},
		&repo.Infrastructure{},
	)

	db.Migrator().DropTable("orgs_users_rel")
}
