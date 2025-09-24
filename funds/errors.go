package funds

import "fmt"

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

type InvalidAmountError struct {
	Amount int
}

func (e InvalidAmountError) Error() string {
	return fmt.Sprintf("invalid amount: %d (must be positive and not exceed 10000)", e.Amount)
}

type InvalidUserIDError struct {
	UserID int
}

func (e InvalidUserIDError) Error() string {
	return fmt.Sprintf("invalid user ID: %d (must be positive)", e.UserID)
}
