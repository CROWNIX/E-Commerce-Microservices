package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	paths := []string{
		".",
		"..",
		"../..",
		"../../..",
		"../../../..",
		"../../../../..",
	}

	viper.SetConfigName("env.json")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	for _, p := range paths {
		viper.AddConfigPath(filepath.Clean(p))
		if err := viper.ReadInConfig(); err == nil {
			if err := viper.Unmarshal(&config); err != nil {
				return err
			}
			return nil
		}
	}

	log.Printf("No env.json file found, fallback to system environment variables")
	return nil
}
