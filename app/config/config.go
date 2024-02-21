package config

import "os"

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
