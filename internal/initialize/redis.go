package initialize

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"go.uber.org/zap"
)

var redisClient *redis.Client

// InitRedis initializes the Redis connection pool
func InitRedis() error {
	redisConfig := global.SysConfig.RedisConfig

	// Create a new Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
		PoolSize: redisConfig.PoolSize,
	})

	// Test the Redis connection
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	global.Logger.Info("Redis successfully connected", zap.String("host", redisConfig.Host))
	return nil
}
