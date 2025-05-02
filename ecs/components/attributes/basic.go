package attributes

import "fmt"

// BasicAttributes represents a player's basic physical characteristics
type BasicAttributes struct {
	Height float64 `json:"height"` // Height in centimeters
	Weight float64 `json:"weight"` // Weight in kilograms
	Age    int     `json:"age"`    // Age in years
}

// NewBasicAttributes creates and initializes a player's basic attributes
func NewBasicAttributes(height, weight float64, age int) BasicAttributes {
	return BasicAttributes{
		Height: height,
		Weight: weight,
		Age:    age,
	}
}

// Display shows the player's basic attributes
func (ba BasicAttributes) Display() {
	fmt.Printf("Height: %.2f cm\n", ba.Height)
	fmt.Printf("Weight: %.2f kg\n", ba.Weight)
	fmt.Printf("Age: %d years\n", ba.Age)
}
