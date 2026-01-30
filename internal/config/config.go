package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI      string
	Port          string
	AppName       string
	AppEnv        string
	RedisAddr     string
	RedisUserName string
	RedisPassword string
	RedisDB       int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	cfg := &Config{
		AppName:       getEnv("APP_NAME", "user-auth"),
		AppEnv:        getEnv("APP_ENV", "production"),
		Port:          getEnv("PORT", "8080"),
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisUserName: getEnv("REDIS_USERNAME", ""),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
	}
	return cfg
}
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		var i int
		_, err := fmt.Sscanf(value, "%d", &i)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
