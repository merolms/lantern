package persistence

import (
    "github.com/meroedu/lantern/internal/entity"
    "github.com/meroedu/lantern/internal/session"
)

type Persister interface {
    entity.Persister
    session.Persister
}
