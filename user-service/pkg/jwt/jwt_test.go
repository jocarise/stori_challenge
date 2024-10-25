package jwt

import (
	"testing"
	"time"
	"user-service/src/models"

	"github.com/golang-jwt/jwt"
)

var mockJWTKey = []byte("60346e6a684673efddc312ab6a2df00ab0de3efa1b5061e452e7624addf57bd6")

func TestGenerateJWT(t *testing.T) {
	user := &models.User{
		ID:    "123",
		Email: "test@example.com",
		Role:  "user",
	}

	token, err := GenerateJWT(user, mockJWTKey)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if token is not empty
	if token == "" {
		t.Fatal("Expected a valid token, got empty string")
	}

	// Validate token
	claims, err := ValidateJWT(token, mockJWTKey)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate claims
	if claims["sub"] != user.ID {
		t.Errorf("Expected user ID %s, got %s", user.ID, claims["sub"])
	}
	if claims["email"] != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, claims["email"])
	}
	if claims["role"] != user.Role {
		t.Errorf("Expected role %s, got %s", user.Role, claims["role"])
	}
}

func TestValidateJWTInvalidToken(t *testing.T) {
	_, err := ValidateJWT("invalid.token.string", mockJWTKey)
	if err == nil {
		t.Fatal("Expected an error for invalid token, got nil")
	}
}

func TestValidateJWTExpiredToken(t *testing.T) {
	// Create an expired token for testing
	claims := jwt.MapClaims{
		"sub":   "123",
		"email": "test@example.com",
		"role":  "user",
		"exp":   time.Now().Add(-time.Hour).Unix(), // Set expiration to past
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	expiredToken, err := token.SignedString(mockJWTKey)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	_, err = ValidateJWT(expiredToken, mockJWTKey)
	if err == nil {
		t.Fatal("Expected an error for expired token, got nil")
	}
}
