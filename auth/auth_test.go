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

func TestParseTokenBadToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzUxMiIsnR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOiI5MGVlZDdiYS1lMzZlLTRlOTQtOGE1NC1iMzM4ZDIyYjVmODUiLCJleHAiOjE2NzkzNTM0OTgsImlhdCI6MTU5Mjk1MzQ5OH0.W3YSuwoli5U8VjZ6B0irzheWMMh96uzTojWOTH4Y2g_DJpb1cOX9HHVvdK9HBkBple2fUpM6SgJBU1mFIm82YA"

	token, err := parseToken(tokenString)

	if err == nil || token != nil {
		t.Errorf("Parse token should fail")
	}
}
