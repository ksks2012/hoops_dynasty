package attributes

import "fmt"

// PhysicalAttributes represents a player's physical abilities
type PhysicalAttributes struct {
	Speed        float64 // Movement speed in transitions and drives
	Stamina      float64 // Lasting energy, affects performance in later stages of a game
	Strength     float64 // Physical stability in confrontations, especially near the basket
	VerticalLeap float64 // Jumping height for rebounds, blocks, and layups
	Agility      float64 // Quick movement and turning ability for offensive drives and defensive reactions
}

// NewPhysicalAttributes creates and initializes a character's physical attributes
func NewPhysicalAttributes(speed, stamina, strength, verticalLeap, agility float64) PhysicalAttributes {
	return PhysicalAttributes{
		Speed:        speed,
		Stamina:      stamina,
		Strength:     strength,
		VerticalLeap: verticalLeap,
		Agility:      agility,
	}
}

// Display shows the character's physical attributes
func (pa PhysicalAttributes) Display() {
	fmt.Printf("Speed: %.2f\n", pa.Speed)
	fmt.Printf("Stamina: %.2f\n", pa.Stamina)
	fmt.Printf("Strength: %.2f\n", pa.Strength)
	fmt.Printf("Vertical Leap: %.2f\n", pa.VerticalLeap)
	fmt.Printf("Agility: %.2f\n", pa.Agility)
}
