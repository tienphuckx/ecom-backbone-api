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
    go test ./test/redis/ -v
    go test ./test/basic/ -v
    go test ./test/<other>/ -v


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


## Test Router

Here are the updated `curl` commands with the port changed to `8999`:

### User Routes

#### 1. **Register a New User**


curl -X POST http://localhost:8999/v1/2024/user/register \
     -H "Content-Type: application/json" \
     -d '{"username": "newuser", "password": "password123", "email": "newuser@example.com"}'


#### 2. **Send OTP to User**


curl -X POST http://localhost:8999/v1/2024/user/send_otp \
     -H "Content-Type: application/json" \
     -d '{"phone": "+1234567890"}'


#### 3. **Get User Information**


curl -X GET http://localhost:8999/v1/2024/user/get_info \
     -H "Authorization: Bearer <your_token_here>"




### Admin Routes

#### 1. **Admin Login**


curl -X POST http://localhost:8999/v1/2024/admin/login \
     -H "Content-Type: application/json" \
     -d '{"username": "admin", "password": "admin123"}'


#### 2. **Activate User (Admin Action)**


curl -X POST http://localhost:8999/v1/2024/admin/active_user/123 \
     -H "Authorization: Bearer <admin_token_here>"


#### 3. **Add Shop (Admin Action)**


curl -X POST http://localhost:8999/v1/2024/admin/add_shop \
     -H "Authorization: Bearer <admin_token_here>" \
     -H "Content-Type: application/json" \
     -d '{"shop_name": "New Shop", "owner": "owner_id"}'




### Admin Managing Users

#### 1. **Update User Information (Admin Action)**


curl -X PUT http://localhost:8999/v1/2024/admin/user/update/123 \
     -H "Authorization: Bearer <admin_token_here>" \
     -H "Content-Type: application/json" \
     -d '{"email": "updatedemail@example.com", "status": "active"}'


#### 2. **Deactivate a User (Admin Action)**


curl -X POST http://localhost:8999/v1/2024/admin/user/deactivate/123 \
     -H "Authorization: Bearer <admin_token_here>"


#### 3. **Delete a User (Admin Action)**


curl -X DELETE http://localhost:8999/v1/2024/admin/user/delete/123 \
     -H "Authorization: Bearer <admin_token_here>"




### Product Routes (Optional)

#### 1. **Add a Product**


curl -X POST http://localhost:8999/v1/2024/product/add \
     -H "Authorization: Bearer <admin_token_here>" \
     -H "Content-Type: application/json" \
     -d '{"product_name": "New Product", "price": 100}'


#### 2. **Get Product List**


curl -X GET http://localhost:8999/v1/2024/product/list \
     -H "Authorization: Bearer <admin_token_here>"




### Check Server Status (For Monitoring)


curl -X GET http://localhost:8999/v1/2024/checkStatus




### Summary:
All `curl` commands are now updated to use port `8999`. You can run these commands to test your API with the new port configuration.

Let me know if you need further modifications or assistance!

