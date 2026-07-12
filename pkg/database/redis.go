package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gosphere/internal/config"
)

func NewRedisClient(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost,
		Password: cfg.RedisPass,
		DB: cfg.RedisDB,
		PoolSize: 20, // Redis bağlantı havuzu
	})

	ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if _, err := client.Ping(ctxTimeout).Result(); err != nil {
		return nil, fmt.Errorf("redis sunucusuna bağlanılamadı: %w", err)
	}

	log.Println("Redis önbellek sunucusuna başarıyla bağlanıldı!")
	return client, nil
}