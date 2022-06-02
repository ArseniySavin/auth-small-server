package controllers

import "context"

// GetToken -
func (da *Controllers) GetToken(ctx context.Context, token string) (map[string]interface{}, error) {
	jwt, err := da.token.GetAccessToken(ctx, token)
	if err != nil {
		return nil, nil
	}

	claims, err := da.jwtMaker.Parse(jwt)
	if err != nil {
		return nil, nil
	}

	claims["active"] = true

	return claims, nil
}
