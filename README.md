### Generate grpc

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative login_grpc/sign_up.proto
```

### Generate models

```shell
sqlboiler psql -c sqlboiler.toml
```

### Migrations
```shell
migrate create -ext sql -dir migrations -seq password_len
export POSTGRESQL_URL='postgres://data:{pass}@localhost:5432/gogo?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path login_psql/migrations up
```