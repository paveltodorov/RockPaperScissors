package challenge

type Challenge struct {
	ID             int    `json:"id"`
	ChallengerID   int    `json:"challengerId"`
	OpponentID     int    `json:"opponentId"`
	Bet            int    `json:"bet"` // Amount each player bets (total pot = bet * 2)
	ChallengerMove string `json:"challengerMove"`
	OpponentMove   string `json:"opponentMove"`
	Status         string `json:"status"`
	WinnerID       int    `json:"winnerId"` // 0 means tie (no winner)
}
