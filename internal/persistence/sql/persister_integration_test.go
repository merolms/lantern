// +build integration

package sql

import (
    "context"
    "testing"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"

    "github.com/meroedu/lantern/internal/session"
    "github.com/stretchr/testify/require"
)

func TestPersister(t *testing.T) {
    var db, err = dbsql.Open("mysql", os.Getenv("persister_dsn"))
    require.NoError(t, err)

    err = db.Ping()
    require.NoError(t, err)

    var (
        dbx       = sqlx.NewDb(db, "mysql")
        persister = NewPersister(dbx)
        sid       = uuid.New()
        now       = time.Now()
        s         = session.Session{
            ID:         sid,
            EntityID:   uuid.New(),
            EntityType: 1,
            ExpiresAt:  now,
            CreatedAt:  now,
            UpdatedAt:  now,
        }
    )

    err = persister.CreateSession(context.Background(), &s)
    require.NoError(t, err)

    sess, err := persister.GetSession(context.Background(), sid)
    require.NoError(t, err)
    require.NotNil(t, sess)
}
