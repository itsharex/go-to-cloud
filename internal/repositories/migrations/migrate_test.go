package migrations

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/conf"
	"gorm.io/gorm"
	"testing"
)

var sortedUp, sortedDown []int

type migration001 struct {
}

func (m *migration001) Up(_ *gorm.DB) {
	println("[up]001")
	sortedUp = append(sortedUp, 1)
}

func (m *migration001) Down(_ *gorm.DB) {
	println("[down]001")
	sortedDown = append(sortedDown, 1)
}

type migration002 struct {
}

func (m *migration002) Up(_ *gorm.DB) {
	println("[up]002")
	sortedUp = append(sortedUp, 2)
}

func (m *migration002) Down(_ *gorm.DB) {
	println("[down]002")
	sortedDown = append(sortedDown, 2)
}

func initTestData() {
	sortedUp = make([]int, 0)
	sortedDown = make([]int, 0)

	migrations = []Migration{
		&migration001{},
		&migration002{},
	}
}
func TestMigrate(t *testing.T) {
	var db *gorm.DB

	if testing.Short() {
		initTestData()
	} else {
		db = conf.GetDbClient()
	}

	Migrate(db)

	if testing.Short() {
		assert.Equal(t, 1, sortedUp[0])
		assert.Equal(t, 2, sortedUp[1])
	}
}

func TestRollback(t *testing.T) {
	var db *gorm.DB

	if testing.Short() {
		initTestData()
	} else {
		db = conf.GetDbClient()
	}

	Rollback(db)

	if testing.Short() {
		assert.Equal(t, 2, sortedDown[0])
		assert.Equal(t, 1, sortedDown[1])
	}
}
