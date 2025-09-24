package challenge

import "fmt"

// Business logic errors
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

type InsufficientFundsError struct {
	UserID   int
	Balance  int
	Required int
}

func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("user %d has insufficient funds: balance=%d, required=%d", e.UserID, e.Balance, e.Required)
}

type UserNotFoundError struct {
	UserID int
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("user %d not found", e.UserID)
}

type ChallengeNotFoundError struct {
	ChallengeID int
}

func (e ChallengeNotFoundError) Error() string {
	return fmt.Sprintf("challenge %d not found", e.ChallengeID)
}

type InvalidChallengeStateError struct {
	ChallengeID   int
	CurrentState  string
	ExpectedState string
}

func (e InvalidChallengeStateError) Error() string {
	return fmt.Sprintf("challenge %d is in invalid state: current=%s, expected=%s", e.ChallengeID, e.CurrentState, e.ExpectedState)
}

type SelfChallengeError struct {
	UserID int
}

func (e SelfChallengeError) Error() string {
	return fmt.Sprintf("user %d cannot challenge themselves", e.UserID)
}

type InvalidMoveError struct {
	Move string
}

func (e InvalidMoveError) Error() string {
	return fmt.Sprintf("invalid move: %s (must be rock, paper, or scissors)", e.Move)
}
