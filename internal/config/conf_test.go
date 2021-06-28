package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var expectedServerConf = ServerConfig{
	Port:            8000,
	IdleTimeout:     2 * time.Minute,
	ReadTimeout:     5 * time.Second,
	WriteTimeout:    5 * time.Second,
	ShutdownTimeout: 15 * time.Second,
}

var expectedDBConf = DBConfig{
	Host:     "localhost",
	Port:     5432,
	Name:     "shortener",
	User:     "shortener",
	Password: "123456789",
	SSLMode:  "verify-full",
}

func TestParseServerConfig(t *testing.T) {
	conf, err := parseServerConfig()

	assert.NoError(t, err)
	assert.Equal(t, expectedServerConf, *conf)
}

func TestParseDBConfig(t *testing.T) {
	conf, err := parseDBConfig()

	assert.NoError(t, err)
	assert.Equal(t, expectedDBConf, *conf)
}

func TestGetConfig(t *testing.T) {
	conf, err := GetConfig()

	assert.NoError(t, err)
	assert.Equal(t, expectedServerConf, *conf.Server)
	assert.Equal(t, expectedDBConf, *conf.DB)
}
