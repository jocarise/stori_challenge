package utils

import (
	"newsletter-service/pkg/regex"
	"regexp"
	"testing"
)

func TestValidateEmailRequirements(t *testing.T) {
	_, err := regexp.Compile(regex.EmailRegex)
	if err != nil {
		t.Fatalf("Failed to compile Email regex: %v", err)
	}

	validEmails := []string{
		"test@example.com",
		"user.name+tag+sorting@example.com",
		"user_name@example.co.uk",
		"user-name@example.com",
		"user@subdomain.example.com",
	}

	invalidEmails := []string{
		"plainaddress",         // No @
		"@missingusername.com", // Missing username
		"username@.com",        // Missing domain
		"username@com",         // No top-level domain
		"username@.com.",       // Trailing dot
		"username@..com",       // Double dots
		"username@com.",        // Trailing dot in domain
	}

	for _, email := range validEmails {
		if !ValidateEmailRequirements(email) {
			t.Errorf("Expected valid email: %s, but it was not validated", email)
		}
	}

	for _, email := range invalidEmails {
		if ValidateEmailRequirements(email) {
			t.Errorf("Expected invalid email: %s, but it was validated", email)
		}
	}
}
