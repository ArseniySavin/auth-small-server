create table clients (
    client_id		varchar(30)  not null, -- code
    client_secret	varchar(150) not null, -- api_key
    active          boolean default false,
    start_work_date timestamptz default now(),
    end_work_date   timestamptz null,
    claims_idref	int null,
    comment			varchar(255),
    
    constraint PK_client_id primary key (client_id),
    constraint UQ_client_id_client_secret unique (client_id, client_secret),
    constraint FK_claims_idref_claims_id foreign key (claims_idref) references claims (claims_id)
);