package auth

import (
	"testing"
)

func TestNewTokenAndParseToken(t *testing.T) {
	t.Parallel()

	uid := "1234-74748-483243-423432432"
	token := NewToken(uid)

	if token == "" {
		t.Errorf("Token is empty")
	}

	parsedToken, err := parseToken(token)

	if err != nil {
		t.Error("parseToken failed", err)
	}

	claims := parsedToken.Claims.(*IridiumClaims)

	if claims.UserId != uid {
		t.Errorf("UserId claim is %s instead of %s", claims.UserId, uid)
	}
}

func TestPasswordHashAndCompare(t *testing.T) {
	t.Parallel()

	password := "mypassword"

	hash, err := GeneratePasswordHash(password)

	if err != nil {
		t.Errorf("Password hash failed %v", err)
	}

	if hash == nil {
		t.Error("Password hash is empty")
	}

	// Test PasswordHashCompare
	err = PasswordHashCompare(hash, password)

	if err != nil {
		t.Errorf("Password hashses did not match: %v", err)
	}
}
