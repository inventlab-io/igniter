package storage

import "github.com/igniter/config"

type TemplateStore interface {
	PutTemplate(path string, template string) string
	GetTemplate(path string) string
}

type ConfigRepo interface {
	GetTemplateStoreOptions(path string) string
	PutTemplateStoreOptions(path string, optionsJson string) string
	GetTemplateStore(path string) TemplateStore
}

type ConfigRepoFactory func(cfg config.ServerConfig) ConfigRepo
