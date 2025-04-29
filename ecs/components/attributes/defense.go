package attributes

import "fmt"

// DefensiveAttributes represents a player's defensive abilities
type DefensiveAttributes struct {
	PerimeterDefense float64 // Ability to defend on the perimeter and limit the opponent's shooting accuracy
	InteriorDefense  float64 // Ability to defend inside the paint, affecting block and contest success
	Block            float64
	Steal            float64
	Rebound          float64
	DefenseIQ        float64
}

// NewDefensiveAttributes creates and initializes a character's defensive attributes
func NewDefensiveAttributes(perimeterDefense, interiorDefense, block, steal, rebound, defenseIQ float64) DefensiveAttributes {
	return DefensiveAttributes{
		PerimeterDefense: perimeterDefense,
		InteriorDefense:  interiorDefense,
		Block:            block,
		Steal:            steal,
		Rebound:          rebound,
		DefenseIQ:        defenseIQ,
	}
}

// Display shows the character's defensive attributes
func (da DefensiveAttributes) Display() {
	fmt.Printf("Perimeter Defense: %.2f\n", da.PerimeterDefense)
	fmt.Printf("Interior Defense: %.2f\n", da.InteriorDefense)
	fmt.Printf("Block: %.2f\n", da.Block)
	fmt.Printf("Steal: %.2f\n", da.Steal)
	fmt.Printf("Rebound: %.2f\n", da.Rebound)
	fmt.Printf("Defense IQ: %.2f\n", da.DefenseIQ)
}
