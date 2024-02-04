package redisrepository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func (r *Repository) Exist(ctx context.Context, key string) (bool, error) {
	_, err := r.client.Get(ctx, key).Result()
	if err == nil {
		return true, nil
	}
	if err != redis.Nil {
		return false, err
	}
	return false, nil
}

func (r *Repository) Set(ctx context.Context, key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func NewRepository(redisHostPort, redisPassword string, redisDB int) (*Repository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHostPort,
		Password: redisPassword,
		DB:       redisDB,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &Repository{
		client: client,
	}, nil
}
