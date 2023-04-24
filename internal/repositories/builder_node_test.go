package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/builder"
	"gorm.io/datatypes"
	"os"
	"testing"
)

func TestGetBuildNodesById(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.NotNil(t, actualNode)

	assert.Equal(t, model.Name, actualNode.Name)
	assert.Equal(t, model.Remark, actualNode.Remark)
	assert.Equal(t, model.Workspace, actualNode.K8sWorkerSpace)
}

func TestDeleteBuilderNode(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.NotNil(t, actualNode)

	assert.NoError(t, DeleteBuilderNode(1, node))

	nilNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, nilNode.ID)
}

func TestUpdateBuilderNode(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	model.Id = node
	model.Name = "test2"
	model.MaxWorkers = 4
	err = UpdateBuilderNode(model, 1, []uint{3, 4})
	assert.NoError(t, err)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.Equal(t, "test", actualNode.Name)  // not changed
	assert.Equal(t, 4, actualNode.MaxWorkers) // changed
	assert.Equal(t, datatypes.JSON("[3,4]"), actualNode.BelongsTo)
}
