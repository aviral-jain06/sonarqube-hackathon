package main

import (
	"fmt"
	"os"
)

// Replaced global variable with a constant
const debug = true

// Simplified nested if statements
func processRequest(input string) {
	if input != "" && len(input) > 5 && input[0] == 'A' && debug {
		fmt.Println("Processing:", input)
	}
}

// Generalized input validation function
func validateInput(s string, minLength, maxLength int) bool {
	if s == "" {
		return false
	}
	length := len(s)
	return length >= minLength && length <= maxLength
}

// Updated to use the generalized validateInput function
func validateName(s string) bool {
	return validateInput(s, 3, 50)
}

// Improved security by using more restrictive permissions
func checkPermissions(filename string) {
	err := os.Chmod(filename, 0600) // Changed to read/write for owner only
	if err != nil {
		fmt.Println("Error changing file permissions:", err)
	}
}

// Added error handling for division by zero
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}