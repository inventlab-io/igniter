package storage

import "github.com/igniter/server/config"

type TemplateStore interface {
	PutTemplate(path string, template string) string
	GetTemplate(path string) string
}

type StoreFactory func(cfg config.ServerConfig) TemplateStore
