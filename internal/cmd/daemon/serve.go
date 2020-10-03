package daemon

import (
    "log"
    "net/http"

    "github.com/gorilla/csrf"
    "github.com/gorilla/mux"

    "github.com/meroedu/lantern/internal/driver"
    "github.com/meroedu/lantern/x/graceful"
)

func Serve(registry driver.Registry) {
    var (
        router = mux.NewRouter()
        CSRF   = csrf.Protect([]byte(registry.Configuration().CSRFAuthKey()))
    )

    registry.RegisterRoutes(router)

    var server = http.Server{
        Addr:    registry.Configuration().PublicListenOn(),
        Handler: CSRF(router),
    }

    if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
        log.Fatalf("Failed to gracefully shutdown Lantern: %s", err.Error())
    }

    log.Println("Lantern was shutdown gracefully")
}
