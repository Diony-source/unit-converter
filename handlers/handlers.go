// Handles user input and output interactions
package handlers

import (
	"fmt"
	"unit-converter/entities"
	"unit-converter/services"
)

// Start initiates the user interaction process
func Start() {
	var value float64
	var conversionType string

	// Prompt user for the value to convert
	fmt.Println("Enter the value to convert:")
	fmt.Scanln(&value)

	// Prompt user for the conversion type
	fmt.Println("Enter the conversion type (e.g., km_to_miles, kg_to_pounds, c_to_f):")
	fmt.Scanln(&conversionType)

	// Create a Conversion struct
	conversion := entities.Conversion{
		Value: value,
		Type:  conversionType,
	}

	// Perform the conversion
	result, err := services.Convert(conversion)
	if err != nil {
		// Display error if any
		fmt.Println("Error:", err)
	} else {
		// Display the conversion result
		fmt.Printf("The converted value is: %.2f\n", result)
	}
}
