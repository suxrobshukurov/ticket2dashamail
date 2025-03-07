package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logFile := "test_log.log"
	defer os.Remove(logFile) // Удаляем файл после теста

	InitLogger(logFile)

	Log.Info().Msg("Test log entry")

	// Проверяем, что файл создан
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Errorf("Файл логов не создан")
	}
}
