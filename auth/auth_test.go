package auth

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	token := NewToken(1)

	if token == "" {
		t.Errorf("Token is empty")
	}
}
