package configs

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Dashamail    DashamailConfig
	Tickets      TicketsConfig
	RequestDelay time.Duration
}

type DashamailConfig struct {
	ApiKey string
	BaseID string
	ApiUrl string
}

type TicketsConfig struct {
	ApiUrl string
	ApiKey string
}

// Читаем `.env` и `config.json`
func NewConfig() *Config {
	// Загружаем переменные окружения из `.env`
	_ = godotenv.Load()

	// Читаем `config.json`
	b, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Ошибка чтения config.json: %v", err)
	}

	// Декодируем JSON
	var cfg struct {
		TicketsApiUrl   string `json:"tickets_api"`
		DashamailApiUrl string `json:"dashamail_api"`
		RequestDelay    int64  `json:"request_delay"`
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		log.Fatalf("Ошибка парсинга config.json: %v", err)
	}

	// Возвращаем структуру конфигурации
	return &Config{
		Dashamail: DashamailConfig{
			ApiKey: os.Getenv("DASHAMAIL_API_KEY"),
			BaseID: os.Getenv("DASHAMAIL_BASE_ID"),
			ApiUrl: cfg.DashamailApiUrl,
		},
		Tickets: TicketsConfig{
			ApiUrl: cfg.TicketsApiUrl,
			ApiKey: os.Getenv("TICKETS_API_KEY"),
		},
		RequestDelay: time.Duration(cfg.RequestDelay) * time.Minute,
	}
}
