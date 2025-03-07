package logger

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger(logFile string) {
	// Создаём папку `logs/`, если её нет
	logDir := filepath.Dir(logFile)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}

	// Открываем лог-файл
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Ошибка открытия файла логов: " + err.Error())
	}

	// Читаем уровень логирования из переменной окружения
	logLevel := os.Getenv("LOG_LEVEL")
	level, err := strconv.Atoi(logLevel)
	if err != nil {
		level = int(zerolog.InfoLevel) // По умолчанию — INFO
	}

	zerolog.SetGlobalLevel(zerolog.Level(level))
	Log = zerolog.New(file).With().Timestamp().Logger()
	Log.Info().Msgf("Logger initialized with level: %d", level)
}

// Удаляет старые логи, если их больше maxFiles
func ManageLogs(logPath string, maxFiles int) {
	for {
		time.Sleep(24 * time.Hour) // Проверяем раз в день

		files, _ := filepath.Glob(filepath.Join(logPath, "*.log"))
		if len(files) > maxFiles {
			// Сортируем файлы по дате
			sort.Strings(files)
			for _, file := range files[:len(files)-maxFiles] {
				os.Remove(file)
			}
		}
	}
}
