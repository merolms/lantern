package configuration

import (
    "strings"

    "github.com/spf13/viper"
)

var _ Provider = new(ViperProvider)

func init() {
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    viper.SetEnvPrefix("LANTERN")
}

type ViperProvider struct{}

func NewViperProvider() *ViperProvider {
    return &ViperProvider{}
}

func (v *ViperProvider) PublicListenOn() string {
    return ":9090"
}

func (v *ViperProvider) CSRFAuthKey() string {
    return viper.GetString("csrf.auth_key")
}
