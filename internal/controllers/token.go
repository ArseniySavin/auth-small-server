package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/ArseniySavin/auth-small-server/typos"
	"time"
)

// MakeToken -
func (da *Controllers) MakeToken(ctx context.Context, clientId, scopes string) (*typos.MakeTokenResponse, error) {
	claims, err := da.token.GetClaimsFromCache(ctx, clientId)
	if err != nil {
		if !errors.Is(err, typos.ErrClaimsNotFound) {
			return nil, err
		}

		claims, err = da.token.GetClaimsFromRepository(ctx, clientId)
		if err != nil {
			return nil, err
		}
	}

	accessToken, err := da.jwtMaker.Make(claims)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()
	hash.Write([]byte(accessToken))
	referenceToken := hex.EncodeToString(hash.Sum(nil))

	err = da.token.SetAccessToken(ctx, accessToken, int64(3600))
	if err != nil {
		return nil, err
	}

	return &typos.MakeTokenResponse{
		Token:     referenceToken,
		TokenType: "Bearer",
		Scope:     scopes,
		Expire:    time.Now().Add(time.Hour).Unix(),
	}, nil
}
