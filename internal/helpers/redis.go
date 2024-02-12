package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		Password: "",
		DB: 0,
	})
}
