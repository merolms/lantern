package driver

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"

    "github.com/meroedu/lantern/internal/driver/configuration"
    "github.com/meroedu/lantern/x/hash"
)

type Registry interface {
    Init() error

    Configuration() configuration.Provider

    CookieManager() sessions.Store

    RegisterRoutes(*mux.Router)

    hash.Provider
}
