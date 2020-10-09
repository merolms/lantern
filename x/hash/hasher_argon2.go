package hash

import (
    "bytes"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "github.com/pkg/errors"

    "golang.org/x/crypto/argon2"
)

var _ Hasher = new(Argon2)

type Argon2 struct{}

func NewArgon2() *Argon2 {
    return &Argon2{}
}

func (a *Argon2) Generate(password []byte) ([]byte, error) {
    // TODO: Salt length should be configurable
    var salt = make([]byte, 20)
    if _, err := rand.Read(salt); err != nil {
        return nil, err
    }

    // TODO: IDKey params should be configurable
    var hash = argon2.IDKey(password, salt, 3, 131072, 1, 32)
    var b bytes.Buffer
    if _, err := fmt.Fprintf(
        &b,
        "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
        argon2.Version, 131072, 3, 1,
        base64.RawStdEncoding.EncodeToString(salt),
        base64.RawStdEncoding.EncodeToString(hash),
    ); err != nil {
        return nil, errors.WithStack(err)
    }

    return b.Bytes(), nil
}
