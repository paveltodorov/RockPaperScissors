package user

import "fmt"

type UserNotFoundError struct {
	UserID int
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("user %d not found", e.UserID)
}

type InvalidCredentialsError struct {
	Username string
}

func (e InvalidCredentialsError) Error() string {
	return fmt.Sprintf("invalid credentials for user: %s", e.Username)
}

type InvalidUsernameError struct {
	Username string
}

func (e InvalidUsernameError) Error() string {
	return fmt.Sprintf("invalid username: %s (must be 3-20 characters, alphanumeric only)", e.Username)
}

type InvalidPasswordError struct {
	Message string
}

func (e InvalidPasswordError) Error() string {
	return fmt.Sprintf("invalid password: %s", e.Message)
}
