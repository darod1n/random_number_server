package app

import (
	"log"
	"net/http"
	"random_number_server/internal/config"
	redisrepository "random_number_server/internal/core/adapters/repositories/redis"
	wshandlers "random_number_server/internal/core/handlers/ws"
	service "random_number_server/internal/core/services"
)

func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	randomRepo, err := redisrepository.NewRepository(
		cfg.RedisHostPort,
		cfg.RedisPassword,
		cfg.RedisDB,
	)
	if err != nil {
		return err
	}

	randomGenerator := service.NewGeneratorRandomNumber(
		randomRepo,
		cfg.MaxNumber,
	)
	handlers := wshandlers.NewWSHandlers(randomGenerator)
	http.HandleFunc("/", handlers.Connection)
	log.Println("Listening on port 8080...")
	return http.ListenAndServe(":8080", nil)
}
