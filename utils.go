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

// validateInput function remains the same
func validateInput(s string) bool {
	if s == "" {
		return false
	}
	if len(s) < 3 {
		return false
	}
	if len(s) > 50 {
		return false
	}
	return true
}

// validateName now calls validateInput to avoid duplication
func validateName(s string) bool {
	return validateInput(s)
}

// checkPermissions now uses more restrictive permissions
func checkPermissions(filename string) {
	err := os.Chmod(filename, 0600) // Changed to read/write for owner only
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