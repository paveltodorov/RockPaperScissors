package challenge

type CreateChallengeRequest struct {
	ChallengerID int    `json:"challenger_id" binding:"required"`
	OpponentID   int    `json:"opponent_id" binding:"required"`
	Bet          int    `json:"bet" binding:"required,min=1"` // Amount each player will bet (total pot = bet * 2)
	Move         string `json:"move" binding:"required,oneof=rock paper scissors"`
}

type AcceptChallengeRequest struct {
	ChallengeID int    `json:"challenge_id" binding:"required"`
	Move        string `json:"move" binding:"required,oneof=rock paper scissors"`
}

type DeclineChallengeRequest struct {
	ChallengeID int `json:"challenge_id" binding:"required"`
	OpponentID  int `json:"opponent_id" binding:"required"`
}

type ChallengeResponse struct {
	ID             int    `json:"id"`
	ChallengerID   int    `json:"challenger_id"`
	OpponentID     int    `json:"opponent_id"`
	Bet            int    `json:"bet"`
	ChallengerMove string `json:"challenger_move,omitempty"`
	OpponentMove   string `json:"opponent_move,omitempty"`
	Status         string `json:"status"`
	WinnerID       int    `json:"winner_id,omitempty"`
}

func (c *Challenge) ToResponse() *ChallengeResponse {
	return &ChallengeResponse{
		ID:             c.ID,
		ChallengerID:   c.ChallengerID,
		OpponentID:     c.OpponentID,
		Bet:            c.Bet,
		ChallengerMove: c.ChallengerMove,
		OpponentMove:   c.OpponentMove,
		Status:         c.Status,
		WinnerID:       c.WinnerID,
	}
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
