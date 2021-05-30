### Setup without docker

1. Create `.env` file with following:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=docker_go_mux
PORT=3000
```

2. Create a database with name is `docker_go_mux` in postgres

3. Run cmd `go run main.go` to start server

### Setup with docker

1. Run cmd `docker-compose up`, server run at port `:8080`
