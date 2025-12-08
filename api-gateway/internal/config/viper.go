package config

import (
	"log"
	"path/filepath"

	appLog "github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/log"
	"github.com/jordanmarcelino/learn-go-microservices/pkg/logger"
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

	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AutomaticEnv()

	configFound := false
	for _, p := range paths {
		viper.AddConfigPath(filepath.Clean(p))
		if err := viper.ReadInConfig(); err == nil {
			configFound = true
			log.Printf("Config file found at: %s", viper.ConfigFileUsed())
			break
		}
	}

	if !configFound {
		log.Printf("No .env file found, using system environment variables")
	}

	config = &Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshaling config: %v", err)
		return err
	}

	log.Printf("Config loaded: AppName=%s, Port=%d, Env=%s", config.AppName, config.RestApiPort, config.AppEnv)

	// Initialize logger with info level
	appLog.SetLogger(logger.NewZeroLogLogger(0))

	return nil
}
