package sql

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/meroedu/lantern/internal/entity"
)

func (p *Persister) CreateEntity(ctx context.Context, entity *entity.Entity) error {
    var (
        query = fmt.Sprintf(`
            INSERT INTO %s (id, entity_type, email, username, status, hashed_password, created_at, updated_at)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        `, entity.TableName())
        args = []interface{}{
            sql.NullString{
                String: entity.ID.String(),
                Valid:  true,
            },
            sql.NullInt32{
                Int32: int32(entity.Type),
                Valid: true,
            },
            sql.NullString{
                String: entity.Email,
                Valid:  true,
            },
            sql.NullString{
                String: entity.Username,
                Valid:  true,
            },
            sql.NullInt32{
                Int32: int32(entity.Status),
                Valid: true,
            },
            sql.NullString{
                String: string(entity.HashedPassword),
                Valid:  true,
            },
            sql.NullTime{
                Time:  entity.CreatedAt,
                Valid: true,
            },
            sql.NullTime{
                Time:  entity.UpdatedAt,
                Valid: true,
            },
        }
    )

    var _, err = p.db.ExecContext(ctx, query, args...)
    return err
}
