package main

import (
	"math/rand"
	"time"
)

var (
	baseStr    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseStrLen = len(baseStr)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// TODO POST URL
	// TODO GET URL
	GenerateRandomString()
}

func GenerateRandomString(str string) string {
	randomNumber := rand.Intn(baseStrLen)
	return string(str[randomNumber])
}
