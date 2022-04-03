package etcdstore

import (
	"context"
	"encoding/json"
	"github.com/igniter/config"
	"github.com/igniter/storage"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type EtcdStore struct {
	client  *clientv3.Client
	context context.Context
}

func ConfigRepoFactory(cfg config.ServerConfig) storage.ConfigRepo {
	store := initStore(cfg.Etcd)
	return &store
}

// GetTemplateStore implements ConfigRepo
func (e *EtcdStore) GetTemplateStore(key string) storage.TemplateStore {

	r, _ := e.client.KV.Get(e.context, key)
	settingsJson := r.Kvs[0].Value
	var opt config.EtcdOptions
	json.Unmarshal(settingsJson, &opt)
	store := initStore(opt)
	return &store
}

func (e *EtcdStore) GetTemplateStoreOptions(key string) string {
	result, _ := e.client.Get(e.context, key)
	return string(result.Kvs[0].Value)
}

func (e *EtcdStore) PutTemplateStoreOptions(key string, optionsJson string) string {
	_, err := e.client.Put(e.context, key, optionsJson)
	if err != nil {
	}
	return "Ok"
}

// PutTemplate implements TemplateStore
func (e *EtcdStore) PutTemplate(key string, template string) string {
	_, err := e.client.KV.Put(e.context, key, template)
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

// GetTemplate implements TemplateStore
func (e EtcdStore) GetTemplate(key string) string {
	r, _ := e.client.KV.Get(e.context, key)
	defer e.client.Close()

	return string(r.Kvs[0].Value)
}

//private functions
func initStore(cfg config.EtcdOptions) EtcdStore {

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.RequestTimeout)*time.Second)
	c, _ := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.ConnectionTimeout) * time.Second,
	})

	e := EtcdStore{}

	e.client = c
	e.context = ctx

	return e
}
