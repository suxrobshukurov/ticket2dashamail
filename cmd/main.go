package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ticket_app/configs"
	"ticket_app/internal/email"
	"ticket_app/internal/storage"
	"ticket_app/pkg/logger"

	"time"
)

func main() {
	fmt.Println("Service started...")

	// Загружаем конфигурацию
	cnf := configs.NewConfig()

	// Запускаем логирование в фоне
	logger.InitLogger("logs/app.log")
	go logger.ManageLogs("logs/", 5)

	// Инициализируем хранилище email'ов
	db := storage.NewFileDB("data/emails.log")

	// Запускаем обработку билетов в фоне
	go func() {
		for {
			logger.Log.Info().Msg("⏳ Начинаем обработку билетов...")
			processTickets(cnf, db)
			logger.Log.Info().Msg("✅ Обработка билетов завершена")

			time.Sleep(cnf.RequestDelay)
		}
	}()

	// Запускаем HTTP-сервер
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Сервис работает\n"))
	})
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info().Msg("🔄 Обработка запущена вручную через API")
		processTickets(cnf, db)
		w.Write([]byte("Обработка билетов завершена\n"))
	})

	server := &http.Server{Addr: ":8080"}
	go func() {
		logger.Log.Info().Msg("🌍 HTTP-сервер запущен на порту 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("Ошибка HTTP-сервера")
		}
	}()

	// Ожидание завершения программы (CTRL+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Log.Warn().Msg("❌ Сервис остановлен пользователем.")
	server.Close()
}

// processTickets получает билеты и отправляет email'ы
func processTickets(cnf *configs.Config, db *storage.FileDB) {
	// Берём билеты за последние 24 часа
	startTime := time.Now().Add(-24 * time.Hour).Format("2006-01-02T15-04-05")
	endTime := time.Now().Format("2006-01-02T15-04-05")

	logger.Log.Info().Msgf("📡 Запрашиваем билеты с %s по %s", startTime, endTime)

	users, err := email.GetUsers(*cnf, startTime, endTime)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Ошибка получения пользователей")
		return
	}

	logger.Log.Info().Msgf("📋 Найдено %d новых пользователей", len(users))

	for _, user := range users {
		if user.Email == "" {
			logger.Log.Warn().Msg("⚠️ Пропущен пользователь без email")
			continue
		}

		if db.EmailExists(user.Email) {
			logger.Log.Warn().Msgf("⚠️ Email уже существует: %s", user.Email)
			continue
		}

		emailSent, err := email.SendToDashamail(&cnf.Dashamail, user)
		if err != nil {
			logger.Log.Error().Err(err).Msgf("📛 Ошибка отправки email: %s", user.Email)
			continue
		}

		db.SaveEmail(emailSent)
		logger.Log.Info().Msgf("📨 Email отправлен и сохранён: %s", emailSent)
	}
}
