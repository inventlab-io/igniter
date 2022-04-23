package storage

import (
	"context"
	"fmt"
	"github.com/igniter/config"
	"github.com/mitchellh/mapstructure"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strings"
	"time"
)

func EtcdConfigRepoFactory(cfg config.ServerConfig) ConfigRepo {
	var opt config.EtcdOptions
	mapstructure.Decode(cfg.Storage.Options, &opt)
	store := etcdInitStore(opt)
	return store
}

func etcdInitStore(cfg config.EtcdOptions) *EtcdStore {

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.RequestTimeout)*time.Second)
	c, _ := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.ConnectionTimeout) * time.Second,
	})

	e := EtcdStore{}

	e.client = c
	e.context = ctx

	return &e
}

type EtcdStore struct {
	client  *clientv3.Client
	context context.Context
}

func (e *EtcdStore) GetStoreOptions(key string) []byte {
	optionsKey := parseStoreOptionsKey(key)
	return e.getData(optionsKey)
}

func (e *EtcdStore) PutStoreOptions(key string, optionsJson string) string {
	optionsKey := parseStoreOptionsKey(key)
	return e.putData(optionsKey, optionsJson)
}

func (e *EtcdStore) DeleteStoreOptions(key string) string {
	optionsKey := parseStoreOptionsKey(key)
	return e.deleteData(optionsKey)
}

func (e *EtcdStore) GetSecretsOptions(key string) []byte {
	optionsKey := parseSecretsOptionsKey(key)
	return e.getData(optionsKey)
}

func (e *EtcdStore) PutSecretsOptions(key string, optionsJson string) string {
	optionsKey := parseSecretsOptionsKey(key)
	return e.putData(optionsKey, optionsJson)
}

func (e *EtcdStore) DeleteSecretsOptions(key string) string {
	optionsKey := parseSecretsOptionsKey(key)
	return e.deleteData(optionsKey)
}

// GetTemplate implements TemplateStore
func (e EtcdStore) GetTemplate(key string) string {
	templateKey := parseTemplateKey(key)
	return string(e.getData(templateKey))
}

// PutTemplate implements TemplateStore
func (e *EtcdStore) PutTemplate(key string, template string) string {
	templateKey := parseTemplateKey(key)
	return e.putData(templateKey, template)
}

// DeleteTemplate implements TemplateStore
func (e EtcdStore) DeleteTemplate(key string) string {
	templateKey := parseTemplateKey(key)
	return string(e.deleteData(templateKey))
}

// GetValues implements ValuesStore
func (e EtcdStore) GetValues(key string) string {
	valuesKey := parseValuesKey(key)
	return string(e.getData(valuesKey))
}

// GetValues implements ValuesStore
func (e EtcdStore) GetValuesInBatch(keys []string) map[string]string {

	var valuesKeys []string

	for _, k := range keys {
		valuesKey := parseValuesKey(k)
		valuesKeys = append(valuesKeys, valuesKey)
	}
	resultBytes := e.getDataInBatch(valuesKeys)

	values := make(map[string]string)
	for k, v := range resultBytes {
		userKey := stripInternalPrefix(k)
		values[userKey] = string(v)
	}
	return values
}

// PutValues implements ValuesStore
func (e *EtcdStore) PutValues(key string, values string) string {
	valuesKey := parseValuesKey(key)
	return e.putData(valuesKey, values)
}

// DeleteValues implements ValuesStore
func (e *EtcdStore) DeleteValues(key string) string {
	valuesKey := parseValuesKey(key)
	return e.deleteData(valuesKey)
}

func (e *EtcdStore) getData(key string) []byte {
	valuesMap := e.getDataInBatch([]string{key})
	if len(valuesMap) > 0 {
		return valuesMap[key]
	} else {
		return nil
	}
}

func (e *EtcdStore) getDataInBatch(keys []string) map[string][]byte {

	valueMap := make(map[string][]byte)

	defer e.client.Close()
	for _, k := range keys {
		r, _ := e.client.KV.Get(e.context, k)
		if r != nil && len(r.Kvs) > 0 && r.Kvs[0].Value != nil && len(r.Kvs[0].Value) > 0 {
			valueMap[k] = r.Kvs[0].Value
		}
	}

	return valueMap
}

func (e *EtcdStore) putData(key string, data string) string {

	_, err := e.client.KV.Put(e.context, key, data)
	defer e.client.Close()

	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	return "Ok"
}
func (e *EtcdStore) deleteData(key string) string {
	_, err := e.client.KV.Delete(e.context, key)
	defer e.client.Close()

	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	return "Ok"
}

func parseStoreOptionsKey(key string) string {
	return fmt.Sprintf(":opt_store:%s", key)
}

func parseSecretsOptionsKey(key string) string {
	return fmt.Sprintf(":opt_secrets:%s", key)
}

func parseTemplateKey(key string) string {
	return fmt.Sprintf(":tpl:%s", key)
}

func parseValuesKey(key string) string {
	return fmt.Sprintf(":val:%s", key)
}

func stripInternalPrefix(key string) string {
	if strings.HasPrefix(key, ":opt_store:") {
		return key[11:]
	} else if strings.HasPrefix(key, ":opt_secrets:") {
		return key[13:]
	} else if strings.HasPrefix(key, ":tpl:") ||
		strings.HasPrefix(key, ":val:") {
		return key[5:]
	} else {
		return key
	}
}
