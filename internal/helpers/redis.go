package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = "../../.env"
	}
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		log.Fatal("REDIS_URL is not set in .env file")
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})
}
