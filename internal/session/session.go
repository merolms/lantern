package session

import (
    "github.com/google/uuid"
    "time"
)

type Session struct {
    ID         uuid.UUID `json:"id" db:"id"`
    EntityID   uuid.UUID `json:"entity_id" db:"entity_id"`
    EntityType int64     `json:"-" db:"entity_type"`
    ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
    CreatedAt  time.Time `json:"-" db:"created_at"`
    UpdatedAt  time.Time `json:"-" db:"updated_at"`
}

func (s Session) TableName() string {
    return "sessions"
}
