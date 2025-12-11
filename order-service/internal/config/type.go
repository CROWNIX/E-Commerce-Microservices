package config

import "time"

type Config struct {
	AppName     string      `mapstructure:"APP_NAME"`
	RestApiPort int         `mapstructure:"REST_API_PORT"`
	AppEnv      string      `mapstructure:"APP_ENV"`
	GinMode     string      `mapstructure:"GIN_MODE"`
	DB          ConfigDB    `mapstructure:"DB"`
	JWT   		JWTConfig   `mapstructure:"JWT"`
	Redis   	RedisConfig `mapstructure:"REDIS"`
}

type ConfigDB struct {
	DSN             string        `mapstructure:"DB_DSN"`
	MaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
	ConnMaxIdleTime time.Duration `mapstructure:"DB_CONN_MAX_IDLE_TIME"`
}

type JWTConfig struct {
	Secret	string	`mapstructure:"JWT_SECRET"`
}

type RedisConfig struct {
	Address	string		`mapstructure:"ADDRESS"`
	Db	    int	    	`mapstructure:"DB"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		return &Config{}
	}

	return config
}
