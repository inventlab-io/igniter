package storage

import (
	"encoding/json"
	"github.com/igniter/config"
)

type ConfigRepoFactory func(cfg config.ServerConfig) ConfigRepo

type TemplateStore interface {
	PutTemplate(path string, template string) string
	GetTemplate(path string) string
}

type ConfigRepo interface {
	GetStoreOptions(path string) []byte
	PutStoreOptions(path string, optionsJson string) string
}

func GetTemplateStore(settingsJson []byte) TemplateStore {

	var opt config.StoreOptions
	json.Unmarshal(settingsJson, &opt)
	var store TemplateStore

	if opt.StorageType == "etcd" {
		var etcdOpt config.EtcdOptions
		json.Unmarshal(settingsJson, &etcdOpt)
		json.Unmarshal(settingsJson, &etcdOpt)
		store = etcdInitStore(etcdOpt)
	}

	return store
}
