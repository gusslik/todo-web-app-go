package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
)

func mustGetEnv(fieldName string) string {
	val := os.Getenv(fieldName)
	if val == "" {
		log.Fatalf("Environment variable %s is not set", fieldName)
	}
	return val
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env:", err)
	}

	Host = mustGetEnv("DBHOST")
	Port = mustGetEnv("DBPORT")
	User = mustGetEnv("DBUSER")
	Password = mustGetEnv("DBPASSWORD")
	Dbname = mustGetEnv("DBNAME")
}
