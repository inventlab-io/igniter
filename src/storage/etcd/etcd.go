package etcd

import (
	"context"
	"github.com/igniter/server/config"
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

func StoreFactory(cfg config.ServerConfig) storage.TemplateStore {
	store := initStore(cfg)
	return &store
}

func initStore(cfg config.ServerConfig) EtcdStore {

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.RequestTimeout)*time.Second)
	c, _ := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Etcd.Endpoints,
		DialTimeout: time.Duration(cfg.Etcd.ConnectionTimeout) * time.Second,
	})

	e := EtcdStore{}

	e.client = c
	e.context = ctx

	return e
}

func (e *EtcdStore) PutTemplate(path string, template string) string {
	_, err := e.client.KV.Put(e.context, path, template)
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

	return "OK"
}

func (e EtcdStore) GetTemplate(path string) string {
	r, _ := e.client.KV.Get(e.context, path)
	defer e.client.Close()

	return string(r.Kvs[0].Value)
}
