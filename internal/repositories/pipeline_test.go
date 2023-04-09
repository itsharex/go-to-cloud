package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/utils"
	"os"
	"testing"
)

func prepareDb() {
	if conf.GetDbClient().Migrator().HasTable(Org{}) {
		conf.GetDbClient().Migrator().DropTable(&Org{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&Org{})

	if conf.GetDbClient().Migrator().HasTable(User{}) {
		conf.GetDbClient().Migrator().DropTable(&User{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&User{})

	if conf.GetDbClient().Migrator().HasTable(&PipelineSteps{}) {
		conf.GetDbClient().Migrator().DropTable(&PipelineSteps{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&PipelineSteps{})

	if conf.GetDbClient().Migrator().HasTable(&Pipeline{}) {
		conf.GetDbClient().Migrator().DropTable(&Pipeline{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&Pipeline{})
}

func TestNewPlan(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &pipeline.PlanModel{
		Name: *utils.StrongPasswordGen(6),
	}

	plan, err := NewPlan(1, 1, model)
	assert.NoError(t, err)
	assert.NotNil(t, plan)

	plan2, err := QueryPipeline(plan.ID)
	assert.NoError(t, err)
	assert.Equal(t, plan2.Name, plan.Name)
}
