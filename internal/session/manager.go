package session

import (
    "context"
    "net/http"
)

type (
    Manager interface {
        FetchFromRequest(context.Context, *http.Request) (*Session, error)
    }
    ManagerProvider interface {
        SessionManager() Manager
    }
)
