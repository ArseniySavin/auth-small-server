package etcd

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/typos"
	"path"
	"sync"
	"time"

	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache"
	etcdClient "go.etcd.io/etcd/clientv3"
)

const Name = "etcd_cache"

// init -
func init() {
	cache.Register(Name, NewEtcd())
}

// Etcd -
type Etcd struct {
	client *etcdClient.Client
	rw     sync.Mutex
}

// NewEtcd -
func NewEtcd() cache.Cache {
	return &Etcd{
		rw: sync.Mutex{},
	}
}

// Open -
func (e *Etcd) Open(dsn string) error {
	e.rw.Lock()
	defer e.rw.Unlock()

	var err error

	e.client, err = etcdClient.New(etcdClient.Config{
		Endpoints:   []string{dsn},
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		return err
	}

	return nil
}

// Close -
func (e *Etcd) Close() error {
	e.rw.Lock()
	defer e.rw.Unlock()

	if e.client == nil {
		return nil
	}

	if err := e.client.Close(); err != nil {
		return err
	}

	e.client = nil

	return nil
}

// Put -
func (e *Etcd) Put(ctx context.Context, data map[string]interface{}, ttl int64) error {
	e.rw.Lock()
	defer e.rw.Unlock()

	if e.client == nil {
		return typos.ErrCacheNotOpen
	}

	leaseId := etcdClient.LeaseID(0)

	if ttl > 0 {
		resp, err := e.client.Grant(ctx, ttl)
		if err != nil {
			return err
		}

		leaseId = resp.ID
	}

	for k, v := range data {
		value, ok := v.(string)
		if !ok {
			return typos.ErrStringConversion
		}

		_, err := e.client.Put(ctx, k, value, etcdClient.WithLease(leaseId))
		if err != nil {
			return err
		}
	}

	return nil
}

// Get -
func (e *Etcd) Get(ctx context.Context, key string) (map[string]interface{}, error) {
	e.rw.Lock()
	defer e.rw.Unlock()

	if e.client == nil {
		return nil, typos.ErrCacheNotOpen
	}

	if len(key) > 0 {
		key = path.Clean(key)
	}

	response, err := e.client.Get(ctx, key, etcdClient.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})

	for _, kv := range response.Kvs {
		result[string(kv.Key)] = string(kv.Value)
	}

	return result, nil
}

// Delete -
func (e *Etcd) Delete(ctx context.Context, key string) error {
	e.rw.Lock()
	defer e.rw.Unlock()

	if e.client == nil {
		return typos.ErrCacheNotOpen
	}

	_, err := e.client.Delete(ctx, key, etcdClient.WithPrefix())
	if err != nil {
		return err
	}

	return nil
}
