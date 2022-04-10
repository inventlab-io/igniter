package storage

import (
	"github.com/igniter/config"
	"github.com/mitchellh/mapstructure"
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

type ConfigRepo interface {
	GetStoreOptions(path string) []byte
	PutStoreOptions(path string, optionsJson string) string
	DeleteStoreOptions(path string) string
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

func createEtcdStore(opt config.StoreOptions) *EtcdStore {
	var etcdOpt config.EtcdOptions
	mapstructure.Decode(opt.Options, &etcdOpt)
	store := etcdInitStore(etcdOpt)
	return store
}
