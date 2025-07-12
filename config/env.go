// config/env.go
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string        `mapstructure:"PORT"`
	MongoURI    string        `mapstructure:"MONGO_URI"`
	MongoDBName string        `mapstructure:"MONGO_DB"`
	JWTSecret   string        `mapstructure:"JWT_SECRET"`
	Timeout     time.Duration `mapstructure:"TIMEOUT"`
	Environment string        `mapstructure:"ENV"` // dev, prod, test, local
}

var Constants *Config

// Load reads configuration from .env or environment variables.
func Load() (bool, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("MONGO_URI")
	viper.BindEnv("MONGO_DB")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("TIMEOUT")
	viper.BindEnv("ENV")

	// Set sensible defaults
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("TIMEOUT", 30)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return false, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return false, fmt.Errorf("unable to decode config: %w", err)
	}

	// Convert seconds to duration
	cfg.Timeout = cfg.Timeout * time.Second

	// Set the constants
	Constants = &cfg
	return true, nil
}
