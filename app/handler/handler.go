package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kuropenguin/my-tiny-url/app/entity"
	"github.com/kuropenguin/my-tiny-url/app/usecase"
)

type IHanlder interface {
	CreateTinyURL(w http.ResponseWriter, r *http.Request)
	GetOriginalURLByTinyURL(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	usecase usecase.IUseCase
}

func NewHandlerImple(usecase usecase.IUseCase) *HandlerImpl {
	return &HandlerImpl{usecase: usecase}
}

type createTinyRequest struct {
	OriginalURL string `json:"original_url"`
}

type createTinyResponse struct {
	TinyURL string `json:"tiny_url"`
}

type getOriginalURLResponse struct {
	OriginalURL string `json:"original_url"`
}

const (
	tinyURLKey = "tiny_url"
)

// curl -X GET -H localhost:8080/health
func (h *HandlerImpl) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

// TODO request validation
// curl -X POST -H "Content-Type: application/json" -d '{"original_url":"google.com"}' localhost:8080/create_tiny_url
func (h *HandlerImpl) CreateTinyURL(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var createTinyRequest createTinyRequest
	if err := json.Unmarshal(reqBody, &createTinyRequest); err != nil || createTinyRequest.OriginalURL == "" {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tinyURL, err := h.usecase.CreateTinyURL(entity.OriginalURL(createTinyRequest.OriginalURL))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// write response 200 and json
	res := createTinyResponse{TinyURL: string(tinyURL)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

}

// TODO request validation
// curl -X GET "localhost:8080/get_original_url?tiny_url=http://localhost:8080/BpLnfgDs"
func (h *HandlerImpl) GetOriginalURLByTinyURL(w http.ResponseWriter, r *http.Request) {
	tinyURL := r.URL.Query().Get(tinyURLKey)
	if tinyURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	originalURL, err := h.usecase.GetOriginalURLByTinyURL(entity.TinyURL(tinyURL))
	if err == usecase.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := getOriginalURLResponse{OriginalURL: string(originalURL)}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}
