package components

import "fmt"

// OffensiveTechnicalAttributes represents a player's offensive abilities
type OffensiveAttributes struct {
	Shooting         float64 `json:"shooting"`          // Shooting accuracy for close, mid-range, and three-point shots
	ShootingDistance float64 `json:"shooting_distance"` // Shooting range for close, mid-range, and three-point shots
	Finishing        float64 `json:"finishing"`         // Scoring efficiency in the paint, especially under defensive pressure
	Passing          float64 `json:"passing"`           // Pass accuracy and ability to create scoring opportunities
	Dribbling        float64 `json:"dribbling"`         // Ability to find space and break through defenders
}

// NewOffensiveAttributes creates and initializes a character's offensive attributes
func NewOffensiveAttributes(shooting, shootingDistance, finishing, passing, dribbling float64) OffensiveAttributes {
	return OffensiveAttributes{
		Shooting:         shooting,
		ShootingDistance: shootingDistance,
		Finishing:        finishing,
		Passing:          passing,
		Dribbling:        dribbling,
	}
}

// Display shows the character's offensive attributes
func (oa OffensiveAttributes) Display() {
	fmt.Printf("Shooting: %.2f\n", oa.Shooting)
	fmt.Printf("Shooting Range: %.2f\n", oa.ShootingDistance)
	fmt.Printf("Finishing: %.2f\n", oa.Finishing)
	fmt.Printf("Passing: %.2f\n", oa.Passing)
	fmt.Printf("Dribbling: %.2f\n", oa.Dribbling)
}
