package generator

import (
	"math/rand"
	"time"

	"github.com/hoops_dynasty/ecs/components/attributes"
	"github.com/hoops_dynasty/ecs/entities"
	"github.com/hoops_dynasty/global"
	"github.com/hoops_dynasty/pkg/setting"
)

// randomInt generates a random integer within the given range (inclusive)
func randomInt(r setting.AttributeRange) int {
	if r.Min >= r.Max {
		return r.Min
	}
	return rand.Intn(r.Max-r.Min+1) + r.Min
}

func randomFloat(r setting.AttributeRange) float64 {
	if r.Min >= r.Max {
		return float64(r.Min)
	}
	return float64(rand.Intn(r.Max-r.Min+1)+r.Min) / 10.0
}

// randomName generates a placeholder name (replace with your name generation logic)
func randomName() string {
	names := []string{"John Doe", "Mike Smith", "Chris Johnson", "Alex Brown"}
	return names[rand.Intn(len(names))]
}

// GeneratePlayer creates a new player with attributes based on the provided settings
func GeneratePlayer(id int) *entities.Player {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create player components based on settings
	basic := &attributes.BasicAttributes{
		Height: randomFloat(global.BasicSetting.HeightRange),
		Weight: randomFloat(global.BasicSetting.WeightRange),
		Age:    randomInt(global.BasicSetting.AgeRange),
	}

	defense := &attributes.DefensiveAttributes{
		PerimeterDefense: randomFloat(global.DefenseSetting.PerimeterDefenseRange),
		InteriorDefense:  randomFloat(global.DefenseSetting.InteriorDefenseRange),
		Block:            randomFloat(global.DefenseSetting.BlockRange),
		Steal:            randomFloat(global.DefenseSetting.StealRange),
		Rebound:          randomFloat(global.DefenseSetting.ReboundRange),
		DefensiveIQ:      randomFloat(global.DefenseSetting.DefensiveIQRange),
	}

	mental := &attributes.MentalAttributes{
		CourtVision:          randomFloat(global.MentalSetting.CourtVisionRange),
		DefensivePositioning: randomFloat(global.MentalSetting.DefensivePositioningRange),
		WorkRate:             randomFloat(global.MentalSetting.WorkRateRange),
		Leadership:           randomFloat(global.MentalSetting.LeadershipRange),
		Focus:                randomFloat(global.MentalSetting.FocusRange),
	}

	offense := &attributes.OffensiveAttributes{
		Shooting:         randomFloat(global.OffenseSetting.ShootingRange),
		ShootingDistance: randomFloat(global.OffenseSetting.ShootingDistanceRange),
		Finishing:        randomFloat(global.OffenseSetting.FinishingRange),
		Passing:          randomFloat(global.OffenseSetting.PassingRange),
		Dribbling:        randomFloat(global.OffenseSetting.DribblingRange),
	}

	physical := &attributes.PhysicalAttributes{
		Speed:        randomFloat(global.PhysicalSetting.SpeedRange),
		Stamina:      randomFloat(global.PhysicalSetting.StaminaRange),
		Strength:     randomFloat(global.PhysicalSetting.StrengthRange),
		VerticalLeap: randomFloat(global.PhysicalSetting.VerticalLeapRange),
		Agility:      randomFloat(global.PhysicalSetting.AgilityRange),
	}

	// Create player entity
	return &entities.Player{
		ID: id,
		Components: map[string]interface{}{
			"basic":    basic,
			"defense":  defense,
			"mental":   mental,
			"offense":  offense,
			"physical": physical,
		},
	}
}
