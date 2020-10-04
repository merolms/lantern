package configuration

import (
    "github.com/google/uuid"
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

func (v *ViperProvider) CSRFAuthKey() string {
    return viper.GetString("csrf.auth_key")
}

func (v *ViperProvider) PublicListenOn() string {
    return ":9090"
}

func (v *ViperProvider) PersisterDriverName() string {
    return "mysql"
}

func (v *ViperProvider) PersisterDSN() string {
    return viper.GetString("persister.dsn")
}

func (v *ViperProvider) DefaultSessionSecret() [][]byte {
    var secrets = viper.GetStringSlice("session.default_secret")
    if len(secrets) == 0 {
        secrets = []string{uuid.New().String()}
        viper.Set("session.default_secret", secrets)
    }

    var result = make([][]byte, len(secrets))
    for i, secret := range secrets {
        result[i] = []byte(secret)
    }

    return result
}

func (v *ViperProvider) SessionSecret() [][]byte {
    var secrets = viper.GetStringSlice("session.secrets")
    if len(secrets) == 0 {
        return v.DefaultSessionSecret()
    }

    var result = make([][]byte, len(secrets))
    for i, secret := range secrets {
        result[i] = []byte(secret)
    }

    return result
}
