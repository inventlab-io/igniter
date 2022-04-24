package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRoutes(r *gin.Engine, svr Server) {

	r.GET("/options/store", func(ctx *gin.Context) { getStoreOptions(ctx, svr) })
	r.PUT("/options/store", func(ctx *gin.Context) { putStoreOptions(ctx, svr) })
	r.DELETE("/options/store", func(ctx *gin.Context) { deleteStoreOptions(ctx, svr) })
	r.GET("/options/secrets/k/:engine", func(ctx *gin.Context) { getSecretsOptions(ctx, svr) })
	r.PUT("/options/secrets/k/:engine", func(ctx *gin.Context) { putSecretsOptions(ctx, svr) })
	r.DELETE("/options/secrets/k/:engine", func(ctx *gin.Context) { deleteSecretsOptions(ctx, svr) })

	r.PUT("/secrets/map/k/*path", func(ctx *gin.Context) { putSecretsMapData(ctx, svr) })
	r.GET("/secrets/map/k/*path", func(ctx *gin.Context) { getSecretsMapData(ctx, svr) })
	r.DELETE("/secrets/map/k/*path", func(ctx *gin.Context) { deleteSecretsMapData(ctx, svr) })
	r.PUT("/secrets/maps/:store/k/*path", func(ctx *gin.Context) { putSecretsMapData(ctx, svr) })
	r.GET("/secrets/map/s/:store/k/*path", func(ctx *gin.Context) { getSecretsMapData(ctx, svr) })
	r.DELETE("/secrets/map/s/:store/k/*path", func(ctx *gin.Context) { deleteSecretsMapData(ctx, svr) })
	r.PUT("/secrets/map/:engine/k/*path", func(ctx *gin.Context) { putSecretsMapData(ctx, svr) })
	r.GET("/secrets/map/:engine/k/*path", func(ctx *gin.Context) { getSecretsMapData(ctx, svr) })
	r.DELETE("/secrets/map/:engine/k/*path", func(ctx *gin.Context) { deleteSecretsMapData(ctx, svr) })
	r.PUT("/secrets/map/:engine/s/:store/k/*path", func(ctx *gin.Context) { putSecretsMapData(ctx, svr) })
	r.GET("/secrets/map/:engine/s/:store/k/*path", func(ctx *gin.Context) { getSecretsMapData(ctx, svr) })
	r.DELETE("/secrets/map/:engine/s/:store/k/*path", func(ctx *gin.Context) { deleteSecretsMapData(ctx, svr) })

	r.PUT("/:datatype/k/*path", func(ctx *gin.Context) { putUserData(ctx, svr) })
	r.GET("/:datatype/k/*path", func(ctx *gin.Context) { getUserData(ctx, svr) })
	r.DELETE("/:datatype/k/*path", func(ctx *gin.Context) { deleteUserData(ctx, svr) })
	r.PUT("/:datatype/:store/k/*path", func(ctx *gin.Context) { putUserData(ctx, svr) })
	r.GET("/:datatype/:store/k/*path", func(ctx *gin.Context) { getUserData(ctx, svr) })
	r.DELETE("/:datatype/:store/k/*path", func(ctx *gin.Context) { deleteUserData(ctx, svr) })

	r.POST("/render/k/*path", func(ctx *gin.Context) { render(ctx, svr) })
	r.POST("/render/:store/k/*path", func(ctx *gin.Context) { render(ctx, svr) })
}

func getStoreOptions(ctx *gin.Context, svr Server) {
	store := ctx.Param("store")
	result := svr.GetStoreOptions(store)
	ctx.JSON(http.StatusOK, result)
}

func putStoreOptions(ctx *gin.Context, svr Server) {
	store := ctx.Param("store")
	options, err := ctx.GetRawData()
	if err != nil {
		fmt.Errorf("Malformed store option request")
	}
	result := svr.PutStoreOptions(store, string(options))
	ctx.String(http.StatusOK, result)
}

func deleteStoreOptions(ctx *gin.Context, svr Server) {
	store := ctx.Param("store")
	result := svr.DeleteStoreOptions(store)
	ctx.String(http.StatusOK, result)
}

func getSecretsOptions(ctx *gin.Context, svr Server) {
	engine := ctx.Param("engine")
	result := svr.GetSecretsOptions(engine)
	ctx.JSON(http.StatusOK, result)
}

