package configuration

type Provider interface {
    PublicListenOn() string
}