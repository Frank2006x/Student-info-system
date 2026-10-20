package config

import (
	"errors"
	"io/fs"

	"github.com/spf13/viper"
)

type Config struct {
	PrimaryDatabaseURL string
	ReplicaDatabaseURL string
}

func Load() (Config, error) {
	settings := viper.New()
	settings.SetConfigFile(".env")
	settings.SetConfigType("env")
	settings.AutomaticEnv()
	if err := settings.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) && !errors.Is(err, fs.ErrNotExist) {
			return Config{}, err
		}
	}

	primaryDatabaseURL := settings.GetString("PRIMARY_DATABASE_URL")
	replicaDatabaseURL := settings.GetString("REPLICA_DATABASE_URL")
	if primaryDatabaseURL == "" || replicaDatabaseURL == "" {
		return Config{}, errors.New("PRIMARY_DATABASE_URL and REPLICA_DATABASE_URL are required")
	}
	return Config{PrimaryDatabaseURL: primaryDatabaseURL, ReplicaDatabaseURL: replicaDatabaseURL}, nil
}
