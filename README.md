# Go Admin

## Prerequisites
_Make sure if [air](https://github.com/air-verse/air) live-reloading is installed_

**Air is yet another live-reloading command line utility for developing Go applications. Run air in your project root directory, leave it alone, and focus on your code.**

To install air you can execute the curl below
```shell
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

## How to start project with air

>> Note: to run with air right now we need run api without docker and database with docker

```shell
$ docker compose up
$ air
```

## How to start project without air

```shell
$ docker compose up
```

```shell
go run cmd/server/main.go
```

## Building with docker
<ul>
    <li>
        Uncomment a block of lines at docker-compose
    </li>
</ul>

```yaml
#  go-admin:
#    build:
#      context: .
#      dockerfile: ./Dockerfile
#    ports:
#      - 3333:8080
#    env_file:
#      - ./.env
#    networks:
#      - go-admin-network
#    depends_on:
#      mysql-go-admin:
#        condition: service_healthy
```

<ul>
    <li>
        Build the image
    </li>
</ul>

```shell
docker build -t go_admin .
``` 

<ul>
    <li>
        Run Project
    </li>
</ul>

```shell
docker run -p 8080:8080 go_admin
```

## Running tests
```shell
go test ./...
```

