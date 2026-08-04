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

---

request validator
go get github.com/go-playground/validator/v10


## dployments

I have used EC2 t2.micro 1vCPU 1 GiB Memory
went i am building my go api binary on ec3, ram is full so it failed
so i decided to build it on github runner machine and copy that binary to my ec3 and run it
so it is working.

I don't even need to pull my github repo to ec2 machine and install go lang
