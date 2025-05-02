package components

type Roster struct {
	Starters []int `json:"starters"`
	Bench    []int `json:"bench"`
}

type Tactics struct {
	OffenseStyle string `json:"offense_style"`
	DefenseStyle string `json:"defense_style"`
	Pace         int    `json:"pace"`
}

type TeamStats struct {
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	PointsPerGame float64 `json:"points_per_game"`
	// NOTE: Other statistical data
}

type Morale struct {
	TeamMorale int `json:"team_morale"`
}
