package config

import "time"

type Auth struct {
	JWTConfig
	// JWTSecret string `env:"AUTH_JWT_SECREET" envDefault:"ChangeMe!"`
}

type JWTConfig struct {
	SigningKey string        `env:"AUTH_JWT_SECREET" envDefault:"ChangeMe!"`
	ExpiresIn  time.Duration `env:"AUTH_JWT_DURATION" envDefault:"24h"`
}
