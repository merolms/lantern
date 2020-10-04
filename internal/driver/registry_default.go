package driver

import (
    "github.com/gorilla/mux"

    "github.com/meroedu/lantern/internal/driver/configuration"
    "github.com/meroedu/lantern/internal/http/health"
)

var _ Registry = new(RegistryDefault)

type RegistryDefault struct {
    c             configuration.Provider
    uploadHandler *health.Handler
}

func NewRegistryDefault() *RegistryDefault {
    return &RegistryDefault{}
}

func (rd *RegistryDefault) WithConfiguration(c configuration.Provider) *RegistryDefault {
    rd.c = c

    return rd
}

func (rd *RegistryDefault) Configuration() configuration.Provider {
    return rd.c
}

func (rd *RegistryDefault) RegisterRoutes(router *mux.Router) {
    rd.HealthHandler().RegisterRoutes(router)
}

func (rd *RegistryDefault) HealthHandler() *health.Handler {
    if rd.uploadHandler == nil {
        rd.uploadHandler = health.NewHandler()
    }
    return rd.uploadHandler
}
