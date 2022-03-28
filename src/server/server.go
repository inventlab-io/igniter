package server

import (
	"github.com/gin-gonic/gin"
	"github.com/igniter/http"
	"github.com/igniter/server/config"
	"github.com/igniter/storage"
	"github.com/igniter/storage/etcd"
)

type Server struct {
	config               config.ServerConfig
	templateStoreFactory storage.StoreFactory
}

func (svr Server) Run(cfg config.ServerConfig) {
	svr.config = cfg
	svr.templateStoreFactory = etcd.StoreFactory

	http.InitGin(func(e *gin.Engine) { initRoutes(e, svr) })
}

func (svr Server) Shutdown() {
}

func (svr Server) PutTemplate(path string, template string) string {
	store := svr.templateStoreFactory(svr.config)

	return store.PutTemplate(path, template)
}

func (svr Server) GetTemplate(path string) string {
	store := svr.templateStoreFactory(svr.config)
	r := store.GetTemplate(path)
	return r
}
