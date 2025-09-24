package user

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response payload for user login
type LoginResponse struct {
	User *UserResponse `json:"user"`
}

// UserResponse represents a user in API responses (without password)
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

// ToResponse converts a User to UserResponse (excludes password)
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Balance:  u.Balance,
	}
}
