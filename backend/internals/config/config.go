package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
}

func Load() (Config, error) {
	settings := viper.New()
	settings.SetConfigFile(".env")
	settings.SetConfigType("env")
	settings.AutomaticEnv()
	if err := settings.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return Config{}, err
		}
	}

	databaseURL := settings.GetString("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}
	return Config{DatabaseURL: databaseURL}, nil
}
