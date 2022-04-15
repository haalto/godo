package config

import "os"

type Database struct {
	DSN string
}

type Config struct {
	Database
}

func GetConfig() *Config {

	conf := Config{Database: Database{
		DSN: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	}}

	if os.Getenv("DATABASE_DSN") != "" {
		conf.Database.DSN = os.Getenv("DATABASE_DSN")
	}

	return &conf
}
