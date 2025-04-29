package attributes

import "fmt"

// MentalAttributes represents a player's mental attributes
type MentalAttributes struct {
	CourtVision          float64 // Ability to recognize teammates' positions and execute effective offense
	DefensivePositioning float64 // Ability to maintain good positioning on defense and anticipate opponents' moves
	WorkRate             float64 // Player's effort and activity on the court, especially in defense
	Leadership           float64 // Ability to motivate teammates and improve teamwork
	Focus                float64 // Concentration level, helps avoid mistakes and stay attentive
}

// NewMentalAttributes creates and initializes a player's mental attributes
func NewMentalAttributes(courtVision, defensivePositioning, workRate, leadership, focus float64) MentalAttributes {
	return MentalAttributes{
		CourtVision:          courtVision,
		DefensivePositioning: defensivePositioning,
		WorkRate:             workRate,
		Leadership:           leadership,
		Focus:                focus,
	}
}

// Display shows the player's mental attributes
func (ma MentalAttributes) Display() {
	fmt.Printf("Court Vision: %.2f\n", ma.CourtVision)
	fmt.Printf("Defensive Positioning: %.2f\n", ma.DefensivePositioning)
	fmt.Printf("Work Rate: %.2f\n", ma.WorkRate)
	fmt.Printf("Leadership: %.2f\n", ma.Leadership)
	fmt.Printf("Focus: %.2f\n", ma.Focus)
}
