package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser   string
	DBPasswd string
	DBAddr   string
	DBName   string
}

func Load() Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	return Config{
		DBUser:   os.Getenv("DATABASE_USER"),
		DBPasswd: os.Getenv("DATABASE_PASS"),
		DBAddr:   os.Getenv("DATABASE_HOST"),
		DBName:   os.Getenv("DATABASE_NAME"),
	}
}
