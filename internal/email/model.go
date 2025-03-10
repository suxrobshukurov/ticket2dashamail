package email

// Данные о билете из API
type Ticket struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Status  string `json:"status"`
	City    string `json:"city"`
	PaidAt  string `json:"paid_at"`
	EventID int    `json:"event_id"`
}

// Ответ от API билетов
type TicketsResponse struct {
	Tickets    []Ticket `json:"tickets"`
	TotalCount int      `json:"total_count"`
}

// Пользователь для Dashamail
type User struct {
	Email string `json:"email"`
	// City  string `json:"merge_4"`
	// Name  string `json:"merge_1"`
	// Phone string `json:"merge_3"`
}

// Запрос к Dashamail API
type DashamailRequest struct {
	Method string `json:"method"`
	ApiKey string `json:"api_key"`
	ListID string `json:"list_id"`
	Email  string `json:"email"`
	City   string `json:"merge_4"`
	Name   string `json:"merge_1"`
	Phone  string `json:"merge_3"`
}
