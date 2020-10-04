package sql

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/pkg/errors"

    "github.com/meroedu/lantern/internal/persistence"
    "github.com/meroedu/lantern/internal/session"
)

var _ persistence.Persister = new(Persister)

type Persister struct {
    db *sqlx.DB
}

func NewPersister(db *sqlx.DB) *Persister {
    return &Persister{db: db}
}

func (p *Persister) GetSession(ctx context.Context, sid uuid.UUID) (*session.Session, error) {
    var s session.Session
    var (
        query = fmt.Sprintf(`
        SELECT id, entity_id
          FROM %s
          WHERE id = ?
          LIMIT 1
    `, s.TableName())

        arg = sql.NullString{
            String: sid.String(),
            Valid:  true,
        }
    )

    if err := p.db.QueryRowxContext(ctx, query, arg).StructScan(&s); err != nil {
        return nil, errors.WithStack(err)
    }

    return &s, nil
}

func (p *Persister) CreateSession(ctx context.Context, session *session.Session) error {
    var (
        query = fmt.Sprintf(`
            INSERT INTO %s (id, entity_id, entity_type, expires_at, created_at, updated_at)
            VALUES (?, ?, ?, ?, ?, ?)
        `, session.TableName())
        // TODO: Implement arg builder
        args = []interface{}{
            sql.NullString{
                String: session.ID.String(),
                Valid:  true,
            },
            sql.NullString{
                String: session.EntityID.String(),
                Valid:  true,
            },
            sql.NullInt64{
                Int64: session.EntityType,
                Valid: true,
            },
            sql.NullTime{
                Time:  session.ExpiresAt,
                Valid: true,
            },
            sql.NullTime{
                Time:  session.CreatedAt,
                Valid: true,
            },
            sql.NullTime{
                Time:  session.UpdatedAt,
                Valid: true,
            },
        }
    )

    var _, err = p.db.ExecContext(ctx, query, args...)
    return err
}
