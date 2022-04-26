package secret

import (
	"context"
	"encoding/json"
	vault "github.com/hashicorp/vault/api"
	"github.com/igniter/config"
	"github.com/imdario/mergo"
	"github.com/mitchellh/mapstructure"
	"log"
)

type VaultClient struct {
	client  *vault.Client
	context context.Context
}

func createVaultClient(opt config.SecretsOptions, optOverrides interface{}) *VaultClient {

	finalOpt := make(map[string]interface{})
	mergo.Merge(&finalOpt, opt.Options)
	mergo.Merge(&finalOpt, optOverrides, mergo.WithOverride)

	var vaultOpt config.VaultOptions
	mapstructure.Decode(finalOpt, &vaultOpt)

	config := vault.DefaultConfig()

	if vaultOpt.Address != "" {
		config.Address = vaultOpt.Address
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	if vaultOpt.Token != "" {
		client.SetToken(vaultOpt.Token)
	}

	var vc VaultClient
	vc.client = client
	return &vc
}

func (v *VaultClient) GetSecret(path string) interface{} {

	// Reading a secret
	secret, err := v.client.Logical().Read(path)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	var dataMap interface{}
	data, _ := json.Marshal(secret)
	json.Unmarshal(data, &dataMap)
	return dataMap
}
