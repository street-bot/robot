package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Init initializes viper configuration based on the environment variable "ENV"
func Init() (*viper.Viper, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, fmt.Errorf("ENV variable not set")
	}
	return InitEnv(env)
}

// InitEnv initializes configuration based on input env
func InitEnv(env string) (*viper.Viper, error) {
	log.Info().Msgf("Running in %s mode!", env)
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	return config, nil
}
