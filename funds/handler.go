package funds

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DepositHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TransactionRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		u, err := s.Deposit(req.UserID, req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := TransactionResponse{
			Success: true,
			Message: "Deposit successful",
			Balance: u.Balance,
		}
		c.JSON(http.StatusOK, response)
	}
}

func WithdrawHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TransactionRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		u, err := s.Withdraw(req.UserID, req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := TransactionResponse{
			Success: true,
			Message: "Withdrawal successful",
			Balance: u.Balance,
		}
		c.JSON(http.StatusOK, response)
	}
}
