package email

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"ticket_app/configs"
	"ticket_app/pkg/logger"
	"time"
)

func SendToDashamail(configs *configs.DashamailConfig, user User) (string, error) {
	payload := DashamailRequest{
		Method: "lists.add_member",
		ApiKey: configs.ApiKey,
		ListID: configs.BaseID,
		Email:  user.Email,
		// City:   user.City,
		// Name:   user.Name,
		// Phone:  user.Phone,
	}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", configs.ApiUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Error().Err(err).Msgf("📛 Ошибка запроса в Dashamail для email: %s", user.Email)
		return "", err
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, _ := io.ReadAll(resp.Body)

	// Парсим JSON-ответ
	var dashamailResponse struct {
		Response struct {
			Msg struct {
				ErrCode int    `json:"err_code"`
				Text    string `json:"text"`
			} `json:"msg"`
			Data struct {
				MemberID    int    `json:"member_id"`
				FixedEmail  bool   `json:"fixed_email"`
				SendConfirm string `json:"send_confirm"`
			} `json:"data"`
		} `json:"response"`
	}
	json.Unmarshal(body, &dashamailResponse)

	// Логируем только важные данные
	if resp.StatusCode == http.StatusOK && dashamailResponse.Response.Msg.ErrCode == 0 {
		logger.Log.Info().Msgf("✅ Email успешно добавлен в Dashamail: %s (ID: %d)", user.Email, dashamailResponse.Response.Data.MemberID)
	} else {
		logger.Log.Warn().Msgf("⚠️ Dashamail вернул ошибку: %s (код: %d)", dashamailResponse.Response.Msg.Text, dashamailResponse.Response.Msg.ErrCode)
	}

	return user.Email, nil
}
