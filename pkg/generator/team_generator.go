package generator

import (
	components "github.com/hoops_dynasty/ecs/components"
)

type Team struct {
	ID         int
	Components map[string]interface{}
}

func NewRandomTeam(id int, playerCount int) *Team {
	team := &Team{
		ID: id,
		Components: map[string]interface{}{
			"roster": &components.Roster{
				Starters: make([]int, 0, 5),
				Bench:    make([]int, 0, playerCount-5),
			},
			// TODO:
			"tactics": &components.Tactics{
				OffenseStyle: "balanced",
				DefenseStyle: "man_to_man",
				Pace:         50,
			},
			"stats":  &components.TeamStats{},
			"morale": &components.Morale{TeamMorale: 50},
		},
	}

	// Generate players and add to the roster
	roster := team.Components["roster"].(*components.Roster)
	for i := 0; i < playerCount; i++ {
		player := GeneratePlayer(id*100 + i)
		if i < 5 {
			roster.Starters = append(roster.Starters, player.ID)
		} else {
			roster.Bench = append(roster.Bench, player.ID)
		}
	}

	return team
}
