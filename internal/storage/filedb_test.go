package storage

import (
	"os"
	"testing"
)

func TestFileDB(t *testing.T) {
	testFile := "test_emails.log"
	defer os.Remove(testFile) // Удаляем тестовый файл после теста

	db := NewFileDB(testFile)

	// Проверяем, что email не существует
	if db.EmailExists("test@example.com") {
		t.Errorf("Email должен отсутствовать, но найден")
	}

	// Сохраняем email
	db.SaveEmail("test@example.com")

	// Проверяем, что email теперь существует
	if !db.EmailExists("test@example.com") {
		t.Errorf("Email должен существовать, но не найден")
	}
}
