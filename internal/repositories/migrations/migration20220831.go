package migrations

import (
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
)

type Migration20220831 struct {
}

func (m *Migration20220831) Up(db *gorm.DB) {
	db.AutoMigrate(
		&repo.Infrastructure{},
		&repo.User{},
		&repo.Org{},
	)

	user := &repo.User{
		Account:     "aaa",
		Password:    "bbb",
		Email:       "cccc",
		Mobile:      "133",
		LastLoginAt: nil,
		Orgs: []*repo.Org{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
		},
	}

	db.Debug().Create(user)
	db.Debug().Save(user)
}

func (m *Migration20220831) Down(db *gorm.DB) {
	db.Migrator().DropTable(
		&repo.Org{},
		&repo.User{},
		&repo.Infrastructure{},
	)

	db.Migrator().DropTable("orgs_users_rel")
}
