package hash

type Hasher interface {
    Generate(password []byte) ([]byte, error)
}

type Provider interface {
    Hasher() Hasher
}
