package email

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ticket_app/configs"
)

func TestGetUsers(t *testing.T) {
	// Создаём фейковый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"tickets": [{"email": "user@example.com", "city": "TestCity", "name": "TestName", "phone": "123456789"}]}`))
	}))
	defer server.Close()

	// Создаём тестовый конфиг
	config := configs.Config{
		Tickets: configs.TicketsConfig{
			ApiUrl: server.URL,
			ApiKey: "test_api_key",
		},
	}

	users, err := GetUsers(config, "2025-03-07T10:00:00", "2025-03-07T11:00:00")
	if err != nil {
		t.Fatalf("Ошибка получения пользователей: %v", err)
	}

	if len(users) != 1 || users[0].Email != "user@example.com" {
		t.Errorf("Ожидался 1 пользователь с email user@example.com, но получено: %+v", users)
	}
}
