create table map_clients_grants (
    map_clients_grants_id   int generated always as identity,
    client_idref            varchar(30) not null,
    grant_idref             varchar(30) not null,
        
    constraint PK_map_clients_grants_id primary key (map_clients_grants_id),
    constraint FK_grant_idref_grant_id foreign key (grant_idref) references grants (grant_id),
    constraint FK_client_idref_client_id foreign key (client_idref) references clients (client_id)
);