package migrations

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

var sortedUp, sortedDown []int

type migration001 struct {
}

func (m *migration001) Up(db *gorm.DB) {
	println("[up]001")
	sortedUp = append(sortedUp, 1)
}

func (m *migration001) Down(db *gorm.DB) {
	println("[down]001")
	sortedDown = append(sortedDown, 1)
}

type migration002 struct {
}

func (m *migration002) Up(db *gorm.DB) {
	println("[up]002")
	sortedUp = append(sortedUp, 2)
}

func (m *migration002) Down(db *gorm.DB) {
	println("[down]002")
	sortedDown = append(sortedDown, 2)
}

func init() {
	sortedUp = make([]int, 0)
	sortedDown = make([]int, 0)

	migrations = []Migration{
		&migration001{},
		&migration002{},
	}
}
func TestMigrate(t *testing.T) {
	if testing.Short() {
		t.Skip("debugger only")
	}

	Migrate()

	assert.Equal(t, 1, sortedUp[0])
	assert.Equal(t, 2, sortedUp[1])
}

func TestRollback(t *testing.T) {
	if testing.Short() {
		t.Skip("debugger only")
	}

	Rollback()

	assert.Equal(t, 2, sortedDown[0])
	assert.Equal(t, 1, sortedDown[1])
}
