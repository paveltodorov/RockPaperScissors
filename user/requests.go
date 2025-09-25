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

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Balance:  u.Balance,
	}
}
