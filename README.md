# ecom-backbone-api
Ecommerce backbone API - using Go

# Libraries used

## 1. Gin
    - Gin is a web framework written in Go (Golang) that provides a set of features like routing, middleware, and binding for web applications.
    - install:
        go get -u github.com/gin-gonic/gin
## 1. Zap
    - Zap is a high-performance, structured logging library for Go, designed and developed by Uber. It is one of the fastest logging libraries available, offering both structured (key-value) and unstructured (printf-style) logging.
    - install:
        go get -u go.uber.org/zap
## 2. Viper
    - Viper is a popular configuration management library for Go, designed to handle dynamic configuration in Go applications. It simplifies working with different types of configuration formats and provides a consistent API to retrieve configuration values.
    - install:
        go get github.com/spf13/viper

## 3. Lumberjack
    - Lumberjack is a log rotation and compression package for Go, designed to be easy to use and integrate with other libraries and frameworks. It provides a simple and configurable way to rotate, compress, and store log files.
    - install:
        go get github.com/natefinch/lumberjack
        go get github.com/natefinch/lumberjack
    - Note: This package is used for logging.

## 4. JWT-Go
    - JWT-Go is a Go package for implementing JWT (JSON Web Tokens) as per RFC 7519. It provides a simple and efficient way to handle JSON Web Tokens in Go applications.
    - install:
    go get -u github.com/dgrijalva/jwt-go
    - Note: This package is used for JWT authentication.

## 5. MySQL Driver 
    go get -u github.com/go-sql-driver/mysql

## 6. Reids
    - Redis is an open-source, in-memory data structure store, used as a database, cache, and message broker.
    - install:
        go get github.com/go-redis/redis/v8

## 7. Kafka
    go get github.com/segmentio/kafka-go

## 8. Gorm
    - GORM is an open-source ORM (Object-Relational Mapping) library for Go, providing a simple and efficient way to interact with a relational database.
    - install:
        go get -u gorm.io/gorm
        go get -u gorm.io/driver/mysql

## 9. Google UUID
    - Google UUID is a package for generating and parsing UUIDs in Go. It provides a simple and efficient way to generate and validate UUIDs.
    - install:
        go get -u github.com/google/uuid

## 10. Benchmark
    - Benchmark is a testing tool in Go that provides a way to measure the performance of code. It is useful for testing the performance of functions, methods, and entire packages.
    - install:
        go get -u golang.org/x/perf/cmd/benchcmp
    
    

# Run Code Coverage
    - Run tests with coverage:
        go test -coverprofile=coverage.out
        go tool cover -html=coverage.out -o coverage.html


# TESTING

## API
    curl -H "Authorization: valid-token" curl http://localhost:8080/user/email/phuc@gmail.com
    curl http://localhost:8080/user/email/phuc@gmail.com

## MySQL
    mysql -h 127.0.0.1 -P 3306 -u root -p --protocol=tcp
    GRANT ALL PRIVILEGES ON *.* TO 'phuc'@'%' IDENTIFIED BY '123456';
    FLUSH PRIVILEGES;
    docker exec -it mysql-db bash
    mysql -uroot -p123456
    use boxfetch;
    show tables;
    desc go_db_user;

## Benchmark
    go test -bench=. -benchmem
    
## Step Update content of a docker image and restart (redis)
    docker-compose down --volumes
    docker system prune -f --volumes
    mkdir -p ./redis_data
    docker-compose up -d

    docker network ps
    docker network inspect docker network inspect nginx-mysql-network
## Port
    sudo lsof -i :8080
