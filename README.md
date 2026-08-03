auth rest api in go

## Packages

```bash
Gin Web Framework = github.com/gin-gonic/gin
Reading Env Variable = github.com/joho/godotenv

Postgres Driver pgx
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool


Databse Migration
scoop bucket add main
scoop install main/migrate


create migration 

migrate create -ext sql -dir db/migrations -seq create_users_table

migration up

migrate -path db/migrations -database "url" up

```

live reload server 
go install github.com/air-verse/air@latest
air init
update .ari.toml file
air

---
password hashing
go get golang.org/x/crypto/bcrypt

---
JWT library
go get github.com/golang-jwt/jwt/v5