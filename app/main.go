package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/kuropenguin/my-tiny-url/app/handler"
	"github.com/kuropenguin/my-tiny-url/app/repository"
	"github.com/kuropenguin/my-tiny-url/app/usecase"
)

func main() {
	router := mux.NewRouter()
	repository := repository.NewMapRepository()
	usecase := usecase.NewUsecaseImpl(repository)
	handler := handler.NewHandlerImple(usecase)
	router.HandleFunc("/create_tiny_url", handler.CreateTinyURL).Methods("POST")
	router.HandleFunc("/get_origin_url", handler.GetOriginURLByTinyURL).Methods("GET")

	log.Println("start server")
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
	// TODO graceful shutdown
}
