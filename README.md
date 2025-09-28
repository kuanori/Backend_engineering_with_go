# Backend Engineering with Go

🚀A learning project built while taking "Backend Engineering with Go" Udemy course

**Course:** [Backend Engineering with Go](https://www.udemy.com/course/backend-engineering-with-go)

## ✨ Features

### 🏗 Architecture & Design
- **Clean Layered Architecture** (Transport/Service/Storage)
- **Repository Pattern** for data access
- **REST API**

### 🔐 Authentication & Authorization
- **JWT Token-based authentication**
- **Role-based authorization** with middleware
- **User registration & activation flow & email**
- **Secure password handling**

### 💾 Database & Persistence
- **PostgreSQL** with connection pooling
- **SQL migrations** with version control
- **Database seeding** for development
- **SQL transactions** for data consistency
- **Optimistic concurrency control**
- **Query timeouts** management

### 📊 Social Features
- **User management** (profiles, registration)
- **Post creation** with comments
- **User following system**
- **Feed algorithm** with pagination
- **Feed filtering** and sorting

### 🛠 Development & DevOps
- **Swagger/OpenAPI** auto-generated documentation
- **Structured logging** system
- **Continuous Integration** setup

### 🚀 Performance & Security
- **Rate limiting** middleware
- **CORS handling**
- **Caching strategies** (user profiles)
- **SQL indexes** optimization
- **Concurrency control** with mutexes
- **Server metrics** monitoring

### ✅ Testing & Quality
- **User handler** unit 
- **Error handling** standardization
- **Payload validation**
- **Performance testing** with Autocannon


### ⌨️ Commands
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
