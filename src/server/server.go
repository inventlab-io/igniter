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
	templateStoreFactory map[string]storage.StoreFactory
}

func (svr Server) Run(cfg config.ServerConfig) {

	svr.config = cfg
	svr.templateStoreFactory = make(map[string]storage.StoreFactory)
	svr.templateStoreFactory["etcd"] = etcd.StoreFactory

	http.InitGin(func(e *gin.Engine) { initRoutes(e, svr) })
}

func (svr Server) Shutdown() {
}

func (svr Server) PutTemplate(store string, path string, template string) string {
	s := svr.templateStoreFactory[store](svr.config)
	return s.PutTemplate(path, template)
}

func (svr Server) GetTemplate(store string, path string) string {
	s := svr.templateStoreFactory[store](svr.config)
	r := s.GetTemplate(path)
	return r
}
