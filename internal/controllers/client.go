package controllers

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/typos"
)

// MakeClientGrantScopeClaim -
func (da *Controllers) MakeClientGrantScopeClaim(ctx context.Context, client typos.ClientRequest) error {
	return da.client.CreateClient(ctx, client)
}

// MakeClientGrant -
func (da *Controllers) MakeClientGrant(ctx context.Context, client typos.ClientRequest) error {
	return da.client.CreateClientWithoutScopeAndClaims(ctx, client)
}

// UpdateClientGrantScopeClaim -
func (da *Controllers) UpdateClientGrantScopeClaim(ctx context.Context, client typos.ClientRequest) error {
	return da.client.UpdateClient(ctx, client)
}
