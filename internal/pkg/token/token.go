package token

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache"
	"github.com/ArseniySavin/auth-small-server/typos"
	jsoniter "github.com/json-iterator/go"
	"path"
)

// GetClaimsFromCache -
func (bo *TokenServices) GetClaimsFromCache(ctx context.Context, clientId string) (map[string]interface{}, error) {
	key := path.Join(BasePathClaims, clientId)

	cacheData, err := cache.GetCtx(ctx, key)
	if err != nil {
		return nil, err
	}

	if len(cacheData) == 0 {
		return nil, typos.ErrClaimsNotFound
	}

	claims := cacheData[key].(string)

	kvs := make(map[string]interface{})
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal([]byte(claims), &kvs)

	return kvs, nil
}

// GetClaimsFromRepository -
func (bo *TokenServices) GetClaimsFromRepository(ctx context.Context, clientId string) (map[string]interface{}, error) {
	key := path.Join(BasePathClaims, clientId)

	claims, err := bo.repo.GetClientClaims(ctx, clientId)
	if err != nil {
		return nil, err
	}

	cache.Put(map[string]interface{}{key: claims}, 3600)

	kvs := make(map[string]interface{})
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal([]byte(claims), &kvs)

	return kvs, nil
}

// GetAccessToken -
func (bo *TokenServices) GetAccessToken(ctx context.Context, reference string) (string, error) {
	key := path.Join(BasePathToken, reference)

	data, err := cache.GetCtx(ctx, key)
	if err != nil {
		return "", err
	}

	jwt, ok := data[key]
	if !ok {
		return "", typos.ErrAccessTokenNotFound
	}

	return jwt.(string), nil
}

// SetAccessToken -
func (bo *TokenServices) SetAccessToken(ctx context.Context, jwt string, ttl int64) error {
	hash := sha256.New()
	hash.Write([]byte(jwt))

	reference := hex.EncodeToString(hash.Sum(nil))

	key := path.Join(BasePathToken, reference)
	data := map[string]interface{}{key: jwt}

	return cache.PutCtx(ctx, data, ttl)
}

// DeleteAccessToken -
func (bo *TokenServices) DeleteAccessToken(ctx context.Context, reference string) error {
	key := path.Join(BasePathToken, reference)

	return cache.DeleteCtx(ctx, key)
}
