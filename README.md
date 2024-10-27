# Go Admin

## Prerequisites
_Make sure if [migration](https://github.com/golang-migrate/migrate) is installed_

```shell
wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb
sudo dpkg -i migrate.linux-amd64.deb
sudo apt-get install make
```

_Make sure if [air](https://github.com/air-verse/air) live-reloading is installed_

**Air is yet another live-reloading command line utility for developing Go applications. Run air in your project root directory, leave it alone, and focus on your code.**

To install air you can execute the curl below
```shell
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

## How to run migrations
<ul>
    <li>
        Add the migration project to your setup   
    </li>
</ul>

_Make sure if make is installed_

```shell
make mg-up
```
>> Note: to run with air right now we need run api without docker and database with docker
<ul>
    <li>
        Running      
    </li>
</ul>

```shell
air
```

## How to start project with air

<p>After run docker execute</p>

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

