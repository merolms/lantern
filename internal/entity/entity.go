package entity

import (
    "time"

    "github.com/google/uuid"
)

type (
    Type   uint8
    Status uint8
)

const (
    TypeUser Type = 1
)

const (
    StatusInactive Status = 0
    StatusActive   Status = 1
)

type Entity struct {
    ID             uuid.UUID      `json:"id" db:"id"`
    CredentialType CredentialType `json:"credential_type" db:"credential_type"`
    Email          string         `json:"email" db:"email"`
    HashedPassword []byte         `json:"hashed_password" db:"hashed_password"`
    Status         Status         `json:"status" db:"status"`
    Type           Type           `json:"type" db:"type"`
    Username       string         `json:"username" db:"username"`
    CreatedAt      time.Time      `json:"created_at" db:"created_at"`
    UpdatedAt      time.Time      `json:"updated_at" db:"updated_at"`
}

func NewIdentity() *Entity {
    var now = time.Now()
    return &Entity{
        ID:             uuid.New(),
        CredentialType: "",
        Email:          "",
        HashedPassword: nil,
        Status:         1,
        Type:           TypeUser,
        Username:       "",
        CreatedAt:      now,
        UpdatedAt:      now,
    }
}

func (e *Entity) WithCredentialType(ct CredentialType) *Entity {
    if e != nil {
        e.CredentialType = ct
    }

    return e
}

func (e *Entity) WithEmail(email string) *Entity {
    if e != nil {
        e.Email = email
    }

    return e
}

func (e *Entity) WithHashedPassword(hash []byte) *Entity {
    if e != nil {
        e.HashedPassword = hash
    }

    return e
}

func (e *Entity) WithUsername(username string) *Entity {
    if e != nil {
        e.Username = username
    }

    return e
}

func (e Entity) TableName() string {
    return "entities"
}
