package funds

type TransactionRequest struct {
	UserID int `json:"user_id" binding:"required"`
	Amount int `json:"amount" binding:"required,min=1"`
}

type TransactionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Balance int    `json:"balance"`
}
