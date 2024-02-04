package config

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Config struct {
	RedisHostPort string
	RedisPassword string
	RedisDB       int

	MaxNumber int64
}

func NewConfig() (*Config, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}

	redisHP := fmt.Sprintf("%s:%s", redisHost, redisPort)

	maxNumber := math.MaxInt64
	return &Config{
		RedisHostPort: redisHP,
		RedisPassword: redisPassword,
		RedisDB:       redisDB,

		MaxNumber: int64(maxNumber),
	}, nil
}
