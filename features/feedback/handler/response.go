package handler

type feedbackResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
	Notes  string `json:"notes"`
}
