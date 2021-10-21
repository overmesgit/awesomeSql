### Generate grpc

```shell
protoc --go_out=grpc --go_opt=paths=source_relative --go-grpc_out=grpc --go-grpc_opt=paths=source_relative sign_up.proto
```

### Generate models

```shell
sqlboiler psql -c sqlboiler.toml
```

### Migrations
```shell
migrate create -ext sql -dir migrations -seq password_len
migrate -database ${POSTGRESQL_URL} -path migrations up
```