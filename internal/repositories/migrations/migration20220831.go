package migrations

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	repo "go-to-cloud/internal/repositories"
	"gorm.io/datatypes"
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
			Account:  models.RootUserName,
			RealName: "系统管理员",
			Kind: func() datatypes.JSON {
				s, _ := json.Marshal([]string{conf.KindRoot})
				return s
			}(),
			Pinyin:     "xitongguanliyuan",
			PinyinInit: "xtgly",
			Orgs:       []*repo.Org{org},
		}
		initRootPassword := "root"
		user.SetPassword(&initRootPassword)
		db.Debug().Create(user)
		db.Debug().Save(user)

		guest := &repo.User{
			Account:  "guest",
			RealName: "游客", Kind: func() datatypes.JSON {
				s, _ := json.Marshal([]string{conf.KindGuest})
				return s
			}(),
			Pinyin:     "youke",
			PinyinInit: "yk",
			Orgs:       []*repo.Org{org},
		}
		initRootPassword2 := "guest"
		guest.SetPassword(&initRootPassword2)
		db.Debug().Create(guest)
		db.Debug().Save(guest)
	}
}

func (m *Migration20220831) Down(db *gorm.DB) {
	db.Migrator().DropTable(
		&repo.Org{},
		&repo.User{},
	)

	db.Migrator().DropTable("orgs_users_rel")
}
