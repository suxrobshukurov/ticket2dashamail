package email

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ticket_app/configs"
)

func TestSendToDashamail(t *testing.T) {
	// Фейковый сервер Dashamail API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := configs.DashamailConfig{
		ApiKey: "test_api_key",
		BaseID: "test_base_id",
		ApiUrl: server.URL,
	}

	user := User{Email: "test@example.com", City: "TestCity", Name: "TestName", Phone: "123456789"}

	emailSent, err := SendToDashamail(&config, user)
	if err != nil {
		t.Fatalf("Ошибка отправки email: %v", err)
	}

	if emailSent != "test@example.com" {
		t.Errorf("Ожидался email test@example.com, но получен: %s", emailSent)
	}
}
