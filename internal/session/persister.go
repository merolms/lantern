package session

import (
    "context"

    "github.com/google/uuid"
)

type (
    PersisterProvider interface {
        SessionPersister() Persister
    }
    Persister interface {
        // GetSession retrieves a session from the store
        GetSession(context.Context, uuid.UUID) (*Session, error)

        // CreateSession adds a session to the store
        CreateSession(context.Context, *Session) error
    }
)
