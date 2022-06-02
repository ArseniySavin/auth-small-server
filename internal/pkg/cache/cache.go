package cache

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/typos"

	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache/io"
)

// Cache -
type Cache interface {
	io.OpenCloser
	Put(ctx context.Context, data map[string]interface{}, ttl int64) error
	Get(ctx context.Context, key string) (map[string]interface{}, error)
	Delete(ctx context.Context, key string) error
}

var (
	cacheMap     map[string]Cache
	defaultCache Cache
)

// Register -
func Register(name string, cache Cache) {
	if cacheMap == nil {
		cacheMap = make(map[string]Cache)
	}

	cacheMap[name] = cache
}

// Open -
func Open(name, dsn string) error {
	if defaultCache != nil {
		return nil
	}

	cache, ok := cacheMap[name]
	if !ok {
		return typos.ErrUnknownCache
	}

	defaultCache = cache

	return defaultCache.Open(dsn)
}

// Close -
func Close() error {
	if defaultCache != nil {
		err := defaultCache.Close()
		defaultCache = nil

		return err
	}

	return nil
}

// Put -
func Put(data map[string]interface{}, ttl int64) error {
	if defaultCache == nil {
		return typos.ErrDefaultCacheUndefined
	}

	return PutCtx(context.Background(), data, ttl)
}

// PutCtx -
func PutCtx(ctx context.Context, data map[string]interface{}, ttl int64) error {
	if defaultCache == nil {
		return typos.ErrDefaultCacheUndefined
	}

	return defaultCache.Put(ctx, data, ttl)
}

// Get -
func Get(key string) (map[string]interface{}, error) {
	if defaultCache == nil {
		return nil, typos.ErrDefaultCacheUndefined
	}

	return GetCtx(context.Background(), key)
}

// GetCtx -
func GetCtx(ctx context.Context, key string) (map[string]interface{}, error) {
	if defaultCache == nil {
		return nil, typos.ErrDefaultCacheUndefined
	}

	return defaultCache.Get(ctx, key)
}

// Delete -
func Delete(key string) error {
	if defaultCache == nil {
		return typos.ErrDefaultCacheUndefined
	}

	return DeleteCtx(context.Background(), key)
}

// DeleteCtx -
func DeleteCtx(ctx context.Context, key string) error {
	if defaultCache == nil {
		return typos.ErrDefaultCacheUndefined
	}

	return defaultCache.Delete(ctx, key)
}
