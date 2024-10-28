package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/spf13/viper"
)

type Config struct {
	App
	DB
	Auth
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		// return nil, err
	}
	viper.AutomaticEnv()

	// config := &Config{
	// 	DBHost:     viper.GetString("DB_HOST"),
	// 	DBPort:     viper.GetString("DB_PORT"),
	// 	DBUser:     viper.GetString("DB_USER"),
	// 	DBPassword: viper.GetString("DB_PASSWORD"),
	// 	DBName:     viper.GetString("DB_NAME"),
	// 	AppPort:    viper.GetString("APP_PORT"),
	// 	JWTSecret:  viper.GetString("JWT_SECRET"),
	// }

	var config Config
	// err := viper.Unmarshal(&config)
	err := env.Parse(&config)

	return &config, err
}
