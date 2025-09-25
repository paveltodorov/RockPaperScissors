package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
	Strategy string `json:"strategy"`
	IsAI     bool   `json:"is_ai"`
}