func putSecretsOptions(ctx *gin.Context, svr Server) {
	engine := ctx.Param("engine")
	options, err := ctx.GetRawData()
	if err != nil {
		fmt.Errorf("Malformed secret engine option request")
	}
	result := svr.PutSecretsOptions(engine, string(options))
	ctx.String(http.StatusOK, result)
}

func deleteSecretsOptions(ctx *gin.Context, svr Server) {
	engine := ctx.Param("engine")
	result := svr.DeleteSecretsOptions(engine)
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
	} else if datatype == "values" {
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
	} else if datatype == "values" {
		result := svr.GetValues(store, path)
		ctx.String(http.StatusOK, result)
	}
}

func deleteUserData(ctx *gin.Context, svr Server) {

	datatype := ctx.Param("datatype")

	if datatype != "template" && datatype != "values" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	path := ctx.Param("path")
	store := ctx.Param("store")

	if datatype == "template" {
		result := svr.DeleteTemplate(store, path)
		ctx.String(http.StatusOK, result)
	} else if datatype == "values" {
		result := svr.DeleteValues(store, path)
		ctx.String(http.StatusOK, result)
	}
}

func putSecretsMapData(ctx *gin.Context, svr Server) {

	path := ctx.Param("path")
	engine := ctx.Param("engine")
	store := ctx.Param("store")
	rawData, err := ctx.GetRawData()
	if err != nil {
		fmt.Errorf("Error putting rawData %s", path)
	}

	result := svr.PutSecretsMap(engine, store, path, string(rawData))
	ctx.String(http.StatusOK, result)
}

func getSecretsMapData(ctx *gin.Context, svr Server) {

	path := ctx.Param("path")
	engine := ctx.Param("engine")
	store := ctx.Param("store")

	result := svr.GetSecretsMap(engine, store, path)
	ctx.String(http.StatusOK, result)
}

func deleteSecretsMapData(ctx *gin.Context, svr Server) {

	path := ctx.Param("path")
	engine := ctx.Param("engine")
	store := ctx.Param("store")
	result := svr.DeleteSecretsMap(engine, store, path)
	ctx.String(http.StatusOK, result)
}

func render(ctx *gin.Context, svr Server) {

	templatePath := ctx.Param("path")
	store := ctx.Param("store")
	rawBody, _ := ctx.GetRawData()
	render := parseRenderRequest(store, rawBody)
	result, ok := svr.Render(store, templatePath, render)
	if ok {
		ctx.String(http.StatusOK, result)
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

func parseRenderRequest(defaultStore string, rawBody []byte) RenderDto {
	var render RenderDto

	var renderRequest map[string]interface{}
	json.Unmarshal(rawBody, &renderRequest)

	renderRequestValues := renderRequest["values"]
	switch t := renderRequestValues.(type) {
	case string:
		render.Values = append(render.Values, makeNewRenderValue(defaultStore, t))
	case []interface{}:
		renderRequestValues := renderRequestValues.([]interface{})
		for _, rrv := range renderRequestValues {
			if strVal, isString := rrv.(string); isString {
				render.Values = append(render.Values, makeNewRenderValue(defaultStore, strVal))
			} else if mapVal, isMap := rrv.(map[string]interface{}); isMap {
				newRenderVal := convertMapToRenderValue(mapVal)
				render.Values = append(render.Values, newRenderVal)
			}

		}
	}
	return render
}

func convertMapToRenderValue(r map[string]interface{}) RenderValue {
	path := r["path"].(string)
	newRv := RenderValue{Path: path}
	storeKeys := r["storeKeys"]

	if strKeys, isString := storeKeys.(string); isString {
		newRv.StoreKeys = append(newRv.StoreKeys, strKeys)
	} else if array, isArray := storeKeys.([]interface{}); isArray {
		newRv = RenderValue{Path: path}
		for _, item := range array {
			newRv.StoreKeys = append(newRv.StoreKeys, item.(string))
		}
	}

	return newRv
}

func makeNewRenderValue(store string, path string) RenderValue {
	rv := RenderValue{}
	rv.Path = path
	rv.StoreKeys = append(rv.StoreKeys, store)
	return rv
}
