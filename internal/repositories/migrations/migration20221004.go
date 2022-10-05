package migrations

import (
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
)

type migration20221004 struct {
}

func (m *migration20221004) Up(db *gorm.DB) {

	if !db.Migrator().HasTable(&repo.CodeRepo{}) {
		err := db.AutoMigrate(&repo.CodeRepo{})
		if err != nil {
			panic(err)
		}
	}
}

func (m *migration20221004) Down(db *gorm.DB) {
	err := db.Migrator().DropTable(
		&repo.CodeRepo{},
	)
	if err != nil {
		panic(err)
	}
}
