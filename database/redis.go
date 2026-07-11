package database

import (
	"context"
	"fmt"
	"log"
	"os"

	// Use v9 for modern context support
	"github.com/redis/go-redis/v9"
)

// Rdb is exported so your main package / handlers can use it
var Rdb *redis.Client
var ctx = context.Background()

// InitRedis starts with a capital letter so it can be called from main
func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PW"),               
		DB:       0,                
	})

	// Ping Redis to verify connection
	_, err := Rdb.Ping(ctx).Result()
	fmt.Print(Rdb.Get(ctx, "Test"));
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully!")
}