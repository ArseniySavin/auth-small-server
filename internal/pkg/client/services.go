package client

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/repository"
	"github.com/ArseniySavin/auth-small-server/typos"
)

const (
	BasePathClient = "/auth/clients"
)

// IClientServices -
type IClientServices interface {
	CreateClient(context.Context, typos.ClientRequest) error
	CreateClientWithoutScopeAndClaims(context.Context, typos.ClientRequest) error
	ChangeClientState(context.Context, string, bool) error
	UpdateClient(context.Context, typos.ClientRequest) error
}

// ClientServices - Services
type ClientServices struct {
	repo repository.Repository
}

// NewClientServices -
func NewClientServices(repo repository.Repository) IClientServices {
	return &ClientServices{
		repo: repo,
	}
}
