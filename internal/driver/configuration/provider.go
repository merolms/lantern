package configuration

type Provider interface {
    CSRFAuthKey() string

    PublicListenOn() string
}
