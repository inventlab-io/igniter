package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/igniter/config"
	"github.com/igniter/http"
	"github.com/igniter/storage"
)

type Server struct {
	config            config.ServerConfig
	configRepoFactory storage.ConfigRepoFactory
}

func (svr Server) Run(cfg config.ServerConfig) {

	svr.config = cfg

	if cfg.Storage.Type == "etcd" {
		svr.configRepoFactory = storage.EtcdConfigRepoFactory
	}

	http.InitGin(func(e *gin.Engine) { initRoutes(e, svr) })
}

func (svr Server) Shutdown() {
}

func (svr Server) GetTemplate(store string, path string) string {
	storeOpt := svr.GetStoreOptions(store)
	templateStore := storage.GetTemplateStore(storeOpt)
	return templateStore.GetTemplate(path)
}

func (svr Server) PutTemplate(store string, path string, template string) string {

	storeOpt := svr.GetStoreOptions(store)
	templateStore := storage.GetTemplateStore(storeOpt)
	return templateStore.PutTemplate(path, template)
}

func (svr Server) DeleteTemplate(store string, path string) string {
	storeOpt := svr.GetStoreOptions(store)
	templateStore := storage.GetTemplateStore(storeOpt)
	return templateStore.DeleteTemplate(path)
}

func (svr Server) GetValues(store string, path string) string {
	storeOpt := svr.GetStoreOptions(store)
	valuesStore := storage.GetValuesStore(storeOpt)
	return valuesStore.GetValues(path)
}

func (svr Server) GetValuesInBatch(store string, paths []string) []string {
	storeOpt := svr.GetStoreOptions(store)
	valuesStore := storage.GetValuesStore(storeOpt)
	return valuesStore.GetValuesInBatch(paths)
}

func (svr Server) PutValues(store string, path string, values string) string {
	storeOpt := svr.GetStoreOptions(store)
	valuesStore := storage.GetValuesStore(storeOpt)
	return valuesStore.PutValues(path, values)
}

func (svr Server) DeleteValues(store string, path string) string {
	storeOpt := svr.GetStoreOptions(store)
	valuesStore := storage.GetValuesStore(storeOpt)
	return valuesStore.DeleteValues(path)
}

func (svr Server) GetStoreOptions(store string) config.StoreOptions {

	var opt config.StoreOptions
	configRepo := svr.configRepoFactory(svr.config)

	if store == "" {
		optJson := configRepo.GetStoreOptions("default")
		if optJson != nil {
			json.Unmarshal(optJson, &opt)
		} else {
			return svr.config.Storage
		}
	} else {
		optJson := configRepo.GetStoreOptions(store)
		json.Unmarshal(optJson, &opt)
	}
	return opt
}

func (svr Server) PutStoreOptions(store string, optionsJson string) string {
	configRepo := svr.configRepoFactory(svr.config)
	if store == "" {
		store = "default"
	}
	return configRepo.PutStoreOptions(store, optionsJson)
}

func (svr Server) DeleteStoreOptions(store string) string {
	configRepo := svr.configRepoFactory(svr.config)
	if store == "" {
		store = "default"
	}
	return configRepo.DeleteStoreOptions(store)
}
