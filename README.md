### **Introduce**
The small auth server was based on Oauth. The server is using parts of the big RFC
### **Docs**
- API
  - [Token](/docs/get_token.md)
  - [Revoke](/docs/revoke_token.md)
  - [Introspection](/docs/introspect_token.md)
  - TODO [Client](/docs/client.md)
  - TODO [Revocation](/docs/revocation.md)


- [Install ETCD](https://etcd.io/docs/v3.5/install/) or using [image](https://hub.docker.com/r/bitnami/etcd) from docker hub
- [Tutorial ETCD](https://etcd.io/docs/v3.5/tutorials/)
- [PostgreSQL](https://www.postgresql.org/docs/14/index.html) or using [image](https://hub.docker.com/_/postgres) from docker hub
- Run local container 
```shell
docker run --name postgres -d -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres
docker run --name etcd -d --network bridge -p 2379:2379 -p 2380:2380 --env ALLOW_NONE_AUTHENTICATION=yes --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 bitnami/etcd
```

### **Config**
- For local running, use __.env__ file in the root directory
- For stage running, use docker-compose.yml in the deployment directory

#### **Data structure into cache**
```
|_ /auth 
|   |_ /clients
|   |   |_ client_id1=client_secret1
|   |   |_ client_id2=client_secret2
|   |   |_ ...
|   |
|   |_ /scopes
|   |   |_ client_id1=list_of_scopes1
|   |   |_ client_id2=list_of_scopes2
|   |   |_ ...
|   |
|   |_ /claims
|   |   |_ client_id1=json_of_claims1
|   |   |_ client_id2=json_of_claims2
|   |   |_ ...
|   |
|   |_ /tokens
|   |   |_ {reference1}={jwt1}
|   |   |_ {reference2}={jwt2}
|   |   |_ ...
|   |
|   |_ /grants
|   |   |_ client_id1=list_of_grants1
|   |   |_ client_id2=list_of_grants2
|   |   |_ ...
```

### **Deployment**
- Create data base __auth__. 
- Run __auth__. After starting the application, immediately executed migration
- Use docker-compose.{enviroment}.yml for TEST | PROD deployment. See deployment directory

### **Cache**
 
```shell
etcdctl --endpoints=http://localhost:2379 get /auth/ --prefix
etcdctl --endpoints=http://localhost:2379 del /auth/ --prefix
```