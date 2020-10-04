package driver

import (
    "github.com/gorilla/mux"
    "github.com/meroedu/lantern/internal/driver/configuration"
)

type Registry interface {
    Init() error

    Configuration() configuration.Provider

    RegisterRoutes(*mux.Router)
}
