package attributes

import "fmt"

// DefensiveAttributes represents a player's defensive abilities
type DefensiveAttributes struct {
	PerimeterDefense float64 // Ability to defend on the perimeter and limit the opponent's shooting accuracy
	InteriorDefense  float64 // Ability to defend inside the paint, affecting block and contest success
}

// NewDefensiveAttributes creates and initializes a character's defensive attributes
func NewDefensiveAttributes(perimeterDefense, interiorDefense float64) DefensiveAttributes {
	return DefensiveAttributes{
		PerimeterDefense: perimeterDefense,
		InteriorDefense:  interiorDefense,
	}
}

// Display shows the character's defensive attributes
func (da DefensiveAttributes) Display() {
	fmt.Printf("Perimeter Defense: %.2f\n", da.PerimeterDefense)
	fmt.Printf("Interior Defense: %.2f\n", da.InteriorDefense)
}
