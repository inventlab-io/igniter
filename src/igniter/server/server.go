package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/igniter/config"
	"github.com/igniter/http"
	"github.com/igniter/storage"
	etcdstore "github.com/igniter/storage/etcd"
	"strings"
)

type Server struct {
	config            config.ServerConfig
	configRepoFactory storage.ConfigRepoFactory
}

func (svr Server) Run(cfg config.ServerConfig) {

	svr.config = cfg

	if cfg.Storage == "etcd" {
		svr.configRepoFactory = etcdstore.ConfigRepoFactory
	}

	http.InitGin(func(e *gin.Engine) { initRoutes(e, svr) })
}

func (svr Server) Shutdown() {
}

func (svr Server) PutTemplate(store string, path string, template string) string {
	configRepo := svr.configRepoFactory(svr.config)
	templateStore := configRepo.GetTemplateStore(optionsKey(store, path))
	return templateStore.PutTemplate(path, template)
}

func (svr Server) GetTemplate(store string, path string) string {
	configRepo := svr.configRepoFactory(svr.config)
	templateStore := configRepo.GetTemplateStore(optionsKey(store, path))
	return templateStore.GetTemplate(path)
}

func (svr Server) PutTemplateStoreOptions(store string, path string, optionsJson string) string {
	configRepo := svr.configRepoFactory(svr.config)
	return configRepo.PutTemplateStoreOptions(optionsKey(store, path), optionsJson)
}

func (svr Server) GetTemplateStoreOptions(store string, path string) string {
	configRepo := svr.configRepoFactory(svr.config)

	return configRepo.GetTemplateStoreOptions(optionsKey(store, path))
}

func optionsKey(store string, path string) string {

	if store == "" {
		store = "default"
	}

	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	key := fmt.Sprintf("/options/%s/template%s", store, path)
	return key
}
