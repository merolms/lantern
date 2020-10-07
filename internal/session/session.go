package session

import (
    "time"

    "github.com/google/uuid"
)

type Session struct {
    ID         uuid.UUID `json:"id" db:"id"`
    EntityID   uuid.UUID `json:"entity_id" db:"entity_id"`
    EntityType uint8     `json:"-" db:"entity_type"`
    ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
    CreatedAt  time.Time `json:"-" db:"created_at"`
    UpdatedAt  time.Time `json:"-" db:"updated_at"`
}

func NewSession() *Session {
    var now = time.Now()
    return &Session{
        ID:         uuid.New(),
        EntityID:   uuid.UUID{},
        EntityType: 0,
        // TODO: Should make expires_at configurable
        ExpiresAt: now.Add(1 * time.Hour),
        CreatedAt: now,
        UpdatedAt: now,
    }
}

func (s *Session) WithEntityID(id uuid.UUID) *Session {
    if s != nil {
        s.EntityID = id
    }

    return s
}

func (s *Session) WithEntityType(et uint8) *Session {
    if s != nil {
        s.EntityType = et
    }

    return s
}

func (s Session) TableName() string {
    return "sessions"
}
