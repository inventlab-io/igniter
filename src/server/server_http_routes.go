package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRoutes(r *gin.Engine, svr Server) {
	r.PUT("/template/:store/k/*path", func(ctx *gin.Context) { putTemplate(ctx, svr) })
	r.GET("/template/:store/k/*path", func(ctx *gin.Context) { getTemplate(ctx, svr) })
}

func putTemplate(ctx *gin.Context, svr Server) {

	templatePath := ctx.Param("path")
	store := ctx.Param("store")
	template, err := ctx.GetRawData()
	if err != nil {
		fmt.Errorf("Error putting template %s", templatePath)
	}

	result := svr.PutTemplate(store, templatePath, string(template))
	ctx.String(http.StatusOK, result)
}

func getTemplate(ctx *gin.Context, svr Server) {
	templatePath := ctx.Param("path")
	store := ctx.Param("store")
	result := svr.GetTemplate(store, templatePath)
	ctx.String(http.StatusOK, result)
}
