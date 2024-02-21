package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type appEnv string

type config struct {
	SecretKey string `env:"SECRET_KEY,expand"`
}

const (
	EnvDev appEnv = "dev"
	EnvPrd appEnv = "prd"
)

func getEnv() appEnv {
	switch env := os.Getenv("ENV"); env {
	case "prd":
		return EnvPrd
	default:
		return EnvDev
	}
}

func IsDev() bool {
	return getEnv() == EnvDev
}

func Load() {
	if IsDev() {
		err := godotenv.Load("../../env/.env.dev")
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file: %v", err))
		}
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
}
