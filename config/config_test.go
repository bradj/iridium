package config

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	t.Parallel()

	c, err := NewConfig("fixtures/config.toml")

	if err != nil {
		t.Error("err should be nil", err)
	}

	if c.DB.Username != "root" {
		t.Errorf("Config DB Username is '%v' and should be 'root'", c.DB.Username)
	}

	if c.DB.Password != "somepassword" {
		t.Errorf("Config DB Password is '%v' and should be 'somepassword'", c.DB.Password)
	}

	if c.DB.Host != "192.168.1.1" {
		t.Errorf("Config DB Host is '%v' and should be '192.168.1.1'", c.DB.Host)
	}

	if c.DB.Port != 8000 {
		t.Errorf("Config DB Port is '%v' and should be 8000", c.DB.Port)
	}
}
