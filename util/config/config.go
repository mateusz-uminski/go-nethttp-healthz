package config

import "github.com/spf13/viper"

const (
	DEFAULT_HOST      = "127.0.0.1"
	DEFAULT_PORT      = 8080
	DEFAULT_LOG_LEVEL = "info"
)

type Config interface {
	GetHost() string
	GetPort() int
	GetLogLevel() string
}

type config struct {
	loader *viper.Viper
}

func New(prefix string) *config {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return &config{
		loader: v,
	}
}

func (c *config) GetHost() string {
	host := "host"
	c.loader.SetDefault(host, DEFAULT_HOST)
	return c.loader.GetString(host)
}

func (c *config) GetPort() int {
	port := "port"
	c.loader.SetDefault(port, DEFAULT_PORT)
	return c.loader.GetInt(port)
}

func (c *config) GetLogLevel() string {
	loglevel := "log_level"
	c.loader.SetDefault(loglevel, DEFAULT_LOG_LEVEL)
	return c.loader.GetString(loglevel)
}
