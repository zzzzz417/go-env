package config

import (
	"log"

	"github.com/joho/godotenv"
)

type config struct {
	numInt   int
	numFloat float64
	flag     bool
	str      string
}

var Setting = &config{}

func Setup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
