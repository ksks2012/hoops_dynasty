package system_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hoops_dynasty/ecs/ecs_systems"
	"github.com/hoops_dynasty/ecs/entities"
	"github.com/hoops_dynasty/global"
	"github.com/hoops_dynasty/pkg/setting"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSaveUnitsState(t *testing.T) {
	// Mock global settings
	global.SaveLoadSetting = &setting.SaveLoadSettingS{
		SavePath:    "./testing/var",
		SaveFileExt: ".json",
	}

	// Create test data
	player1 := &entities.Player{ID: 1, Components: map[string]interface{}{"basic": "test1"}}
	player2 := &entities.Player{ID: 2, Components: map[string]interface{}{"basic": "test2"}}
	units := []*entities.Player{player1, player2}

	// Define test filename
	filename := "test_units"

	// Call SaveUnitsState
	err := ecs_systems.SaveUnitsState(units, filename)
	assert.NoError(t, err)

	// Verify file exists
	filePath := global.SaveLoadSetting.SavePath + "/" + filename + "sv" + global.SaveLoadSetting.SaveFileExt
	_, err = os.Stat(filePath)
	assert.NoError(t, err)

	// Verify file content
	file, err := os.Open(filePath)
	assert.NoError(t, err)
	defer file.Close()

	var savedUnits []*entities.Player
	err = json.NewDecoder(file).Decode(&savedUnits)
	assert.NoError(t, err)
	assert.Equal(t, units, savedUnits)

	// Cleanup
	err = os.RemoveAll(global.SaveLoadSetting.SavePath)
	assert.NoError(t, err)
}

func TestSaveUnitsState_YAML(t *testing.T) {
	// Mock global settings
	global.SaveLoadSetting = &setting.SaveLoadSettingS{
		SavePath:    "./testing/var",
		SaveFileExt: ".yaml",
	}

	// Create test data
	player1 := &entities.Player{ID: 1, Components: map[string]interface{}{"basic": "test1"}}
	player2 := &entities.Player{ID: 2, Components: map[string]interface{}{"basic": "test2"}}
	units := []*entities.Player{player1, player2}

	// Define test filename
	filename := "test_units"

	// Call SaveUnitsState
	err := ecs_systems.SaveUnitsState(units, filename)
	assert.NoError(t, err)

	// Verify file exists
	filePath := global.SaveLoadSetting.SavePath + "/" + filename + "sv" + global.SaveLoadSetting.SaveFileExt
	_, err = os.Stat(filePath)
	assert.NoError(t, err)

	// Verify file content
	file, err := os.Open(filePath)
	assert.NoError(t, err)
	defer file.Close()

	var savedUnits []*entities.Player
	err = yaml.NewDecoder(file).Decode(&savedUnits)
	assert.NoError(t, err)
	assert.Equal(t, units, savedUnits)

	// Cleanup
	err = os.RemoveAll(global.SaveLoadSetting.SavePath)
	assert.NoError(t, err)
}

func TestSaveUnitsState_EmptyFilename(t *testing.T) {
	// Mock global settings
	global.SaveLoadSetting = &setting.SaveLoadSettingS{
		SavePath:    "./testing/var",
		SaveFileExt: ".json",
	}

	// Create test data
	player1 := &entities.Player{ID: 1, Components: map[string]interface{}{"basic": "test1"}}
	units := []*entities.Player{player1}

	// Call SaveUnitsState with empty filename
	err := ecs_systems.SaveUnitsState(units, "")
	assert.Error(t, err)
	assert.Equal(t, "filename cannot be empty", err.Error())
}
