package config

import (
	"os"
	"strconv"
)

// AppConfig represents the application's configuration.
type AppConfig struct {
	DatabasePath     string
	RSSFetchLimit    int
	ArticleSendLimit int
	SMTPServer       string
	SMTPUser         string
	SMTPPassword     string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		DatabasePath:     getEnv("DATABASE_PATH", "../database/newsletter-aggregator.db"),
		RSSFetchLimit:    getIntEnv("RSS_FETCH_LIMIT", 50),
		ArticleSendLimit: getIntEnv("ARTICLE_SEND_LIMIT", 3),
		SMTPServer:       getEnv("SMTP_SERVER", "smtp.your-email-provider.com:port"),
		SMTPUser:         getEnv("SMTP_USER", "your_email@example.com"),
		SMTPPassword:     getEnv("SMTP_PASSWORD", "your_email_password"),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func getIntEnv(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intValue
}