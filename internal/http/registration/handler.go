package registration

import (
    "fmt"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
    "github.com/meroedu/lantern/internal/session"
    "net/http"
)

type (
    handlerDependencies interface {
        session.HandlerProvider
    }
    Handler struct {
        d handlerDependencies
    }
)

func NewHandler(d handlerDependencies) *Handler {
    return &Handler{d: d}
}

func (handler *Handler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/register", handler.d.SessionHandler().IsNotAuthenticated(handler.handleRegistration, nil))
}

func (handler *Handler) handleRegistration(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement handleRegistration
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte(fmt.Sprintf(`{"status": 201, "identity": {"id": %s}}`, uuid.New().String())))
}
