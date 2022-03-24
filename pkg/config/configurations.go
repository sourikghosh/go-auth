package config

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthConfig struct {
	PORT        string `mapstructure:"PORT"`
	Mode        string `mapstructure:"MODE"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	LogEncoding string `mapstructure:"LOG_ENCODING"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	DB_DSN      string `mapstructure:"DB_DSN"`
}

var defaultsValue = map[string]string{
	"PORT": "6969",
	"MODE": Development,
}

func LoadConfig(path string) (*AuthConfig, error) {
	// "" -> loads timezone as UTC:
	loc, err := time.LoadLocation("")
	if err != nil {
		return nil, err
	}

	time.Local = loc
	//Checks the defaults Value map and sets the default
	for key, val := range defaultsValue {
		viper.SetDefault(key, val)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg AuthConfig

	err = viper.Unmarshal(&cfg)

	if cfg.DB_DSN == "" || cfg.JWTSecret == "" {
		return nil, errors.New("invalid envs found")
	}

	if cfg.Mode == Production {
		gin.SetMode(gin.ReleaseMode)
	}

	return &cfg, err
}
