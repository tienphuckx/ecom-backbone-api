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

	//Redis database initialization
	if err := InitRedis(); err != nil {
		global.Logger.Fatal("Failed to initialize Redis", zap.Error(err))
	}
	/*

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
		Addr:    ":8999",
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
