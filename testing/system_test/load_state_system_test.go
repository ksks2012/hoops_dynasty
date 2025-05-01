package system_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/hoops_dynasty/ecs/ecs_systems"
	"github.com/hoops_dynasty/ecs/entities"
	"github.com/hoops_dynasty/global"
	"github.com/hoops_dynasty/pkg/setting"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestLoadUnitsStateYAML(t *testing.T) {
	// Mock global settings
	global.SaveLoadSetting = &setting.SaveLoadSettingS{
		SavePath:    "testing",
		SaveFileExt: ".yaml",
	}

	// Create test data
	testFilename := "test_units"
	testFilePath := global.SaveLoadSetting.SavePath + "/" + testFilename + "sv" + global.SaveLoadSetting.SaveFileExt
	testUnits := []entities.Player{
		{ID: 1, Name: "Player1", Components: make(map[string]interface{})},
		{ID: 2, Name: "Player2", Components: make(map[string]interface{})},
	}

	// Create test directory
	err := os.MkdirAll(global.SaveLoadSetting.SavePath, os.ModePerm)
	assert.NoError(t, err)
	defer os.RemoveAll(global.SaveLoadSetting.SavePath)

	// Write test data to file
	file, err := os.Create(testFilePath)
	assert.NoError(t, err)
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(testUnits)
	assert.NoError(t, err)
	encoder.Close()

	// Test successful load
	loadedUnits, err := ecs_systems.LoadUnitsState(testFilename)
	assert.NoError(t, err)
	assert.Equal(t, testUnits, loadedUnits)

	// Test file not found
	_, err = ecs_systems.LoadUnitsState("nonexistent_file")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, os.ErrNotExist))

	// Test invalid file format
	invalidFilePath := global.SaveLoadSetting.SavePath + "/invalid_file" + global.SaveLoadSetting.SaveFileExt
	invalidFileContent := []byte("invalid content")
	err = os.WriteFile(invalidFilePath, invalidFileContent, os.ModePerm)
	assert.NoError(t, err)

	_, err = ecs_systems.LoadUnitsState("invalid_file")
	assert.Error(t, err)
}

func TestLoadUnitsStateJSON(t *testing.T) {
	// Mock global settings
	global.SaveLoadSetting = &setting.SaveLoadSettingS{
		SavePath:    "testing",
		SaveFileExt: ".json",
	}

	// Create test data
	testFilename := "test_units_json"
	testFilePath := global.SaveLoadSetting.SavePath + "/" + testFilename + "sv" + global.SaveLoadSetting.SaveFileExt
	testUnits := []entities.Player{
		{ID: 1, Name: "Player1", Components: make(map[string]interface{})},
		{ID: 2, Name: "Player2", Components: make(map[string]interface{})},
	}

	// Create test directory
	err := os.MkdirAll(global.SaveLoadSetting.SavePath, os.ModePerm)
	assert.NoError(t, err)
	defer os.RemoveAll(global.SaveLoadSetting.SavePath)

	// Write test data to file
	file, err := os.Create(testFilePath)
	assert.NoError(t, err)
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(testUnits)
	assert.NoError(t, err)

	// Test successful load
	loadedUnits, err := ecs_systems.LoadUnitsState(testFilename)
	assert.NoError(t, err)
	assert.Equal(t, testUnits, loadedUnits)

	// Test file not found
	_, err = ecs_systems.LoadUnitsState("nonexistent_file_json")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, os.ErrNotExist))

	// Test invalid file format
	invalidFilePath := global.SaveLoadSetting.SavePath + "/invalid_file_json" + global.SaveLoadSetting.SaveFileExt
	invalidFileContent := []byte("invalid content")
	err = os.WriteFile(invalidFilePath, invalidFileContent, os.ModePerm)
	assert.NoError(t, err)

	_, err = ecs_systems.LoadUnitsState("invalid_file_json")
	assert.Error(t, err)
}
