package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(".env don`t load")
	}

	return &Config{
		Db: DbConfig{
			DSN: os.Getenv("DSN"),
		},
	}
}
