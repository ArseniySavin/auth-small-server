create table scopes (
    scope_id  varchar(30) not null,
    comment   varchar(255) not null,
    
    constraint PK_scope_id primary key (scope_id)
);

insert into scopes ("scope_id", "comment") values
	('all', 'All permits for api');