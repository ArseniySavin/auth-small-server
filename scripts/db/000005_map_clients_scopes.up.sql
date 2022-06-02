create table map_clients_scopes (
    map_clients_scopes_id   int generated always as identity,
    client_idref	        varchar(30) not null,
    scope_idref		        varchar(30) not null,
        
    constraint PK_map_clients_scopes_id primary key (map_clients_scopes_id),
    constraint FK_scope_idref_scope_id foreign key (scope_idref) references scopes (scope_id),
    constraint FK_client_idref_client_id foreign key (client_idref) references clients (client_id)
);