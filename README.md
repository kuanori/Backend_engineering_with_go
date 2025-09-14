# Backend_engineering_with_go

### cmd - entrypoint of app
### bin - binaries


### Layers
#### Transport
#### Service
#### Storage


```bash
go run cmd/api/*.go
```

### migrate lib
```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

```bash
make migrate-up
```

```bash
swag init
```