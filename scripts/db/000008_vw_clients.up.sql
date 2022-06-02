create view vw_clients as
select 
	cl.client_id,
	cl.client_secret,
	cl."comment",
	cm.claims,
	(
		select 
			string_agg(s."scope_id"::text, ',')
		from 
			scopes s
			join map_clients_scopes mcs
				on
					s.scope_id = mcs.scope_idref
			join clients cl1
				on 
					cl1.client_id = mcs.client_idref
					and cl1.client_id = cl.client_id 
	) as scopes,
	(
		select 
			string_agg(g.grant_id::text, ',')
		from 
			grants g
			join map_clients_grants mcg
				on
					g.grant_id = mcg.grant_idref
			join clients cl1
				on 
					cl1.client_id = mcg.client_idref
					and cl1.client_id = cl.client_id 
	) as grants
from
	clients cl
	left join claims cm
		on
			cl.claims_idref = cm.claims_id;