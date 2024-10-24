package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/magiconair/properties/assert"
)

// Initialize a Redis client for testing
func setupTestRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Adjust this if necessary
		Password: "123456",         // Use your Redis password
		DB:       0,                // Use default DB
	})

	// Optionally, you can test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return rdb
}

// TestRedisSetGet tests basic Redis operations (set and get)
func TestRedisSetGet(t *testing.T) {
	rdb := setupTestRedis()
	ctx := context.Background()

	// Set a key-value pair in Redis
	err := rdb.Set(ctx, "test_key", "test_value", 0).Err()
	assert.Equal(t, err, nil) // Assert that no error occurred

	// Get the value for the key
	val, err := rdb.Get(ctx, "test_key").Result()
	assert.Equal(t, err, nil)          // Assert that no error occurred
	assert.Equal(t, val, "test_value") // Assert that the value matches
}

// TestRedisKeyExpiration tests that a key expires in Redis after a given time
func TestRedisKeyExpiration(t *testing.T) {
	rdb := setupTestRedis()
	ctx := context.Background()

	// Set a key-value pair with an expiration of 1 second
	err := rdb.Set(ctx, "expire_key", "expire_value", 1*time.Second).Err()
	assert.Equal(t, err, nil) // Assert that no error occurred

	// Get the value before expiration
	val, err := rdb.Get(ctx, "expire_key").Result()
	assert.Equal(t, err, nil)            // Assert that no error occurred
	assert.Equal(t, val, "expire_value") // Assert that the value matches

	// Wait for the key to expire
	time.Sleep(2 * time.Second)

	// Try to get the key after expiration
	val, err = rdb.Get(ctx, "expire_key").Result()
	assert.Equal(t, err != nil, true) // Assert that an error occurred (key not found)
	assert.Equal(t, err, redis.Nil)   // Assert that the error is redis.Nil
	assert.Equal(t, val, "")          // Assert that the value is empty
}
