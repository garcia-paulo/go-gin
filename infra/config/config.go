package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource      string        `mapstructure:"DB_SOURCE"`
	ServerPort    string        `mapstructure:"SERVER_PORT"`
	TokenKey      string        `mapstructure:"TOKEN_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err.Error())
	}

	return config
}
