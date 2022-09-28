package migrations

import (
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
		}
	}

	var cnt int64
	if err := db.Find(&repo.CasbinRule{}, 1).Count(&cnt).Error; cnt == 0 && err == nil {
		r := &repo.CasbinRule{
			Id:    1,
			PType: "g",
			V0:    "root",
			V1:    "*",
		}
		db.Debug().Create(r)
	}
	if err := db.Find(&repo.CasbinRule{}, 2).Count(&cnt).Error; cnt == 0 && err == nil {
		r := &repo.CasbinRule{
			Id:    2,
			PType: "g",
			V0:    "root",
			V1:    "guest",
		}
		db.Debug().Create(r)
	}
	if err := db.Find(&repo.CasbinRule{}, 3).Count(&cnt).Error; cnt == 0 && err == nil {
		r := &repo.CasbinRule{
			Id:    3,
			PType: "p",
			V0:    "guest",
			V1:    "/api/user/info",
			V2:    "GET",
		}
		db.Debug().Create(r)
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
