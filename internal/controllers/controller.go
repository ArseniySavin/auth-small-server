package controllers

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/client"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/jwt"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/token"
	"github.com/ArseniySavin/auth-small-server/typos"
)

// IControllers -
type IControllers interface {
	MakeToken(context.Context, string, string) (*typos.MakeTokenResponse, error)
	DeleteToken(context.Context, string) error
	GetToken(context.Context, string) (map[string]interface{}, error)
	MakeClientGrantScopeClaim(context.Context, typos.ClientRequest) error
	MakeClientGrant(context.Context, typos.ClientRequest) error
	ChangeClientState(context.Context, string, bool) error
	UpdateClientGrantScopeClaim(context.Context, typos.ClientRequest) error
}

// Controllers -
type Controllers struct {
	token    token.ITokenServices
	client   client.IClientServices
	jwtMaker jwt.IJwtServices
}

// NewControllers -
func NewControllers(token token.ITokenServices, client client.IClientServices, jwtMaker jwt.IJwtServices) IControllers {
	return &Controllers{
		token:    token,
		client:   client,
		jwtMaker: jwtMaker,
	}
}
