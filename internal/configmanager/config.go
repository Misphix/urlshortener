package configmanager

import (
	"time"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	HTTPServer HTTPServerConfig
	Database   DatabaseConfig
	Redis      RedisConfig
	ShortURL   ShortURLConfig
}

type HTTPServerConfig struct {
	Domain string
	Port   uint
}

type DatabaseConfig struct {
	DSN string
}

type RedisConfig struct {
	Address     string
	DialTimeout time.Duration
	Expiration  time.Duration
}

type ShortURLConfig struct {
	ExpireDuration time.Duration
}

func Get() (*Config, error) {
	if config == nil {
		c, err := get()
		if err != nil {
			return nil, err
		}
		config = c
	}
	return config, nil
}

func get() (*Config, error) {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
