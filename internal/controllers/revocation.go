package controllers

import "context"

// ChangeClientState -
func (da *Controllers) ChangeClientState(ctx context.Context, clientId string, isActive bool) error {
	return da.client.ChangeClientState(ctx, clientId, isActive)
}
