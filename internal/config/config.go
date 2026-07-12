package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string `mapstructure:"APP_PORT"`
	AppEnv      string `mapstructure:"APP_ENV"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	DBSSLMode   string `mapstructure:"DB_SSL_MODE"`
	RedisHost   string `mapstructure:"REDIS_HOST"`
	RedisPass   string `mapstructure:"REDIS_PASSWORD"`
	RedisDB     int    `mapstructure:"REDIS_DB"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	JWTExpire   int    `mapstructure:"JWT_EXPIRE_HOURS"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Could not read .env file, using system environment variables: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetDBConnString PostgreSQL için connection string üretir
func (c *Config) GetDBConnString() string {
	return "host=" + c.DBHost +
			" port=" + c.DBPort +
			" user=" + c.DBUser +
			" password=" + c.DBPassword +
			" dbname=" + c.DBName +
			" sslmode=" + c.DBSSLMode
}