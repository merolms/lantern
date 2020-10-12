package entity

import "context"

type (
    PersistenceProvider interface {
        IdentityPersister() Persister
    }
    Persister interface {
        CreateEntity(context.Context, *Entity) error
    }
)
