package configuration

var _ Provider = new(ViperProvider)

type ViperProvider struct{}

func NewViperProvider() *ViperProvider {
    return &ViperProvider{}
}

func (v *ViperProvider) PublicListenOn() string {
    return ":9090"
}
