package initialize

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"go.uber.org/zap"
)

// ServerStart initializes all required services and starts the server
func ServerStart() {
	// Load configuration file and initialize global SysConfig struct
	LoadConfig()
	// Initialize logger first
	InitLogger()

	// MySQL database initialization
	if err := InitMySql(); err != nil {
		global.Logger.Fatal("Failed to initialize MySQL", zap.Error(err))
	}

	/*

		//Redis database initialization
		if err := InitRedis(); err != nil {
			global.Logger.Fatal("Failed to initialize Redis", zap.Error(err))
		}

		//Kafka producer initialization
		if err := InitKafka(); err != nil {
			global.Logger.Fatal("Failed to initialize Kafka", zap.Error(err))
		}

	*/

	// Logging server start
	global.Logger.Info("Server is starting...")

	// Router initialization
	r := InitRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Run the server in a separate goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatal("Failed to start API server", zap.Error(err))
		}
	}()

	// Listen for system interrupt signals (graceful shutdown)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	global.Logger.Info("Shutting down server...")

	// Close MySQL, Redis, and Kafka connections
	CloseMySql()
	CloseRedis()
	CloseKafka()

	// Graceful shutdown of the HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	global.Logger.Info("Server exited successfully")
}

func CloseMySql() {
	if db != nil {
		if err := db.Close(); err != nil {
			global.Logger.Error("Failed to close MySQL connection", zap.Error(err))
		} else {
			global.Logger.Info("MySQL connection closed successfully")
		}
	}
}

func CloseRedis() {
	if redisClient != nil {
		if err := redisClient.Close(); err != nil {
			global.Logger.Error("Failed to close Redis connection", zap.Error(err))
		} else {
			global.Logger.Info("Redis connection closed successfully")
		}
	}
}

func CloseKafka() {
	if kafkaWriter != nil {
		err := kafkaWriter.Close()
		if err != nil {
			global.Logger.Error("Failed to close Kafka producer", zap.Error(err))
		} else {
			global.Logger.Info("Kafka producer closed successfully")
		}
	}
}

// InitLogger initializes the global logger
// func InitLogger() {
// 	logConfig := global.SysConfig.LogConfig
// 	global.Logger = logger.NewLogger(logConfig)

// 	// Log that logger was successfully initialized
// 	global.Logger.Info("Logger initialized successfully")
// }

// InitMySql initializes MySQL database connection
// func InitMySql() error {
// 	mySQLConfig := global.SysConfig.MySQLConfig

// 	// Add logic to initialize MySQL here...
// 	fmt.Println("Initializing MySQL with the following config:")
// 	fmt.Println("DB Name:", mySQLConfig.DbName)
// 	fmt.Println("Host:", mySQLConfig.Host)
// 	fmt.Println("User:", mySQLConfig.User)
// 	fmt.Println("Port:", mySQLConfig.Port)

// 	// Add connection logic and return error if something goes wrong
// 	return nil
// }

// InitRedis initializes Redis connection
// func InitRedis() error {
// 	redisConfig := global.SysConfig.RedisConfig

// 	// Add logic to initialize Redis here...
// 	fmt.Println("Initializing Redis with the following config:")
// 	fmt.Println("Host:", redisConfig.Host)
// 	fmt.Println("Port:", redisConfig.Port)

// 	// Add connection logic and return error if something goes wrong
// 	return nil
// }

// InitKafka initializes Kafka producer
// func InitKafka() error {
// 	kafkaConfig := global.SysConfig.KafkaConfig

// 	// Add logic to initialize Kafka here...
// 	fmt.Println("Initializing Kafka with the following config:")
// 	fmt.Println("Brokers:", kafkaConfig.Brokers)

// 	// Add connection logic and return error if something goes wrong
// 	return nil
// }
