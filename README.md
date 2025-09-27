# Backend_engineering_with_go

### cmd - entrypoint of app
### bin - binaries


### Layers
#### Transport
#### Service
#### Storage


```bash
go run ./cmd/api
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

```bash
npx autocannon http://localhost:8080/v1/users/2 --connections 10 --duration 5 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJzb2NpYWwiLCJleHAiOjE3NTkyMTc4NDgsImlhdCI6MTc1ODk1ODY0OCwiaXNzIjoic29jaWFsIiwibmJmIjoxNzU4OTU4NjQ4LCJzdWIiOjEzN30.x3wIW1uByni6qpvqPF4-P3o1dinICeubMIB_-pK65e8"
```

```bash
npx autocannon -r 1000 -d 2 -c 10 --renderStatusCodes http://localhost:8080/v1/health
```
```bash
npx autocannon -c 1 -r 100 -d 1 --renderStatusCodes http://localhost:8080/v1/health
```