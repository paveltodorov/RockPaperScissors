package ai

import (
	"net/http"
	"rockpaperscissors/challenge"

	"github.com/gin-gonic/gin"
)

func CreateAIHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateAIPlayerRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		aiPlayer, err := s.CreateAIPlayer(req.Strategy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, CreateAIPlayerResponse{
			User: aiPlayer,
		})
	}
}

func ListAIHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		aiPlayers := s.ListAIUsers()

		var responses []*AIPlayerResponse
		for _, aiPlayer := range aiPlayers {
			responses = append(responses, &AIPlayerResponse{
				ID:       aiPlayer.ID,
				Username: aiPlayer.Username,
				Balance:  aiPlayer.Balance,
				Strategy: aiPlayer.Strategy,
			})
		}

		c.JSON(http.StatusOK, responses)
	}
}

func AIChallengeHandler(s *Service, challengeSvc *challenge.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AIChallengeRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
			return
		}

		// AI makes a move
		context := &GameContext{
			OpponentHistory: []string{}, // TO DO: add opponent history
			BetAmount:       req.Bet,
			Round:           1,
		}

		shouldAcceptChallenge := s.ShouldAcceptChallenge(req.Bet, req.Strategy)
		if !shouldAcceptChallenge {
			c.JSON(http.StatusOK, AIChallengeResponse{
				Accepted:  false,
				Success:   true,
				Message:   "AI should not accept challenge",
				Challenge: nil,
				AIMove:    "",
			})
			return
		}

		aiMove := s.MakeMove(req.Strategy, context)

		// Create challenge
		ch, err := challengeSvc.Create(req.AIID, req.PlayerID, req.Bet, aiMove.String())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ch.Status = "resolved" // an AI challenge is always resolved

		c.JSON(http.StatusOK, AIChallengeResponse{
			Accepted:  true,
			Success:   true,
			Message:   "AI challenge created",
			Challenge: ch.ToResponse(),
			AIMove:    aiMove.String(), // For testing purposes
		})
	}
}
