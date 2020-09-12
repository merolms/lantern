package health

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
    router.
        HandleFunc("/liveness", handler.Liveness()).
        Methods("GET")
}

func (handler *Handler) Liveness() func(http.ResponseWriter, *http.Request) {
    return func(writer http.ResponseWriter, request *http.Request) {
        writer.WriteHeader(http.StatusOK)
        writer.Write([]byte(`{code: "success", "message": "ok"}`))
    }
}
