package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"go.uber.org/zap"
)

// Redis client instance (global)
var RedisClient *redis.Client

// InitRedis initializes the Redis client with the configuration from the settings
func InitRedis() error {
	redisConfig := global.SysConfig.RedisConfig

	// Create Redis client with larger timeout
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		PoolSize:     redisConfig.PoolSize,
		MinIdleConns: redisConfig.MinIdleConns,
		DialTimeout:  time.Duration(redisConfig.DialTimeout) * time.Second,  // 5 seconds
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout) * time.Second,  // 3 seconds
		WriteTimeout: time.Duration(redisConfig.WriteTimeout) * time.Second, // 3 seconds
	})

	// Test the connection with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	global.Logger.Info("Redis connected successfully", zap.String("pong", pong))
	return nil
}

// CloseRedis gracefully closes the Redis connection
func CloseRedis() error {
	err := RedisClient.Close()
	if err != nil {
		global.Logger.Error("Failed to close Redis connection", zap.Error(err))
		return err
	}
	global.Logger.Info("Redis connection closed gracefully")
	return nil
}
