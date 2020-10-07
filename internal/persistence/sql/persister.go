package sql

import (
    "github.com/jmoiron/sqlx"

    "github.com/meroedu/lantern/internal/persistence"
)

var _ persistence.Persister = new(Persister)

type Persister struct {
    db *sqlx.DB
}

func NewPersister(db *sqlx.DB) *Persister {
    return &Persister{db: db}
}
