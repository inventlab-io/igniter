package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/igniter/config"
	"github.com/igniter/http"
	"github.com/igniter/storage"
	"text/template"
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

func (svr Server) GetValuesInBatch(store string, paths []string) map[string]string {
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

func (svr Server) Render(store string, templatePath string, render RenderDto) (result string, ok bool) {

	templateValueMap := make(map[string]interface{})
	storeValueMap := prefetchValuesByBatch(render, svr)

	t := svr.GetTemplate(store, templatePath)

	values := render.Values
	for valueIndex := len(values) - 1; valueIndex >= 0; valueIndex-- {
		val := values[valueIndex]
		storeKeys := val.StoreKeys

		if len(storeKeys) == 0 {
			storeKeys = append(storeKeys, "")
		}

		for storeIndex := len(storeKeys) - 1; storeIndex >= 0; storeIndex-- {
			store := storeKeys[storeIndex]

			rawValue := storeValueMap[store][val.Path]
			var vm map[string]interface{}
			json.Unmarshal([]byte(rawValue), &vm)
			for k, v := range vm {
				templateValueMap[k] = v
			}
		}
	}

	tmpl, err := template.New(templatePath).Parse(t)

	if err != nil {
		return "", false
	} else {

		buf := new(bytes.Buffer)
		tmpl.Execute(buf, templateValueMap)
		return buf.String(), true
	}
}

func prefetchValuesByBatch(render RenderDto, svr Server) map[string]map[string]string {

	//storeMap[store] = valuePath
	storeMap := make(map[string][]string)
	for _, val := range render.Values {
		storeKeys := val.StoreKeys
		if len(storeKeys) == 0 {
			storeKeys = append(storeKeys, "")
		}
		for _, s := range storeKeys {
			storeMap[s] = append(storeMap[s], val.Path)
		}
	}

	//storeValueMap[store][valuePath] = Value
	storeValueMap := make(map[string]map[string]string)
	for store, valuePaths := range storeMap {
		valMap := svr.GetValuesInBatch(store, valuePaths)

		for valuePath, val := range valMap {
			if storeValueMap[store] == nil {
				storeValueMap[store] = make(map[string]string)
			}
			storeValueMap[store][valuePath] = val
		}
	}

	return storeValueMap
}
