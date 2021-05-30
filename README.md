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

3. Run cmd `go run main.go` and visit `http://localhost:3000/api/posts`

### Setup with docker
1. Change `docker-compose.yml` with following:
```
dockerfile: ./docker/app/Dockerfile.prod
-> dockerfile: ./docker/app/Dockerfile.dev

command: sh /scripts/prod.sh
-> command: sh /scripts/dev.sh
```
2. Run cmd `docker-compose up` and visit `http://localhost:80/api/posts`
