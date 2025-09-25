package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		u, err := s.Login(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		response := LoginResponse{User: u.ToResponse()}
		c.JSON(http.StatusOK, response)
	}
}

func LogoutHandler() gin.HandlerFunc {
	// for now, just return a success message
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "logged out"})
	}
}

func ListHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		users := s.List()
		var responses []*UserResponse
		for _, user := range users {
			responses = append(responses, user.ToResponse())
		}
		c.JSON(http.StatusOK, responses)
	}
}

func StatsHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := s.GetByID(c.GetInt("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		response := UserStatsResponse{
			Wins:   user.Stats.Wins,
			Losses: user.Stats.Losses,
			Ties:   user.Stats.Ties,

			RockChoices:     user.Stats.RockChoices,
			PaperChoices:    user.Stats.PaperChoices,
			ScissorsChoices: user.Stats.ScissorsChoices,

			AcceptedChallenges: user.Stats.AcceptedChallenges,
			DeclinedChallenges: user.Stats.DeclinedChallenges,
			CreatedChallenges:  user.Stats.CreatedChallenges,
		}

		c.JSON(http.StatusOK, response)
	}
}
