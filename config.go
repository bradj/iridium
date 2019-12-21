package main

import (
	"github.com/BurntSushi/toml"
)

// TomlConfig struct
type TomlConfig struct {
	DB DB `toml:"database"`
}

// DB struct
type DB struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

// NewConfig creates a new config at configLocation
func NewConfig(filename string) (c TomlConfig, err error) {
	_, err = toml.DecodeFile(filename, &c)

	return c, err
}
