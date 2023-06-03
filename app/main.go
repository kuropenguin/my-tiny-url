package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuropenguin/my-tiny-url/app/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create_tiny_url", handler.CreateTinyURL).Methods("POST")
	router.HandleFunc("/get_origin_url", handler.GetOriginURLByTinyURL).Methods("GET")

	http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
	// TODO graceful shutdown
}
