package storage

import (
	"github.com/igniter/config"
)

type ConfigRepoFactory func(cfg config.ServerConfig) ConfigRepo

type TemplateStore interface {
	PutTemplate(path string, template string) string
	GetTemplate(path string) string
	DeleteTemplate(path string) string
}

type ValuesStore interface {
	PutValues(path string, value string) string
	GetValues(path string) string
	GetValuesInBatch(paths []string) map[string]string
	DeleteValues(path string) string
}

type SecretsMapStore interface {
	PutSecretsMap(engine string, path string, value string) string
	GetSecretsMap(engine string, path string) string
	DeleteSecretsMap(engine string, path string) string
}

type ConfigRepo interface {
	GetStoreOptions(path string) []byte
	PutStoreOptions(path string, optionsJson string) string
	DeleteStoreOptions(path string) string
	GetSecretsOptions(path string) []byte
	PutSecretsOptions(path string, optionsJson string) string
	DeleteSecretsOptions(path string) string
}

func GetTemplateStore(opt config.StoreOptions) TemplateStore {

	var store TemplateStore

	if opt.Type == "etcd" {
		store = createEtcdStore(opt)
	}

	return store
}

func GetValuesStore(opt config.StoreOptions) ValuesStore {

	var store ValuesStore

	if opt.Type == "etcd" {
		store = createEtcdStore(opt)
	}

	return store
}

func GetSecretsMapStore(opt config.StoreOptions) SecretsMapStore {

	var store SecretsMapStore

	if opt.Type == "etcd" {
		store = createEtcdStore(opt)
	}

	return store
}
