package configuration

type Provider interface {
    CSRFAuthKey() string

    PersisterDriverName() string
    PersisterDSN() string

    PublicListenOn() string
}
