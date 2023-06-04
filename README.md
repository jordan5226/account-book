# account-book
An "Account Book" backend implementation in Golang + Gin framework.  
  
# Import
```
1. http/https server framework
go get -u github.com/gin-gonic/gin

2. Database framework
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

3. Set database connection string as environment variable
go get github.com/joho/godotenv

4. Database version control
go get -u github.com/golang-migrate/migrate/v4

5. Data Validation
go get github.com/go-playground/validator/v10

6. Test
go get github.com/stretchr/testify
```
  
# Deploy
- Method(1) Run container by commands  
1. Run Postgres docker container
```
docker pull postgres:15.3-alpine3.18
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15.3-alpine3.18
```
2. Create DB
```
docker exec -it postgres /bin/sh
createdb --username=root --owner=root acctbook
```
  or  
```
docker exec -it postgres createdb --username=root --owner=root acctbook
```
  
3. Create network for communication between different containers
```
docker network create acctbook-network
docker network connect acctbook-network postgres
```
4. Run backend app docker container
```
docker build -t jordan/acctbook .
docker run --name acctbook --network acctbook-network -p 8080:8080 -e GIN_MODE=release -e PG_URL=postgres://root:password@postgres:5432/acctbook?sslmode=disable jordan/acctbook
```
- Method(2) Alternatively, a docker yaml file can be used to compose containers:  
1. Create file [docker-compose.yaml](docker-compose.yaml)
```
services:
  postgres:
    image: postgres:15.3-alpine3.18
    volumes:
      - ./docker-entrypoint-initdb.d:/docker-entrypoint-initdb.d
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=password
      - POSTGRES_MULTIPLE_DATABASES=acctbook
    ports:
      - "5432:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready"]
      interval: 10s
      timeout: 5s
      retries: 5

  backend:
    image: jordan/acctbook
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
      - PG_URL=postgres://root:password@postgres:5432/acctbook?sslmode=disable
    depends_on:
      postgres:
        condition: service_healthy
```
2. Compose up
```
docker compose up
```
3. If you want to abandon the compose
```
docker compose down
docker images // use it to find useless image
docker rmi <IMAGE ID> // assign image ID to remove image
```
  
# Usage
Connect using RESTful API  

1. Create User  
  ```
  POST 127.0.0.1/acctbook/user
  JSON:
  {
      "name": "Yami",
      "uid": "yami0001",
      "pwd": "12345678",
      "currency": "TWD"
  }
  ```  
2. Get User Data
  ```
  GET 127.0.0.1/acctbook/user
  JSON:
  {
      "uid": "yami0001",
      "pwd": "12345678"
  }
  ```  
