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

		aiPlayer, err := s.users.GetByID(req.AIID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		challenger, err := s.users.GetByID(req.PlayerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context := &GameContext{
			OpponentStats: challenger.Stats,
			BetAmount:     req.Bet,
			Round:         1,
		}

		// Create challenge
		ch, err := challengeSvc.Create(req.PlayerID, req.AIID, req.Bet, req.Move)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		shouldAcceptChallenge := s.ShouldAcceptChallenge(req.Bet, aiPlayer.Strategy)
		if !shouldAcceptChallenge {
			_, err := challengeSvc.Decline(ch.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, AIChallengeResponse{
				Accepted:  false,
				Success:   true,
				Message:   "AI should not accept challenge",
				Challenge: nil,
				AIMove:    "",
			})
		}

		aiMove := s.MakeMove(aiPlayer.Strategy, context)
		_, err = challengeSvc.Accept(ch.ID, aiMove.String())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, AIChallengeResponse{
			Accepted:  true,
			Success:   true,
			Message:   "AI challenge accepted",
			Challenge: ch.ToResponse(),
			AIMove:    aiMove.String(),
		})
	}
}
