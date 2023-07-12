package entity

import "math/rand"

const (
	baseURL    = "http://localhost:8080/"
	baseStr    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseStrLen = len(baseStr)
	tinyLen    = 8
)

type TinyURL string

type OriginalURL string

func GenerateTinyURL() TinyURL {
	return TinyURL(baseURL + generateRandomString(tinyLen))
}

func generateRandomString(loopNum int) string {
	url := ""
	for i := 0; i < loopNum; i++ {
		url += generateRandomChar(baseStr)
	}
	return url
}

func generateRandomChar(str string) string {
	randomNumber := rand.Intn(baseStrLen)
	return string(str[randomNumber])
}
