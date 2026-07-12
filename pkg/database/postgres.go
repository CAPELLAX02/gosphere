package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresPool, PostgreSQL için optimize edilmiş bir bağlantı havuzu oluşturur
func NewPostgresPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("Could not parse the configuration: %w", err)
	}

	// Yüksek Trafik Optimizasyonu (Connection Pool Ayarları)
	config.MaxConns = 25 				   	  // Aynı anda maksimum 25 bağlantı
	config.MinConns = 5 					  // Boşta bile en az 5 bağlantı hazır beklesin
	config.MaxConnLifetime = time.Hour 		  // Bir bağlantı maksimum 1 saat yaşasın
	config.MaxConnIdleTime = 30 * time.Minute // 30 dakika boşta kalan bağlantı kapatılsın

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("Could not create database pool: %w", err)
	}

	// Bağlantıyı test et (Ping)
	ctxTimeout, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	if err := pool.Ping(ctxTimeout); err != nil {
		return nil, fmt.Errorf("Could not reach database (Ping error): %w", err)
	}

	log.Println("PostgreSQL connection pool (pgxpool) extablished successfully!")
	return pool, nil
}
