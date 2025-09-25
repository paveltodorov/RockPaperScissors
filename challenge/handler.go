package challenge

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateChallengeRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		ch, err := s.Create(req.ChallengerID, req.OpponentID, req.Bet, req.Move)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := APIResponse{
			Success: true,
			Message: "Challenge created successfully",
			Data:    ch.ToResponse(),
		}
		c.JSON(http.StatusOK, response)
	}
}

func ListHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		challenges := s.List()
		var responses []*ChallengeResponse
		for _, ch := range challenges {
			responses = append(responses, ch.ToResponse())
		}
		c.JSON(http.StatusOK, responses)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		id, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
			return
		}
		c.Set("user_id", id)
		c.Next()
	}
}

func ListPendingHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		challenges, err := s.ListPendingByUserID(c.GetInt("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var responses []*ChallengeResponse
		for _, ch := range challenges {
			responses = append(responses, ch.ToResponse())
		}
		c.JSON(http.StatusOK, responses)
	}
}

func AcceptHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AcceptChallengeRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		ch, err := s.Accept(req.ChallengeID, req.Move)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := APIResponse{
			Success: true,
			Message: "Challenge accepted successfully",
			Data:    ch.ToResponse(),
		}
		c.JSON(http.StatusOK, response)
	}
}

func DeclineHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DeclineChallengeRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		ch, err := s.Decline(req.ChallengeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := APIResponse{
			Success: true,
			Message: "Challenge declined successfully",
			Data:    ch.ToResponse(),
		}
		c.JSON(http.StatusOK, response)
	}
}
