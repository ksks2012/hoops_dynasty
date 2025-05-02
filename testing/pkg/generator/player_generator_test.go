package testing

import (
	"testing"

	components "github.com/hoops_dynasty/ecs/components/attributes"
	"github.com/hoops_dynasty/global"
	"github.com/hoops_dynasty/pkg/generator"
	"github.com/hoops_dynasty/pkg/setting"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePlayer(t *testing.T) {
	// Mock global settings
	global.BasicSetting = &setting.BasicSettingS{
		HeightRange: setting.AttributeRange{Min: 180, Max: 210},
		WeightRange: setting.AttributeRange{Min: 70, Max: 120},
		AgeRange:    setting.AttributeRange{Min: 18, Max: 40},
	}

	global.DefenseSetting = &setting.DefenseSettingS{
		PerimeterDefenseRange: setting.AttributeRange{Min: 50, Max: 100},
		InteriorDefenseRange:  setting.AttributeRange{Min: 50, Max: 100},
		BlockRange:            setting.AttributeRange{Min: 50, Max: 100},
		StealRange:            setting.AttributeRange{Min: 50, Max: 100},
		ReboundRange:          setting.AttributeRange{Min: 50, Max: 100},
		DefensiveIQRange:      setting.AttributeRange{Min: 50, Max: 100},
	}

	global.MentalSetting = &setting.MentalSettingS{
		CourtVisionRange:          setting.AttributeRange{Min: 50, Max: 100},
		DefensivePositioningRange: setting.AttributeRange{Min: 50, Max: 100},
		WorkRateRange:             setting.AttributeRange{Min: 50, Max: 100},
		LeadershipRange:           setting.AttributeRange{Min: 50, Max: 100},
		FocusRange:                setting.AttributeRange{Min: 50, Max: 100},
	}

	global.OffenseSetting = &setting.OffenseSettingS{
		ShootingRange:         setting.AttributeRange{Min: 50, Max: 100},
		ShootingDistanceRange: setting.AttributeRange{Min: 50, Max: 100},
		FinishingRange:        setting.AttributeRange{Min: 50, Max: 100},
		PassingRange:          setting.AttributeRange{Min: 50, Max: 100},
		DribblingRange:        setting.AttributeRange{Min: 50, Max: 100},
	}

	global.PhysicalSetting = &setting.PhysicalSettingS{
		SpeedRange:        setting.AttributeRange{Min: 50, Max: 100},
		StaminaRange:      setting.AttributeRange{Min: 50, Max: 100},
		StrengthRange:     setting.AttributeRange{Min: 50, Max: 100},
		VerticalLeapRange: setting.AttributeRange{Min: 50, Max: 100},
		AgilityRange:      setting.AttributeRange{Min: 50, Max: 100},
	}

	// Generate a player
	player := generator.GeneratePlayer(1)

	// Assert player is not nil
	assert.NotNil(t, player)

	// Assert player ID
	assert.Equal(t, 1, player.ID)

	// Assert player components
	assert.NotNil(t, player.Components["basic"])
	assert.NotNil(t, player.Components["defense"])
	assert.NotNil(t, player.Components["mental"])
	assert.NotNil(t, player.Components["offense"])
	assert.NotNil(t, player.Components["physical"])

	// Assert basic attributes
	basic := player.Components["basic"].(*components.BasicAttributes)
	assert.GreaterOrEqual(t, basic.Height, 180.0)
	assert.LessOrEqual(t, basic.Height, 210.0)
	assert.GreaterOrEqual(t, basic.Weight, 70.0)
	assert.LessOrEqual(t, basic.Weight, 120.0)
	assert.GreaterOrEqual(t, basic.Age, 18)
	assert.LessOrEqual(t, basic.Age, 40)

	// Assert defensive attributes
	defense := player.Components["defense"].(*components.DefensiveAttributes)
	assert.GreaterOrEqual(t, defense.PerimeterDefense, 50.0)
	assert.LessOrEqual(t, defense.PerimeterDefense, 100.0)

	// Assert offensive attributes
	offense := player.Components["offense"].(*components.OffensiveAttributes)
	assert.GreaterOrEqual(t, offense.Shooting, 50.0)
	assert.LessOrEqual(t, offense.Shooting, 100.0)

	// Assert physical attributes
	physical := player.Components["physical"].(*components.PhysicalAttributes)
	assert.GreaterOrEqual(t, physical.Speed, 50.0)
	assert.LessOrEqual(t, physical.Speed, 100.0)
}
