package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(secret []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func TestJWTMiddleware(t *testing.T) {
	SECRET := []byte("JWT_SECRET")

	tests := []struct {
		name           string
		authorization  string
		expectedStatus int
	}{
		{
			name:           "Valid Token",
			authorization:  "Bearer " + GenerateToken(SECRET),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Missing Token",
			authorization:  "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Token",
			authorization:  "Bearer invalid.token.here",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Format",
			authorization:  "Basic somecredentials",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authorization != "" {
				req.Header.Set("Authorization", tt.authorization)
			}

			rr := httptest.NewRecorder()

			middleware := JWTMiddleware(SECRET)
			middleware(handler).ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, status)
			}
		})
	}
}
