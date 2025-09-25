package user

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User *UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

type UserStatsResponse struct {
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

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Balance:  u.Balance,
	}
}
