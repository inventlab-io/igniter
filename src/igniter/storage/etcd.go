package storage

import (
	"context"
	"fmt"
	"github.com/igniter/config"
	"github.com/mitchellh/mapstructure"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
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
	optionsKey := parseOptionsKey(key)
	return e.getData(optionsKey)
}

func (e *EtcdStore) PutStoreOptions(key string, optionsJson string) string {
	optionsKey := parseOptionsKey(key)
	return e.putData(optionsKey, optionsJson)
}

func (e *EtcdStore) DeleteStoreOptions(key string) string {
	optionsKey := parseOptionsKey(key)
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

// GetValues implements ValuesStore
func (e EtcdStore) GetValues(key string) string {
	valuesKey := parseValuesKey(key)
	return string(e.getData(valuesKey))
}

// PutValues implements ValuesStore
func (e *EtcdStore) PutValues(key string, values string) string {
	valuesKey := parseValuesKey(key)
	return e.putData(valuesKey, values)
}

func (e *EtcdStore) getData(key string) []byte {
	r, _ := e.client.KV.Get(e.context, key)
	defer e.client.Close()

	if r != nil && len(r.Kvs) > 0 && r.Kvs[0].Value != nil && len(r.Kvs[0].Value) > 0 {
		return r.Kvs[0].Value
	} else {
		return nil
	}
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

func parseOptionsKey(key string) string {
	return fmt.Sprintf(":opt:%s", key)
}

func parseTemplateKey(key string) string {
	return fmt.Sprintf(":tpl:%s", key)
}

func parseValuesKey(key string) string {
	return fmt.Sprintf(":val:%s", key)
}
