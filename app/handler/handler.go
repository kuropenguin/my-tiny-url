package handler

import (
	"fmt"
	"net/http"
)

// TODO request validation

func CreateTinyURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

func GetOriginURLByTinyURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}
