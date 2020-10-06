package session

import (
    "errors"
    "net/http"
)

type (
    HandlerProvider interface {
        SessionHandler() *Handler
    }
    handlerDependencies interface {
        ManagerProvider
    }
    Handler struct {
        d handlerDependencies
    }
)

func NewHandler(d handlerDependencies) *Handler {
    return &Handler{d: d}
}

func (handler *Handler) IsNotAuthenticated(onUnAuthenticated http.HandlerFunc, onAuthenticated http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if _, err := handler.d.SessionManager().FetchFromRequest(r.Context(), r); err != nil {
            if errors.Is(err, ErrNoActiveSessionFound) {
                if onUnAuthenticated != nil {
                    onUnAuthenticated(w, r)
                    return
                }
            }

            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            w.WriteHeader(http.StatusInternalServerError)
            _, _ = w.Write([]byte(`{"status": 500, "message": "An internal server error occurred. Please contact the system administrator."}`))
            return
        }

        if onAuthenticated != nil {
            onAuthenticated(w, r)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusForbidden)
        _, _ = w.Write([]byte(`{"status": 403, "message": "This endpoint can only be accessed without a login session. Please log out and try again."}`))
    }
}
