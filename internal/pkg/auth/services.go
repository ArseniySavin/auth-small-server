package auth

import (
	"github.com/ArseniySavin/auth-small-server/internal/pkg/repository"
)

const (
	BasePathAuth   = "/auth/clients"
	BasePathScopes = "/auth/scopes"
	BasePathGrants = "/auth/grants"
)

// IAuthServices
type IAuthServices interface {
	Authenticate(clientId, clientSecret string) (bool, error)
	Authorizate(clientId, wantScopes string) error
	AuthorizateGrants(clientId, endpoint string) error
}

// AuthServices -
type AuthServices struct {
	repo repository.Repository
}

// NewAuthServices -
func NewAuthServices(repo repository.Repository) IAuthServices {
	return &AuthServices{
		repo: repo,
	}
}
