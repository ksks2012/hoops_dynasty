package entities

import (
	"github.com/hoops_dynasty/ecs/components/attributes"
)

type Player struct {
	ID         int                    `json:"id"`
	Name       string                 `json:"name"`
	Components map[string]interface{} `json:"components"`
	Potential  int                    `json:"potential"`
	Morale     int                    `json:"morale"`
	// Contract    Contract            `json:"contract"`
	// Personality Personality        `json:"personality"`
}

func NewPlayer(id int, name string, age int) *Player {
	return &Player{
		ID: id,
		Components: map[string]interface{}{
			"basic":    &attributes.BasicAttributes{Height: 180, Weight: 70, Age: age},
			"offense":  &attributes.OffensiveAttributes{Shooting: 50, Passing: 50, Finishing: 50, Dribbling: 50},
			"defense":  &attributes.DefensiveAttributes{Block: 50, Steal: 50, Rebound: 50, DefensiveIQ: 50},
			"physical": &attributes.PhysicalAttributes{Speed: 50, Stamina: 50, Strength: 50, Agility: 50},
			"mental":   &attributes.MentalAttributes{CourtVision: 50, Leadership: 50, WorkRate: 50, Focus: 50},
		},
	}
}

func (p *Player) GetComponent(name string) interface{} {
	return p.Components[name]
}
