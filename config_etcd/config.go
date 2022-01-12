package etcd

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

var etcdCli *clientv3.Client

func init() {
	var (
		ends      string
		endpoints []string
		timeout   time.Duration
		err       error
	)
	timeout = time.Second * 10
	ends = os.Getenv("ETCD_ENDPOINTS")
	if ends == "" {
		ends = "127.0.0.1:2379"
	}
	endpoints = strings.Split(ends, ",")
	cfg := clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          timeout,
		DialKeepAliveTime:    timeout,
		DialKeepAliveTimeout: timeout,
	}
	if etcdCli, err = clientv3.New(cfg); err != nil {
		log.Panic(err)
	}
}

func Put(ctx context.Context, key string, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return clientv3.NewKV(etcdCli).Put(ctx, key, value, opts...)
}

func Watch(ctx context.Context, key string, ops ...clientv3.OpOption) (clientv3.WatchChan, error) {
	watchCh := etcdCli.Watch(ctx, key, ops...)
	return watchCh, nil
}

func Del(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	return etcdCli.Delete(ctx, key, opts...)
}

func Close() error {
	return etcdCli.Close()
}

func Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return etcdCli.Get(ctx, key, opts...)
}

func NewSession(ttl int) (*concurrency.Session, error) {
	return concurrency.NewSession(etcdCli, concurrency.WithTTL(ttl))
}
