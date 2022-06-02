package repository

import (
	"context"
	"database/sql"
)

// GetClientSecret -
func (repo BaseRepository) GetClientSecret(ctx context.Context, code string) (string, error) {
	query := `
select 
	c.client_secret
from
	clients c 
where 
	c.client_id = $1
	and c.active = true 
	and now() <= coalesce(c.end_work_date, now());`

	row := repo.DB.QueryRowContext(ctx, query, code)

	var clientSecret string
	err := row.Scan(&clientSecret)
	if err != nil {
		return "", err
	}

	return clientSecret, nil
}

// GetClientScopes -
func (repo BaseRepository) GetClientScopes(ctx context.Context, code string) (string, error) {
	query := `
select 
	string_agg(s.scope_id::text, ',') as scopes
from
	scopes s
	join map_clients_scopes mcs 
		on
			s.scope_id = mcs.scope_idref
	join clients c
		on
			c.client_id = mcs.client_idref
where 
	c.client_id = $1`

	row := repo.DB.QueryRowContext(ctx, query, code)

	var clientScopes *string
	err := row.Scan(&clientScopes)
	if err != nil {
		return "", err
	}
	if clientScopes == nil {
		return "", nil
	}

	return *clientScopes, nil
}

// GetClient -
func (repo BaseRepository) GetClient(ctx context.Context, code string) (*Client, error) {
	query := `
select 
    c.client_secret,
    c.active,
    c.start_work_date,
    c.end_work_date,
	c.claims_idref
from
	clients c
where 
	c.client_id = $1`

	row := repo.DB.QueryRowContext(ctx, query, code)

	var client Client
	var claimIdref *int
	client.ClientId = code
	err := row.Scan(&client.Secret, &client.Active, &client.StartWDate, &client.EndWDate, &claimIdref)
	if err != nil {
		return nil, err
	}

	if claimIdref == nil {
		client.ClaimIdref = 0
	}

	return &client, nil
}

// GetClientClaims -
func (repo BaseRepository) GetClientClaims(ctx context.Context, code string) (string, error) {
	query := `
select 
	cl.claims_id
from
	claims cl
	join clients c
		on
			cl.claims_id = c.claims_idref 
where 
	c.client_id = $1`

	row := repo.DB.QueryRowContext(ctx, query, code)

	var claims string
	err := row.Scan(&claims)
	if err == sql.ErrNoRows {
		return "null", nil
	}
	if err != nil {
		return "", err
	}
	if &claims == nil {
		return "", nil
	}

	return claims, nil
}

// GetClientClaims -
func (repo BaseRepository) GetClientGrants(ctx context.Context, code string) (string, error) {
	query := `
select 
	string_agg(g.grant_id::text, ',') as grants
from
	grants g
	join map_clients_grants mcg 
		on
			g.grant_id = mcg.grant_idref 
	join clients c
		on
			c.client_id = mcg.client_idref
where 
	c.client_id = $1`

	row := repo.DB.QueryRowContext(ctx, query, code)

	var grants string
	err := row.Scan(&grants)
	if err != nil {
		return "", err
	}

	return grants, nil
}

// GetScopes -
func (repo BaseRepository) GetScopes(ctx context.Context) (map[string]*Scope, error) {
	query := `
select 
	s.scope_id
from
	scopes s`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	scopes := make(map[string]*Scope)
	for rows.Next() {
		var row Scope
		err = rows.Scan(&row.Id)

		if err != nil {
			break
		}

		scopes[row.Id] = &row
	}

	return scopes, nil
}

// GetGrants -
func (repo BaseRepository) GetGrants(ctx context.Context) (map[string]*Grant, error) {
	query := `
select 
	s.grant_id
from
	grants s`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	grants := make(map[string]*Grant)
	for rows.Next() {
		var row Grant
		err = rows.Scan(&row.Id)

		if err != nil {
			break
		}

		grants[row.Id] = &row
	}

	return grants, nil
}
