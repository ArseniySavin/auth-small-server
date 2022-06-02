package repository

import (
	"context"
	"database/sql"
	"time"
)

// AddClient -
func (repo BaseRepository) AddClient(ctx context.Context, tx *sql.Tx, client, secret string, claimId *int) error {
	query := `
insert into clients(client_id, client_secret, active, start_work_date, claims_idref) 
	values($1, $2, $3, $4, $5);
`
	row, err := tx.ExecContext(ctx, query, client, secret, true, time.Now(), claimId)
	row = row
	if err != nil {
		return err
	}

	return nil
}

// AddClaims -
func (repo BaseRepository) AddClaims(ctx context.Context, tx *sql.Tx, claims string) (int, error) {
	query := `
	insert into claims(claims) 
		values(json_in($1)) returning claims_id;
`
	id := 0

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return id, err
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, claims).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

// AddGrant -
func (repo BaseRepository) AddGrant(ctx context.Context, tx *sql.Tx, client, grant string) error {
	query := `
insert into map_clients_grants(client_idref, grant_idref) 
	values($1, $2);
`
	_, err := tx.ExecContext(ctx, query, client, grant)
	if err != nil {
		return err
	}

	return nil
}

// AddScope -
func (repo BaseRepository) AddScope(ctx context.Context, tx *sql.Tx, client, scope string) error {
	query := `
insert into map_clients_scopes(client_idref, scope_idref) 
	values($1, $2);
`
	_, err := tx.ExecContext(ctx, query, client, scope)
	if err != nil {
		return err
	}

	return nil
}

// ChangeClientState -
func (repo BaseRepository) ChangeClientState(ctx context.Context, tx *sql.Tx, client string, isActive bool) error {
	query := `
update clients set active = $2 where client_id = $1;
`
	_, err := tx.ExecContext(ctx, query, client, isActive)
	if err != nil {
		return err
	}

	return nil
}

// UpdateClaim -
func (repo BaseRepository) UpdateClaim(ctx context.Context, tx *sql.Tx, claimId int, claim string) error {
	query := `
update claims set claims = $2 where claims_id = $1;
`
	_, err := tx.ExecContext(ctx, query, claimId, claim)
	if err != nil {
		return err
	}

	return nil
}

// UpdateClientClaimId -
func (repo BaseRepository) UpdateClientClaimId(ctx context.Context, tx *sql.Tx, claimId int, client string) error {
	query := `
update clients set claims_idref = $2 where client_id = $1;
`
	_, err := tx.ExecContext(ctx, query, client, claimId)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSecret -
func (repo BaseRepository) UpdateSecret(ctx context.Context, tx *sql.Tx, client, secret string) error {
	query := `
update clients set client_secret = $2 where client_id = $1;
`
	_, err := tx.ExecContext(ctx, query, client, secret)
	if err != nil {
		return err
	}

	return nil
}
