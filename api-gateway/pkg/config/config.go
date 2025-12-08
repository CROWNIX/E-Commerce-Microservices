package config

import (
	// "log"
	"os"
	"path/filepath"
	// "github.com/spf13/viper"
)

var (
	JWT_CONFIG  *JwtConfig
	SMTP_CONFIG *SMTPConfig
)

func parseConfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd)
}
