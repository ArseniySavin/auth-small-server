package controllers

import "context"

// DeleteToken -
func (da *Controllers) DeleteToken(ctx context.Context, token string) error {
	return da.token.DeleteAccessToken(ctx, token)
}
