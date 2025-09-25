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

	challenger.Stats.CreatedChallenges++
	switch move {
	case "rock":
		challenger.Stats.RockChoices++
	case "paper":
		challenger.Stats.PaperChoices++
	case "scissors":
		challenger.Stats.ScissorsChoices++
	}

	ch := &Challenge{
		ChallengerID:   challengerID,
		OpponentID:     opponentID,
		Bet:            bet,
		ChallengerMove: move,
		Status:         "pending",
	}
	chId := s.repo.AddChallenge(ch)
	ch.ID = chId
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
	switch move {
	case "rock":
		opponent.Stats.RockChoices++
	case "paper":
		opponent.Stats.PaperChoices++
	case "scissors":
		opponent.Stats.ScissorsChoices++
	}

	result := game.DecideWinnerString(ch.ChallengerMove, ch.OpponentMove)
	challenger, _ := s.users.GetByID(ch.ChallengerID)

	opponent.Stats.AcceptedChallenges++

	switch result {
	case game.Tie:
		// Tie - no money transfer, both players keep their bets
		ch.Status = "resolved"
		ch.WinnerID = 0 // No winner in a tie
		challenger.Stats.Ties++
		opponent.Stats.Ties++
		s.users.Update(challenger)
		s.users.Update(opponent)
	case game.Player1Wins:
		// Challenger wins - gets both bets
		ch.Status = "resolved"
		ch.WinnerID = ch.ChallengerID
		winner := challenger
		loser := opponent
		winner.Balance += ch.Bet * 2 // Winner gets both bets
		loser.Balance -= ch.Bet      // Loser loses their bet
		winner.Stats.Wins++
		loser.Stats.Losses++
		s.users.Update(winner)
		s.users.Update(loser)
	case game.Player2Wins:
		// Opponent wins - gets both bets
		ch.Status = "resolved"
		ch.WinnerID = ch.OpponentID
		winner := opponent
		loser := challenger
		winner.Balance += ch.Bet * 2 // Winner gets both bets
		loser.Balance -= ch.Bet      // Loser loses their bet
		winner.Stats.Wins++
		loser.Stats.Losses++
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

func (s *Service) ListPendingByUserID(userID int) ([]*Challenge, error) {
	challenges := s.repo.ListChallenges()
	var pendingChallenges []*Challenge
	for _, challenge := range challenges {
		if challenge.Status == "pending" && (challenge.ChallengerID == userID || challenge.OpponentID == userID) {
			pendingChallenges = append(pendingChallenges, challenge)
		}
	}
	return pendingChallenges, nil
}
