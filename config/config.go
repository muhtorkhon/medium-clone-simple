package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")
	
	conf := viper.New()

	conf.AutomaticEnv()

	cfg := Config{
		Postgres: Postgres{
			Host: conf.GetString("HOST"),
			Port: conf.GetString("PORT"),
			User: conf.GetString("USER"),
			Password: conf.GetString("PASSWORD"),
			Database: conf.GetString("DB_NAME"),
		},
	}
	return cfg
}