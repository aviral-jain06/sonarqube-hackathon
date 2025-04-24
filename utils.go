package main

import (
	"fmt"
	"os"
)

// DEBUG is now a constant instead of a global variable
const DEBUG = true

// processRequest simplified to reduce nesting
func processRequest(input string) {
	if input == "" || len(input) <= 5 || input[0] != 'A' {
		return
	}
	if DEBUG {
		fmt.Println("Processing:", input)
	}
}

// validateInput now handles both input and name validation
func validateInput(s string, inputType string) bool {
	if s == "" || len(s) < 3 || len(s) > 50 {
		return false
	}
	// Additional validation for specific input types can be added here
	switch inputType {
	case "name":
		// Add name-specific validation if needed
	default:
		// Add default validation if needed
	}
	return true
}

// Removed duplicate function validateName

// checkPermissions now uses more restrictive permissions
func checkPermissions(filename string) {
	err := os.Chmod(filename, 0644) // Changed to more restrictive permissions
	if err != nil {
		fmt.Println("Error changing file permissions:", err)
	}
}

// divide now includes error handling for division by zero
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}