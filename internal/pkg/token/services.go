package token

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/repository"
)

const (
	BasePathToken  = "/auth/tokens"
	BasePathClaims = "/auth/claims"
)

// ITokenServices -
type ITokenServices interface {
	GetClaimsFromCache(ctx context.Context, clientId string) (map[string]interface{}, error)
	GetClaimsFromRepository(ctx context.Context, clientId string) (map[string]interface{}, error)
	GetAccessToken(ctx context.Context, reference string) (string, error)
	SetAccessToken(ctx context.Context, jwt string, ttl int64) error
	DeleteAccessToken(ctx context.Context, reference string) error
}

// TokenServices -
type TokenServices struct {
	repo repository.Repository
}

// NewTokenServices -
func NewTokenServices(repo repository.Repository) ITokenServices {
	return &TokenServices{
		repo: repo,
	}
}
