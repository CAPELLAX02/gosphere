package database

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations veritabanı şemasını en güncel versiyona yükseltir
func RunMigrations(connString string) {
	m, err := migrate.New("file://migrations", connString)
	if err != nil {
		log.Fatalf("Migrasyon nesnesi oluşturulamadı: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrasyon çalıştırılamadı: %v", err)
	}

	log.Println("Veritabanı şeması doğrulandı / en güncel versiyona yükseltildi!")
}