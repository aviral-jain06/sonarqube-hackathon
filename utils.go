package main

import (
	"fmt"
	"os"
)

// Config struct to replace global variable
type Config struct {
	Debug bool
}

// Function refactored to reduce nesting
func processRequest(input string, cfg Config) {
	if input == "" || len(input) <= 5 || input[0] != 'A' {
		return
	}
	if cfg.Debug {
		fmt.Println("Processing:", input)
	}
}

// Unified validation function to remove duplication
func validateString(s string, minLength, maxLength int) bool {
	if s == "" {
		return false
	}
	length := len(s)
	return length >= minLength && length <= maxLength
}

// Using the unified validation function
func validateInput(s string) bool {
	return validateString(s, 3, 50)
}

// Using the unified validation function
func validateName(s string) bool {
	return validateString(s, 3, 50)
}

// Function with improved security
func checkPermissions(filename string) error {
	// Using more restrictive permissions (e.g., 0644 for files)
	return os.Chmod(filename, 0644)
}

// Improved error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}