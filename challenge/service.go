package challenge

import (
	"rockpaperscissors/game"
	"rockpaperscissors/user"
)

const (
	MaxBet = 10000
)

type Repository interface {
	AddChallenge(c *Challenge) int
	FindChallengeByID(id int) (*Challenge, bool)
	ListChallenges() []*Challenge
	UpdateChallenge(c *Challenge) error
}

type Service struct {
	repo  Repository
	users *user.Service
}

func NewService(r Repository, us *user.Service) *Service {
	return &Service{repo: r, users: us}
}

func (s *Service) Create(challengerID, opponentID, bet int, move string) (*Challenge, error) {
	if err := s.validateCreateInput(challengerID, opponentID, bet, move); err != nil {
		return nil, err
	}

	challenger, err := s.users.GetByID(challengerID)
	if err != nil {
		return nil, UserNotFoundError{UserID: challengerID}
	}

	opponent, err := s.users.GetByID(opponentID)
	if err != nil {
		return nil, UserNotFoundError{UserID: opponentID}
	}

	if challenger.Balance < bet {
		return nil, InsufficientFundsError{
			UserID:   challengerID,
			Balance:  challenger.Balance,
			Required: bet,
		}
	}

	if opponent.Balance < bet {
		return nil, InsufficientFundsError{
			UserID:   opponentID,
			Balance:  opponent.Balance,
			Required: bet,
		}
	}

	ch := &Challenge{
		ChallengerID:   challengerID,
		OpponentID:     opponentID,
		Bet:            bet,
		ChallengerMove: move,
		Status:         "pending",
	}
	s.repo.AddChallenge(ch)
	return ch, nil
}

func (s *Service) validateCreateInput(challengerID, opponentID, bet int, move string) error {
	if challengerID <= 0 {
		return ValidationError{Field: "challenger_id", Message: "must be positive"}
	}
	if opponentID <= 0 {
		return ValidationError{Field: "opponent_id", Message: "must be positive"}
	}

	if challengerID == opponentID {
		return SelfChallengeError{UserID: challengerID}
	}

	if bet <= 0 {
		return ValidationError{Field: "bet", Message: "must be positive"}
	}
	if bet > MaxBet {
		return ValidationError{Field: "bet", Message: "cannot exceed 10000"}
	}

	validMoves := map[string]bool{"rock": true, "paper": true, "scissors": true}
	if !validMoves[move] {
		return InvalidMoveError{Move: move}
	}

	return nil
}

func (s *Service) Accept(challengeID int, move string) (*Challenge, error) {
	if challengeID <= 0 {
		return nil, ValidationError{Field: "challenge_id", Message: "must be positive"}
	}

	validMoves := map[string]bool{"rock": true, "paper": true, "scissors": true}
	if !validMoves[move] {
		return nil, InvalidMoveError{Move: move}
	}

	ch, ok := s.repo.FindChallengeByID(challengeID)
	if !ok {
		return nil, ChallengeNotFoundError{ChallengeID: challengeID}
	}

	if ch.Status != "pending" {
		return nil, InvalidChallengeStateError{
			ChallengeID:   challengeID,
			CurrentState:  ch.Status,
			ExpectedState: "pending",
		}
	}

	opponent, err := s.users.GetByID(ch.OpponentID)
	if err != nil {
		return nil, UserNotFoundError{UserID: ch.OpponentID}
	}

	if opponent.Balance < ch.Bet {
		return nil, InsufficientFundsError{
			UserID:   ch.OpponentID,
			Balance:  opponent.Balance,
			Required: ch.Bet,
		}
	}

	ch.OpponentMove = move
	result := game.DecideWinner(ch.ChallengerMove, ch.OpponentMove)

	switch result {
	case -1:
		// Tie - no money transfer, both players keep their bets
		ch.Status = "resolved"
		ch.WinnerID = 0 // No winner in a tie
	case 0:
		// Challenger wins - gets both bets
		ch.Status = "resolved"
		ch.WinnerID = ch.ChallengerID
		winner, _ := s.users.GetByID(ch.ChallengerID)
		loser, _ := s.users.GetByID(ch.OpponentID)
		winner.Balance += ch.Bet * 2 // Winner gets both bets
		loser.Balance -= ch.Bet      // Loser loses their bet
		s.users.Update(winner)
		s.users.Update(loser)
	case 1:
		// Opponent wins - gets both bets
		ch.Status = "resolved"
		ch.WinnerID = ch.OpponentID
		winner, _ := s.users.GetByID(ch.OpponentID)
		loser, _ := s.users.GetByID(ch.ChallengerID)
		winner.Balance += ch.Bet * 2 // Winner gets both bets
		loser.Balance -= ch.Bet      // Loser loses their bet
		s.users.Update(winner)
		s.users.Update(loser)
	}

	s.repo.UpdateChallenge(ch)
	return ch, nil
}

func (s *Service) Decline(challengeID int) (*Challenge, error) {
	if challengeID <= 0 {
		return nil, ValidationError{Field: "challenge_id", Message: "must be positive"}
	}

	ch, ok := s.repo.FindChallengeByID(challengeID)
	if !ok {
		return nil, ChallengeNotFoundError{ChallengeID: challengeID}
	}

	if ch.Status != "pending" {
		return nil, InvalidChallengeStateError{
			ChallengeID:   challengeID,
			CurrentState:  ch.Status,
			ExpectedState: "pending",
		}
	}

	ch.Status = "declined"
	s.repo.UpdateChallenge(ch)
	return ch, nil
}

func (s *Service) List() []*Challenge {
	return s.repo.ListChallenges()
}
