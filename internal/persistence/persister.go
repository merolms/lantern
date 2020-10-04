package persistence

import "github.com/meroedu/lantern/internal/session"

type Persister interface {
    session.Persister
}
