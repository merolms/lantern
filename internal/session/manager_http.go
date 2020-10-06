package session

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/pkg/errors"
    "net/http"

    "github.com/gorilla/sessions"
)

const DefaultCookieName = "LANTERN_SESSION"

var ErrNoActiveSessionFound = errors.New("request does not have valid authentication session")

var _ Manager = new(HTTPManager)

type (
    CookieProvider interface {
        CookieManager() sessions.Store
    }
    managerDependencies interface {
        CookieProvider
        PersistenceProvider
    }
    HTTPManager struct {
        d managerDependencies

        cookieName string
    }
)

func NewHTTPManager(d managerDependencies) *HTTPManager {
    return &HTTPManager{d: d, cookieName: DefaultCookieName}
}

func (manager *HTTPManager) FetchFromRequest(ctx context.Context, r *http.Request) (*Session, error) {
    var s, err = manager.d.CookieManager().Get(r, manager.cookieName)
    if err != nil {
        // TODO: Should have more information to debug here
        return nil, errors.WithStack(ErrNoActiveSessionFound)
    }

    var id, ok = s.Values["sid"].(string)
    if !ok {
        return nil, errors.WithStack(ErrNoActiveSessionFound)
    }

    sid, err := uuid.Parse(id)
    if err != nil {
        return nil, errors.WithStack(err)
    }

    sess, err := manager.d.SessionPersister().GetSession(ctx, sid)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrNoActiveSessionFound
        }

        return nil, errors.WithStack(err)
    }

    return sess, nil
}
