package config_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mateusz-uminski/go-nethttp-healthz/util/config"
	"github.com/stretchr/testify/assert"
)

const envPrefix = "TEST_APP"

func TestGetHost(t *testing.T) {
	t.Run("should load default value when HOST env is not set", func(t *testing.T) {
		defer unsetenvs()
		c := newConfig()

		// when
		actualHost := c.GetHost()

		// then
		assert.Equal(t, config.DEFAULT_HOST, actualHost)
	})

	t.Run("should load custom value when HOST env is set", func(t *testing.T) {
		defer unsetenvs()
		value := "0.0.0.0"
		setenv("HOST", value)
		c := newConfig()

		// when
		actualHost := c.GetHost()

		// then
		assert.Equal(t, value, actualHost)
	})
}

func TestGetPort(t *testing.T) {
	t.Run("should load default value when PORT env is not set", func(t *testing.T) {
		defer unsetenvs()
		c := newConfig()

		// when
		actualPort := c.GetPort()

		// then
		assert.Equal(t, config.DEFAULT_PORT, actualPort)
	})

	t.Run("should load custom value when PORT env is set", func(t *testing.T) {
		defer unsetenvs()
		value := 9999
		setenv("PORT", strconv.Itoa(value))
		c := newConfig()

		// when
		actualPort := c.GetPort()

		// then
		assert.Equal(t, value, actualPort)
	})
}

func TestGetLogLevel(t *testing.T) {
	t.Run("should load default value when LOG_LEVEL env is not set", func(t *testing.T) {
		defer unsetenvs()
		c := newConfig()

		// when
		actualLogLevel := c.GetLogLevel()

		// then
		assert.Equal(t, config.DEFAULT_LOG_LEVEL, actualLogLevel)
	})

	t.Run("should load custom value when LOG_LEVEL env is set", func(t *testing.T) {
		defer unsetenvs()
		value := "error"
		setenv("LOG_LEVEL", value)
		c := newConfig()

		// when
		actualLogLevel := c.GetLogLevel()

		// then
		assert.Equal(t, value, actualLogLevel)
	})
}

func unsetenvs() {
	os.Unsetenv("HOST")
	os.Unsetenv("PORT")
	os.Unsetenv("LOG_LEVEL")
}

func setenv(name string, value string) {
	nameWithPrefix := envPrefix + "_" + name
	os.Setenv(nameWithPrefix, value)
}

func newConfig() config.Config {
	return config.New(envPrefix)
}
