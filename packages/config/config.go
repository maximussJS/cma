package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port            string
	BasePath        string
	SlackWebhookUrl string
}

var GlobalConfig *Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println(`Environment variable "PORT" not found, used default :8080`)
		port = ":8080"
	}

	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		log.Println(`Environment variable "BASE_PATH" not found, used default data`)
		basePath = "data"
	}

	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")

	if slackWebhookUrl == "" {
		log.Fatalf(`Environment variable "SLACK_WEBHOOK_URL" not found, update .env file`)
	}

	GlobalConfig = &Config{
		Port:            port,
		BasePath:        basePath,
		SlackWebhookUrl: slackWebhookUrl,
	}
}
