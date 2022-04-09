package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRoutes(r *gin.Engine, svr Server) {

	r.GET("/options/store", func(ctx *gin.Context) { getOptions(ctx, svr) })
	r.PUT("/options/store", func(ctx *gin.Context) { putOptions(ctx, svr) })
	r.GET("/options/store/k/:store", func(ctx *gin.Context) { getOptions(ctx, svr) })
	r.PUT("/options/store/k/:store", func(ctx *gin.Context) { putOptions(ctx, svr) })

	r.PUT("/:datatype/k/*path", func(ctx *gin.Context) { putUserData(ctx, svr) })
	r.GET("/:datatype/k/*path", func(ctx *gin.Context) { getUserData(ctx, svr) })
	r.PUT("/:datatype/:store/k/*path", func(ctx *gin.Context) { putUserData(ctx, svr) })
	r.GET("/:datatype/:store/k/*path", func(ctx *gin.Context) { getUserData(ctx, svr) })
}

func getOptions(ctx *gin.Context, svr Server) {
	store := ctx.Param("store")
	result := svr.GetStoreOptions(store)
	ctx.JSON(http.StatusOK, result)
}

func putOptions(ctx *gin.Context, svr Server) {

	store := ctx.Param("store")
	options, err := ctx.GetRawData()
	if err != nil {
		fmt.Errorf("Malformed template option request")
	}
	result := svr.PutStoreOptions(store, string(options))
	ctx.String(http.StatusOK, result)
}

func putUserData(ctx *gin.Context, svr Server) {

	datatype := ctx.Param("datatype")

	if datatype != "template" && datatype != "values" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	path := ctx.Param("path")
	store := ctx.Param("store")
	rawData, err := ctx.GetRawData()
	if err != nil {
		fmt.Errorf("Error putting rawData %s", path)
	}

	if datatype == "template" {
		result := svr.PutTemplate(store, path, string(rawData))
		ctx.String(http.StatusOK, result)
	} else {
		result := svr.PutValues(store, path, string(rawData))
		ctx.String(http.StatusOK, result)
	}
}

func getUserData(ctx *gin.Context, svr Server) {

	datatype := ctx.Param("datatype")

	if datatype != "template" && datatype != "values" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	path := ctx.Param("path")
	store := ctx.Param("store")

	if datatype == "template" {
		result := svr.GetTemplate(store, path)
		ctx.String(http.StatusOK, result)
	} else {
		result := svr.GetValues(store, path)
		ctx.String(http.StatusOK, result)
	}
}
