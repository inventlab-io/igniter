package server

import (
	"fmt"
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

	if cfg.Storage == "etcd" {
		svr.configRepoFactory = storage.EtcdConfigRepoFactory
	}

	http.InitGin(func(e *gin.Engine) { initRoutes(e, svr) })
}

func (svr Server) Shutdown() {
}

func (svr Server) PutTemplate(store string, path string, template string) string {
	tplStoreOpt := svr.GetStoreOptions(store)
	templateStore := storage.GetTemplateStore(tplStoreOpt)
	return templateStore.PutTemplate(path, template)
}

func (svr Server) GetTemplate(store string, path string) string {
	tplStoreOpt := svr.GetStoreOptions(store)
	templateStore := storage.GetTemplateStore(tplStoreOpt)
	return templateStore.GetTemplate(path)
}

func (svr Server) PutStoreOptions(store string, optionsJson string) string {
	configRepo := svr.configRepoFactory(svr.config)
	return configRepo.PutStoreOptions(optionsKey(store), optionsJson)
}

func (svr Server) GetStoreOptions(store string) []byte {
	configRepo := svr.configRepoFactory(svr.config)
	return configRepo.GetStoreOptions(optionsKey(store))
}

func optionsKey(store string) string {

	if store == "" {
		store = "default"
	}

	key := fmt.Sprintf("/options/store/%s", store)
	return key
}
