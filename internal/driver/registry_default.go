package driver

import (
    dbsql "database/sql"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "github.com/jmoiron/sqlx"
    "github.com/pkg/errors"

    "github.com/meroedu/lantern/internal/driver/configuration"
    "github.com/meroedu/lantern/internal/entity"
    "github.com/meroedu/lantern/internal/http/health"
    "github.com/meroedu/lantern/internal/http/registration"
    "github.com/meroedu/lantern/internal/persistence"
    "github.com/meroedu/lantern/internal/persistence/sql"
    "github.com/meroedu/lantern/internal/session"
    "github.com/meroedu/lantern/x/hash"
)

var _ Registry = new(DefaultRegistry)

type DefaultRegistry struct {
    c configuration.Provider

    persister persistence.Persister

    sessionManager session.Manager

    cookieManager *sessions.CookieStore

    passwordHasher hash.Hasher

    healthHandler       *health.Handler
    registrationHandler *registration.Handler
    sessionHandler      *session.Handler
}

func (r *DefaultRegistry) Hasher() hash.Hasher {
    if r.passwordHasher == nil {
        r.passwordHasher = hash.NewArgon2()
    }

    return r.passwordHasher
}

func (r *DefaultRegistry) IdentityPersister() entity.Persister {
    return r.persister
}

func NewRegistryDefault() *DefaultRegistry {
    return &DefaultRegistry{}
}

func (r *DefaultRegistry) Init() error {
    if r.persister != nil {
        return errors.New("default registry init: must not be called more than once")
    }

    var db, err = dbsql.Open(r.c.PersisterDriverName(), r.c.PersisterDSN())
    if err != nil {
        return errors.WithStack(err)
    }

    if err := db.Ping(); err != nil {
        return errors.WithStack(err)
    }

    var dbx = sqlx.NewDb(db, r.c.PersisterDriverName())
    r.persister = sql.NewPersister(dbx)

    return nil
}

func (r *DefaultRegistry) WithConfiguration(c configuration.Provider) *DefaultRegistry {
    r.c = c

    return r
}

func (r *DefaultRegistry) Configuration() configuration.Provider {
    return r.c
}

func (r *DefaultRegistry) CookieManager() sessions.Store {
    if r.cookieManager == nil {
        var cs = sessions.NewCookieStore(r.c.SessionSecret()...)
        cs.Options.Secure = true
        cs.Options.HttpOnly = true
        cs.Options.SameSite = http.SameSiteLaxMode

        r.cookieManager = cs
    }

    return r.cookieManager
}

func (r *DefaultRegistry) SessionManager() session.Manager {
    if r.sessionManager == nil {
        r.sessionManager = session.NewHTTPManager(r)
    }

    return r.sessionManager
}

func (r *DefaultRegistry) SessionPersister() session.Persister {
    return r.persister
}

func (r *DefaultRegistry) RegisterRoutes(router *mux.Router) {
    r.HealthHandler().RegisterRoutes(router)
    r.RegistrationHandler().RegisterRoutes(router)
}

func (r *DefaultRegistry) HealthHandler() *health.Handler {
    if r.healthHandler == nil {
        r.healthHandler = health.NewHandler()
    }

    return r.healthHandler
}

func (r *DefaultRegistry) RegistrationHandler() *registration.Handler {
    if r.registrationHandler == nil {
        r.registrationHandler = registration.NewHandler(r)
    }

    return r.registrationHandler
}

func (r *DefaultRegistry) SessionHandler() *session.Handler {
    if r.sessionHandler == nil {
        r.sessionHandler = session.NewHandler(r)
    }

    return r.sessionHandler
}
