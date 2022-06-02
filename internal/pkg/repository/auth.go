package repository

import (
	"context"
	"database/sql"
)

// Repository -
type Repository interface {
	StartTransaction(context.Context) (*sql.Tx, error)
	GetClient(context.Context, string) (*Client, error)
	GetClientSecret(context.Context, string) (string, error)
	GetClientScopes(context.Context, string) (string, error)
	GetClientClaims(context.Context, string) (string, error)
	GetClientGrants(context.Context, string) (string, error)
	GetGrants(context.Context) (map[string]*Grant, error)
	GetScopes(context.Context) (map[string]*Scope, error)
	AddClient(context.Context, *sql.Tx, string, string, *int) error
	AddClaims(context.Context, *sql.Tx, string) (int, error)
	AddGrant(context.Context, *sql.Tx, string, string) error
	AddScope(context.Context, *sql.Tx, string, string) error
	ChangeClientState(context.Context, *sql.Tx, string, bool) error
	UpdateClaim(context.Context, *sql.Tx, int, string) error
	UpdateClientClaimId(context.Context, *sql.Tx, int, string) error
	UpdateSecret(context.Context, *sql.Tx, string, string) error
}

// BaseRepository -
type BaseRepository struct {
	DB *sql.DB
}

// NewRepository -
func NewRepository(db *sql.DB) Repository {
	return &BaseRepository{db}
}

func (repo BaseRepository) StartTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
