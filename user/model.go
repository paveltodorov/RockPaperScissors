package user

type Stats struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Ties   int `json:"ties"`

	RockChoices     int `json:"rock_choices"`
	PaperChoices    int `json:"paper_choices"`
	ScissorsChoices int `json:"scissors_choices"`

	AcceptedChallenges int `json:"accepted_challenges"`
	DeclinedChallenges int `json:"declined_challenges"`
	CreatedChallenges  int `json:"created_challenges"`
}
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
	Strategy string `json:"strategy"`
	IsAI     bool   `json:"is_ai"`
	Stats    Stats  `json:"stats"`
}
