insert into clients (client_id, client_secret, active, claims_idref, "comment") values
('auth', 'Im5mSzlAZCFDYSpkNiI=', true, null, '');

-- grants
insert into map_clients_grants (client_idref, grant_idref) values
('auth', 'token'),
('auth', 'introspect'),
('auth', 'revoke'),
('auth', 'list-token'),
('auth', 'client'),
('auth', 'revocation');




