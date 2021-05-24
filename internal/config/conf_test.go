package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	conf, err := GetConfig()
	expectedConf := Config{
		Port:            8000,
		IdleTimeout:     2 * time.Minute,
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
		ShutdownTimeout: 15 * time.Second,
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedConf, *conf)
}
