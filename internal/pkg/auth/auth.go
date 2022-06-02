package auth

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache"
	"github.com/ArseniySavin/auth-small-server/typos"
	"path"
	"strings"
)

// Authenticate -
func (a AuthServices) Authenticate(clientId, clientSecret string) (bool, error) {
	key := path.Join(BasePathAuth, clientId)

	data, err := cache.Get(key)
	if err != nil {
		return false, err
	}

	if val, ok := data[key]; ok {
		secret, ok := val.(string)
		if !ok {
			return false, typos.ErrConversion
		}

		return secret == clientSecret, nil
	}

	secret, err := a.repo.GetClientSecret(context.Background(), clientId)
	if err != nil {
		return false, err
	}

	cache.Put(map[string]interface{}{key: secret}, 3600)

	return secret == clientSecret, nil
}

// Authorizate -
func (a AuthServices) Authorizate(clientId, wantScopes string) error {
	scopes := ""
	key := path.Join(BasePathScopes, clientId)

	data, err := cache.Get(key)
	if err != nil {
		return err
	}

	if val, ok := data[key]; ok {
		scopes, ok = val.(string)
		if !ok {
			return typos.ErrConversion
		}
	} else {
		scopes, err := a.repo.GetClientScopes(context.Background(), clientId)
		if err != nil {
			return err
		}

		cache.Put(map[string]interface{}{key: scopes}, 3600)
	}

	for _, scope := range strings.Split(wantScopes, " ") {
		if !strings.Contains(scopes, scope) {
			return typos.ErrInvalidScope
		}
	}

	return nil
}

// AuthorizateGrants -
func (a AuthServices) AuthorizateGrants(clientId, endpoint string) error {
	grants := ""
	key := path.Join(BasePathGrants, clientId)

	data, err := cache.Get(key)
	if err != nil {
		return err
	}

	if val, ok := data[key]; ok {
		grants, ok = val.(string)
		if !ok {
			return typos.ErrConversion
		}
	} else {
		grants, err = a.repo.GetClientGrants(context.Background(), clientId)
		if err != nil {
			return err
		}

		cache.Put(map[string]interface{}{key: grants}, 3600)
	}

	if !strings.Contains(grants, endpoint) {
		return typos.ErrInvalidGrant
	}

	return nil
}
