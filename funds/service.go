package funds

import (
	"rockpaperscissors/user"
)

type Service struct {
	users *user.Service
}

func NewService(u *user.Service) *Service {
	return &Service{users: u}
}

func (s *Service) Deposit(userID, amount int) (*user.User, error) {
	if userID <= 0 {
		return nil, InvalidUserIDError{UserID: userID}
	}
	if amount <= 0 {
		return nil, InvalidAmountError{Amount: amount}
	}

	u, err := s.users.GetByID(userID)
	if err != nil {
		return nil, UserNotFoundError{UserID: userID}
	}

	u.Balance += amount
	s.users.Update(u)
	return u, nil
}

func (s *Service) Withdraw(userID, amount int) (*user.User, error) {
	if userID <= 0 {
		return nil, InvalidUserIDError{UserID: userID}
	}
	if amount <= 0 {
		return nil, InvalidAmountError{Amount: amount}
	}

	u, err := s.users.GetByID(userID)
	if err != nil {
		return nil, UserNotFoundError{UserID: userID}
	}

	if u.Balance < amount {
		return nil, InsufficientFundsError{
			UserID:   userID,
			Balance:  u.Balance,
			Required: amount,
		}
	}

	u.Balance -= amount
	s.users.Update(u)
	return u, nil
}
