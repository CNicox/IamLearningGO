package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Argon2   Argon2Config
}

type Mode string

const (
	Debug   Mode = "debug"
	Release Mode = "release"
)

type EnableDisable string

const (
	Enable  EnableDisable = "enable"
	Disable EnableDisable = "disable"
)

type ServerConfig struct {
	Port int
	Mode Mode
}
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  EnableDisable
}
type JWTConfig struct {
	Secret string
	TTL    time.Duration
}
type Argon2Config struct {
	Memory      int
	Iterations  int
	Parallelism int
	SaltLength  int
	KeyLength   int
}

func handleConfigError(err error) {
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func LoadConfig() *Config {
	config := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("gin_web_api/config/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	handleConfigError(err)
	err = viper.Unmarshal(&config)
	handleConfigError(err)
	return &config
}
