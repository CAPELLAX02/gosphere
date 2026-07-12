package main

import (
	"fmt"
	"log"
	"gosphere/internal/config"
)

func main() {
	fmt.Println("Running Go Backend API...")

	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load the configuration: %v", err)
	}
	
	log.Printf("Configuration loaded successfully. Working environment [%s] | Port: [%s]", cfg.AppEnv, cfg.AppPort)
	log.Printf("Database Target: %s@%s:%s/%s", cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)
}