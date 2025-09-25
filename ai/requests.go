package ai

import (
	"rockpaperscissors/challenge"
	"rockpaperscissors/user"
)

type CreateAIPlayerRequest struct {
	Strategy string `json:"strategy" binding:"required,oneof=random defensive aggressive smart"`
}

type CreateAIPlayerResponse struct {
	User *user.User `json:"user"`
}

type AIPlayerResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
	Strategy string `json:"strategy"`
}

type AIChallengeRequest struct {
	AIID     int    `json:"ai_id" binding:"required"`
	PlayerID int    `json:"player_id" binding:"required"`
	Bet      int    `json:"bet" binding:"required,min=1"`
	Move     string `json:"move"`
}

type AIChallengeResponse struct {
	Accepted  bool                         `json:"accepted,omitempty"`
	Challenge *challenge.ChallengeResponse `json:"challenge,omitempty"`
	AIMove    string                       `json:"ai_move,omitempty"`
	Success   bool                         `json:"success"`
	Message   string                       `json:"message"`
}
