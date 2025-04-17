package responses

type ErrorResponse struct {
	Timestamp int    `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}
