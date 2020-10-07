package registration

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/meroedu/lantern/internal/entity"
    "github.com/meroedu/lantern/internal/session"
    "github.com/meroedu/lantern/x/hash"
    "github.com/pkg/errors"
    "log"
    "net/http"
)

type (
    BodyPayload struct {
        Password string `json:"password"`
        Traits   struct {
            Email    string `json:"email,omitempty"`
            Username string `json:"username,omitempty"`
        } `json:"traits"`
    }
    handlerDependencies interface {
        session.HandlerProvider
        hash.Provider
        entity.PersistenceProvider
        session.PersistenceProvider
    }
    Handler struct {
        d handlerDependencies
    }
)

func (payload *BodyPayload) validate() error {
    if payload.Password == "" {
        return errors.New("invalid input: password should not be empty")
    }

    if payload.Traits.Email == "" && payload.Traits.Username == "" {
        return errors.New("invalid input: traits should not be empty")
    }

    return nil
}

func NewHandler(d handlerDependencies) *Handler {
    return &Handler{d: d}
}

func (handler *Handler) RegisterRoutes(r *mux.Router) {
    r.
        HandleFunc("/register", handler.d.SessionHandler().IsNotAuthenticated(handler.handleRegistration, nil)).
        Methods("POST")
}

func (handler *Handler) handleRegistration(w http.ResponseWriter, r *http.Request) {
    var payload BodyPayload
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        // TODO: Implement JSON writer reduce duplicate line of code
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusBadRequest)
        _, _ = w.Write([]byte(`{"status": 400, "message": "The request was malformed or contained invalid parameters"}`))
        return
    }

    if err := payload.validate(); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusBadRequest)
        _, _ = w.Write([]byte(`{"status": 400, "message": "The request was malformed or contained invalid parameters"}`))
        return
    }

    var hpw, err = handler.d.Hasher().Generate([]byte(payload.Password))
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte(`{"status": 500, "message": "An internal server error occurred, please contact the system administrator"}`))
        return
    }

    var i = entity.
        NewIdentity().
        WithCredentialType(entity.CredentialTypePassword).
        WithHashedPassword(hpw).
        WithEmail(payload.Traits.Email)

    err = handler.d.IdentityPersister().CreateEntity(context.Background(), i)
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte(`{"status": 500, "message": "An internal server error occurred, please contact the system administrator"}`))
        log.Println("create entity: ", err)
        return
    }

    var s = session.
        NewSession().
        WithEntityID(i.ID).
        WithEntityType(uint8(entity.TypeUser))
    err = handler.d.SessionPersister().CreateSession(r.Context(), s)
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte(`{"status": 500, "message": "An internal server error occurred, please contact the system administrator"}`))
        log.Println("create session: ", err)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte(fmt.Sprintf(`{"status": 201, "identity": {"id": "%s"}}`, i.ID.String())))
}
