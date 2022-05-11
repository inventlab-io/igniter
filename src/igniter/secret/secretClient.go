package secret

import (
	"github.com/igniter/config"
)

type SecretClient interface {
	GetSecret(path string) interface{}
}

func GetSecretClient(opt config.SecretsOptions, overrides interface{}) SecretClient {

	var client SecretClient

	if opt.Type == "vault" {
		client = createVaultClient(opt, overrides)
	}

	return client
}
