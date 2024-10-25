package utils

import (
	"regexp"
	"testing"
	"user-service/pkg/regex"
)

func TestValidatePasswordRequirements(t *testing.T) {
	validPasswords := []string{
		"Valid1@",
		"Str0ng@Password",
		"Password123!",
		"Another1!",
	}

	invalidPasswords := []string{
		"short",           // Too short
		"nouppercase1!",   // No uppercase
		"NOLOWERCASE1!",   // No lowercase
		"WithoutNumber!",  // No number
		"WithoutSpecial1", // No special character
	}

	for _, password := range validPasswords {
		if !ValidatePasswordRequirements(password) {
			t.Errorf("Expected valid password: %s, but it was not validated", password)
		}
	}

	for _, password := range invalidPasswords {
		if ValidatePasswordRequirements(password) {
			t.Errorf("Expected invalid password: %s, but it was validated", password)
		}
	}
}

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

func TestHashAndComparePassword(t *testing.T) {
	password := "Valid1@Password"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	err = ComparePasswords(hashedPassword, password)
	if err != nil {
		t.Errorf("Expected passwords to match, but they did not: %v", err)
	}

	err = ComparePasswords(hashedPassword, "WrongPassword!")
	if err == nil {
		t.Error("Expected passwords to not match, but they matched")
	}
}
