package session

import (
    "context"
    "github.com/google/uuid"
    "github.com/pkg/errors"
    "net/http"

    "github.com/gorilla/sessions"
)

const DefaultCookieName = "LANTERN_SESSION"

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
        return nil, errors.WithStack(err)
    }

    var id, ok = s.Values["sid"].(string)
    if !ok {
        return nil, errors.New("request does not have a valid authentication session")
    }

    sid, err := uuid.Parse(id)
    if err != nil {
        return nil, errors.WithStack(err)
    }

    return manager.d.SessionPersister().GetSession(ctx, sid)
}
