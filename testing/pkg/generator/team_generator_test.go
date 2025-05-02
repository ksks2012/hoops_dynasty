package testing

import (
	"testing"

	components "github.com/hoops_dynasty/ecs/components"
	"github.com/hoops_dynasty/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func TestNewRandomTeam(t *testing.T) {
	playerCount := 10
	teamID := 1

	team := generator.NewRandomTeam(teamID, playerCount)

	// Test team ID
	assert.Equal(t, teamID, team.ID, "Team ID should match the input ID")

	// Test roster component
	roster, ok := team.Components["roster"].(*components.Roster)
	assert.True(t, ok, "Roster component should exist and be of type *components.Roster")
	assert.Len(t, roster.Starters, 5, "Starters should have 5 players")
	assert.Len(t, roster.Bench, playerCount-5, "Bench should have the remaining players")

	// Test tactics component
	tactics, ok := team.Components["tactics"].(*components.Tactics)
	assert.True(t, ok, "Tactics component should exist and be of type *components.Tactics")
	assert.Equal(t, "balanced", tactics.OffenseStyle, "OffenseStyle should be 'balanced'")
	assert.Equal(t, "man_to_man", tactics.DefenseStyle, "DefenseStyle should be 'man_to_man'")
	assert.Equal(t, 50, tactics.Pace, "Pace should be 50")

	// Test stats component
	stats, ok := team.Components["stats"].(*components.TeamStats)
	assert.True(t, ok, "Stats component should exist and be of type *components.TeamStats")
	assert.Equal(t, 0, stats.Wins, "Wins should be 0")
	assert.Equal(t, 0, stats.Losses, "Losses should be 0")
	assert.Equal(t, 0.0, stats.PointsPerGame, "PointsPerGame should be 0.0")

	// Test morale component
	morale, ok := team.Components["morale"].(*components.Morale)
	assert.True(t, ok, "Morale component should exist and be of type *components.Morale")
	assert.Equal(t, 50, morale.TeamMorale, "TeamMorale should be 50")
}

func TestNewRandomTeam_PlayerIDs(t *testing.T) {
	playerCount := 8
	teamID := 2

	team := generator.NewRandomTeam(teamID, playerCount)

	roster := team.Components["roster"].(*components.Roster)

	// Test player IDs in starters
	for i, playerID := range roster.Starters {
		expectedID := teamID*100 + i
		assert.Equal(t, expectedID, playerID, "Starter player ID should match the expected ID")
	}

	// Test player IDs in bench
	for i, playerID := range roster.Bench {
		expectedID := teamID*100 + 5 + i
		assert.Equal(t, expectedID, playerID, "Bench player ID should match the expected ID")
	}
}
