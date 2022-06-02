package client

import (
	"context"
	"fmt"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/tools"
	"github.com/ArseniySavin/auth-small-server/typos"
	"strings"
)

// CreateClient -
func (bo *ClientServices) CreateClient(ctx context.Context, client typos.ClientRequest) error {
	grants, err := bo.repo.GetGrants(ctx)
	if err != nil {
		return err
	}
	scopes, err := bo.repo.GetScopes(ctx)
	if err != nil {
		return err
	}

	tx, err := bo.repo.StartTransaction(ctx)
	if err != nil {
		return err
	}

	scopeId, err := bo.repo.AddClaims(ctx, tx, client.Data.String())
	if err != nil {
		tx.Rollback()
		return err
	}

	err = bo.repo.AddClient(ctx, tx, client.ClientId, client.ClientSecret.String(), &scopeId)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range client.GrantTypes {
		err = bo.repo.AddGrant(ctx, tx, client.ClientId, grants[item].Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range client.Scope {
		itemStr := fmt.Sprint(item)
		err = bo.repo.AddScope(ctx, tx, client.ClientId, scopes[itemStr].Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return nil
}

// ChangeClientState -
func (bo *ClientServices) ChangeClientState(ctx context.Context, clientId string, isActive bool) error {
	tx, err := bo.repo.StartTransaction(ctx)
	if err != nil {
		return err
	}

	err = bo.repo.ChangeClientState(ctx, tx, clientId, isActive)
	if err != nil {
		tx.Commit()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return nil
}

// UpdateClient -
func (bo *ClientServices) UpdateClient(ctx context.Context, client typos.ClientRequest) error {
	grants, err := bo.repo.GetGrants(ctx)
	if err != nil {
		return err
	}
	scopes, err := bo.repo.GetScopes(ctx)
	if err != nil {
		return err
	}

	clientGrantsStr, err := bo.repo.GetClientGrants(ctx, client.ClientId)
	if err != nil {
		return err
	}
	clientGrants := strings.Split(clientGrantsStr, ",")

	clientScopeStr, err := bo.repo.GetClientScopes(ctx, client.ClientId)
	if err != nil {
		return err
	}

	clientScopes := strings.Split(clientScopeStr, ",")

	tx, err := bo.repo.StartTransaction(ctx)
	if err != nil {
		return err
	}

	clientDto, err := bo.repo.GetClient(ctx, client.ClientId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if clientDto.ClaimIdref == 0 {
		claimId, err := bo.repo.AddClaims(ctx, tx, client.Data.String())
		if err != nil {
			tx.Rollback()
			return err
		}

		err = bo.repo.UpdateClientClaimId(ctx, tx, claimId, client.ClientId)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		err = bo.repo.UpdateClaim(ctx, tx, clientDto.ClaimIdref, client.Data.String())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	grantItems := tools.GetDifference(client.GrantTypes, clientGrants)
	for _, item := range grantItems {
		err = bo.repo.AddGrant(ctx, tx, client.ClientId, grants[item].Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	scopeItems := tools.GetDifference(client.Scope, clientScopes)
	for _, item := range scopeItems {
		err = bo.repo.AddScope(ctx, tx, client.ClientId, scopes[item].Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if client.ClientSecret != "" {
		bo.repo.UpdateSecret(ctx, tx, client.ClientId, client.ClientSecret.String())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return nil
}

// CreateClientWithoutScopeAndClaims -
func (bo *ClientServices) CreateClientWithoutScopeAndClaims(ctx context.Context, client typos.ClientRequest) error {
	grants, err := bo.repo.GetGrants(ctx)
	if err != nil {
		return err
	}

	tx, err := bo.repo.StartTransaction(ctx)
	if err != nil {
		return err
	}

	err = bo.repo.AddClient(ctx, tx, client.ClientId, client.ClientSecret.String(), nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range client.GrantTypes {
		err = bo.repo.AddGrant(ctx, tx, client.ClientId, grants[item].Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return nil
}
