package storage

import (
	"bufio"
	"os"
	"sync"
)

type FileDB struct {
	filename string
	mu       sync.Mutex
}

func NewFileDB(filename string) *FileDB {
	// Создаём папку `data/`, если её нет
	if err := os.MkdirAll("data", 0755); err != nil {
		panic("Ошибка создания папки data: " + err.Error())
	}

	// Создаём файл, если его нет
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			panic("Ошибка создания файла email'ов: " + err.Error())
		}
		file.Close()
	}

	return &FileDB{filename: filename}
}

// Проверяем, существует ли email в файле
func (f *FileDB) EmailExists(email string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := os.Open(f.filename)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == email {
			return true
		}
	}
	return false
}

// Сохраняем email в файл
func (f *FileDB) SaveEmail(email string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Ошибка записи в файл email'ов: " + err.Error())
	}
	defer file.Close()

	file.WriteString(email + "\n")
}
