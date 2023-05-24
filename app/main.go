package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	baseStr    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseStrLen = len(baseStr)
	urlLen     = 8
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// TODO POST URL
	// TODO GET URL
	url := GenerateRandomString(urlLen)
	fmt.Println(url)
}

func GenerateRandomString(loopNum int) string {
	url := ""
	for i := 0; i < loopNum; i++ {
		url += GenerateRandomChar(baseStr)
	}
	return url
}

func GenerateRandomChar(str string) string {
	randomNumber := rand.Intn(baseStrLen)
	return string(str[randomNumber])
}
