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

	if !db.Migrator().HasTable(&repo.ArtifactRepo{}) {
		err := db.AutoMigrate(&repo.ArtifactRepo{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&repo.ArtifactDockerImages{}) {
		err := db.AutoMigrate(&repo.ArtifactDockerImages{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&repo.K8sRepo{}) {
		err := db.AutoMigrate(&repo.K8sRepo{})
		if err != nil {
			panic(err)
		}
	}

	//if !db.Migrator().HasTable(&repo.GitRepo{}) {
	//	err := db.AutoMigrate(&repo.GitRepo{})
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	if !db.Migrator().HasTable(&repo.Project{}) {
		err := db.AutoMigrate(&repo.Project{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&repo.ProjectSourceCode{}) {
		err := db.AutoMigrate(&repo.ProjectSourceCode{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&repo.Pipeline{}) {
		err := db.AutoMigrate(&repo.Pipeline{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&repo.PipelineSteps{}) {
		err := db.AutoMigrate(&repo.PipelineSteps{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&repo.BuilderNode{}) {
		err := db.AutoMigrate(&repo.BuilderNode{})
		if err != nil {
			panic(err)
		}
	}
}

func (m *migration20221004) Down(db *gorm.DB) {
	err := db.Migrator().DropTable(
		&repo.CodeRepo{},
		&repo.ArtifactRepo{},
		&repo.ArtifactDockerImages{},
		&repo.K8sRepo{},
		//&repo.GitRepo{},
		&repo.Project{},
		&repo.ProjectSourceCode{},
		&repo.Pipeline{},
		&repo.PipelineSteps{},
		&repo.BuilderNode{},
	)
	if err != nil {
		panic(err)
	}
}
