package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gosphere/internal/config"
	"gosphere/pkg/database"
)

func main() {
	fmt.Println("🚀 Pulse Backend API sistemleri başlatılıyor...")

	// 1. Konfigürasyonu Yükle
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Konfigürasyon yüklenemedi: %v", err)
	}
	log.Printf("Konfigürasyon yüklendi | Ortam: [%s] | Port: [%s]", cfg.AppEnv, cfg.AppPort)

	// 2. Otomatik Veritabanı Migrasyonlarını Çalıştır
	database.RunMigrations(cfg.GetDBConnString())

	// 3. PostgreSQL Bağlantı Havuzunu Kur
	ctx := context.Background()
	pgPool, err := database.NewPostgresPool(ctx, cfg.GetDBConnString())
	if err != nil {
		log.Fatalf("PostgreSQL bağlantı hatası: %v", err)
	}
	defer pgPool.Close()

	// 4. Redis Bağlantısını Kur
	redisClient, err := database.NewRedisClient(ctx, cfg)
	if err != nil {
		log.Fatalf("Redis bağlantı hatası: %v", err)
	}
	defer redisClient.Close()

	log.Println("Tüm altyapı sistemleri hazır! Sunucu istekleri beklemeye geçiyor...")

	// 5. Graceful Shutdown (Temiz Kapanış Mekanizması)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Sistem kapatılıyor, veritabanı ve Redis havuzları temizleniyor...")
}
