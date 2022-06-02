create table grants (
    grant_id	varchar(30) not null,
    comment		varchar(255) not null,
    
    constraint PK_grant_id primary key (grant_id)
);

insert into grants (grant_id, comment) values
	('token', 'to use for get token'),
	('introspect', 'to use for get introspect'),
	('revoke', 'to use for get revoke token'),
	('list-token', 'get list tokens'),
    ('client', 'to use for create client'),
    ('revocation', 'to use for revoke client');