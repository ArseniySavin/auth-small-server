create table claims (
    claims_id	int generated always as identity,
	claims		json not null,
        
    constraint PK_claims_id PRIMARY KEY (claims_id)
);