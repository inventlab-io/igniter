package secret

import (
	"context"
	vault "github.com/hashicorp/vault/api"
	"github.com/igniter/config"
	"github.com/mitchellh/mapstructure"
	"log"
)

type VaultClient struct {
	client  *vault.Client
	context context.Context
}

func createVaultClient(opt config.SecretsOptions) VaultClient {
	var vaultOpt config.VaultOptions
	mapstructure.Decode(opt.Options, &vaultOpt)

	config := vault.DefaultConfig()
	config.Address = vaultOpt.Address

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	if vaultOpt.Token != "" {
		client.SetToken(vaultOpt.Token)
	}

	var vc VaultClient
	vc.client = client
	return vc
}

func (v *VaultClient) GetSecret(path string) map[string]interface{} {

	// Reading a secret
	secret, err := v.client.Logical().Read(path)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("data type assertion failed: %T %#v", secret.Data["data"], secret.Data["data"])
	}
	return data
}
