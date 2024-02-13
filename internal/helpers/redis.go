package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	err := godotenv.Load()
	redisUrl := os.Getenv("REDIS_URL")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if redisUrl == "" {
		log.Fatal("REDIS_URL is not set in .env file")
	}
	
	return redis.NewClient(&redis.Options{
		Addr: redisUrl,
		Password: "",
		DB: 0,
	})
}
