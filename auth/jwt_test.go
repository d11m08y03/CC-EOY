package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateJWT(t *testing.T) {
	userID := uint(1)
	email := "test@example.com"

	token, err := GenerateJWT(userID, email)
	if err != nil {
		t.Fatalf("GenerateJWT returned an error: %v", err)
	}

	// Validate the generated JWT
	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("ValidateJWT returned an error: %v", err)
	}

	// Verify the claims
	if claims.OrganisorID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, claims.OrganisorID)
	}
	if claims.Email != email {
		t.Errorf("Expected Email %v, got %v", email, claims.Email)
	}
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		t.Errorf("Token expiry time is invalid")
	}
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"

	_, err := ValidateJWT(invalidToken)
	if err == nil {
		t.Error("Expected an error for invalid token, got nil")
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	// Create a token with a past expiration time
	expiredClaims := &Claims{
		OrganisorID: 1,
		Email:  "test@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	// Validate the expired token
	_, err = ValidateJWT(tokenStr)
	if err == nil {
		t.Error("Expected an error for expired token, got nil")
	}
}
