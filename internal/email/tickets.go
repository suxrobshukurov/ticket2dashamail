package email

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ticket_app/configs"
)

// Получает список пользователей, купивших билеты за указанный период
func GetUsers(config configs.Config, startTime, endTime string) ([]User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Tickets.ApiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+config.Tickets.ApiKey)

	// Устанавливаем параметры запроса
	q := req.URL.Query()
	q.Add("start_paid_at", startTime)
	q.Add("end_paid_at", endTime)
	q.Add("status", "bought")
	q.Add("per_page", "1000")
	req.URL.RawQuery = q.Encode()

	// Выполняем запрос
	// logger.Log.Info().Msgf("📡 Отправляем запрос: %s", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Проверяем статус-код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Распарсим JSON
	var ticketResponse TicketsResponse
	if err := json.Unmarshal(body, &ticketResponse); err != nil {
		return nil, err
	}

	// Преобразуем билеты в пользователей
	var users []User
	for _, t := range ticketResponse.Tickets {
		if t.Email == "" {
			// logger.Log.Warn().Msgf("⚠️ В API получен билет без email: %+v", t)
			continue
		}

		users = append(users, User{
			Email: t.Email,
			City:  t.City,
			Name:  t.Name,
			Phone: t.Phone,
		})
	}

	return users, nil
}
