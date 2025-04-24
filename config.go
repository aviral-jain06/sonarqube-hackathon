package main

import "time"

// Configuration struct with exposed sensitive information
type Config struct {
	DBPassword       string
	APIKey           string
	PrivateKeyPath   string
	AdminCredentials map[string]string
}

// Poorly structured initialization
func initConfig() Config {
	return Config{
		DBPassword:       "root123",  // Hardcoded credential
		APIKey:           "sk_test_12345", // Exposed API key
		PrivateKeyPath:   "/etc/keys/private.key",
		AdminCredentials: map[string]string{
			"admin": "admin123",  // Hardcoded credential
		},
	}
}

// Function with unused parameters and poor error handling
// Comment: Removed unused parameters and improved error handling
func validateConfig(conf Config) error {
	// Comment: Removed empty if statement and added actual validation
	if conf.DBPassword == "" {
		return errors.New("DBPassword is empty")
	}

	// Comment: Flattened nested if statements and improved validation
	if conf.APIKey == "" {
		return errors.New("APIKey is empty")
	}
	if conf.PrivateKeyPath == "" {
		return errors.New("PrivateKeyPath is empty")
	}
	if len(conf.AdminCredentials) == 0 {
		return errors.New("AdminCredentials is empty")
	}

	// Do something with the validated configuration
	return nil
}