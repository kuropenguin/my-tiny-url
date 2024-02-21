package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env string

const (
	EnvDev env = "dev"
	EnvPrd env = "prd"
)

func getEnv() env {
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

}
