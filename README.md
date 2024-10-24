# ecom-backbone-api
Ecommerce backbone API - using Go

# Libraries used

## 1. Gin
    - Gin is a web framework written in Go (Golang) that provides a set of features like routing, middleware, and binding for web applications.
    - install:
        go get -u github.com/gin-gonic/gin
## 1. ZAP
    - Zap is a high-performance, structured logging library for Go, designed and developed by Uber. It is one of the fastest logging libraries available, offering both structured (key-value) and unstructured (printf-style) logging.
    - install:
        go get -u go.uber.org/zap
## 2. VIPER
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
    go get github.com/go-redis/redis/v8

## 7. Kafka
    go get github.com/segmentio/kafka-go


    

# Run Code Coverage
    - Run tests with coverage:
        go test -coverprofile=coverage.out
        go tool cover -html=coverage.out -o coverage.html


# Test
curl -H "Authorization: valid-token" curl http://localhost:8080/user/email/phuc@gmail.com
curl http://localhost:8080/user/email/phuc@gmail.com
mysql -h 127.0.0.1 -P 3306 -u root -p --protocol=tcp
GRANT ALL PRIVILEGES ON *.* TO 'phuc'@'%' IDENTIFIED BY '123456';
FLUSH PRIVILEGES;


