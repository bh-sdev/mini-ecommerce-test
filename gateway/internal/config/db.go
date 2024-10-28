package config

import "fmt"

type DB struct {
	DATABASE_URL string `env:"DATABASE_URL"`
	DBHost       string `env:"DB_HOST" envDefault:"localhost"`
	DBPort       string `env:"DB_PORT" envDefault:"5432"`
	DBUser       string `env:"DB_HOST" envDefault:"postgres"`
	DBPassword   string `env:"DB_HOST" envDefault:"safepassword"`
	DBName       string `env:"DB_HOST" envDefault:"postgres"`
	//
	AutoMigrate bool `env:"DB_AUTO_MIGRATE" envDefault:"true"`
}

func (conf *DB) Dsn() string {
	if conf.DATABASE_URL != "" {
		return conf.DATABASE_URL
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DBHost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
		// viper.GetString("DB_HOST"),
		// viper.GetString("DB_USER"),
		// viper.GetString("DB_PASSWORD"),
		// viper.GetString("DB_NAME"),
		// viper.GetString("DB_PORT"),
	)
	return dsn
}
