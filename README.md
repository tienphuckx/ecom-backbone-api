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
    
# Test
curl -H "Authorization: valid-token" curl http://localhost:8080/user/email/phuc@gmail.com
curl http://localhost:8080/user/email/phuc@gmail.com