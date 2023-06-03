package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kuropenguin/my-tiny-url/app/entity"
	"github.com/kuropenguin/my-tiny-url/app/usecase"
)

type IHanlder interface {
	CreateTinyURL(w http.ResponseWriter, r *http.Request)
	GetOriginURLByTinyURL(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	usecase usecase.IUseCase
}

func NewHandlerImple(usecase usecase.IUseCase) *HandlerImpl {
	return &HandlerImpl{usecase: usecase}
}

type createTinyResponse struct {
	TinyURL string `json:"tiny_url"`
}
type createTinyRequest struct {
	OriginURL string `json:"origin_url"`
}

// TODO request validation
// curl -X POST -H "Content-Type: application/json" -d '{"origin_url":"google.com"}' localhost:8080/create_tiny_url
func (h *HandlerImpl) CreateTinyURL(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var createTinyRequest createTinyRequest
	if err := json.Unmarshal(reqBody, &createTinyRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	tinyURL, err := h.usecase.CreateTinyURL(entity.OriginURL(createTinyRequest.OriginURL))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	// write response 200 and json
	res := createTinyResponse{TinyURL: string(tinyURL)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

}

// TODO request validation
func (h *HandlerImpl) GetOriginURLByTinyURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}
