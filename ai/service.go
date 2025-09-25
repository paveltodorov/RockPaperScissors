package ai

import (
	"math/rand"
	"rockpaperscissors/game"
	"rockpaperscissors/user"
	"time"
)

type Service struct {
	users *user.Service
}

func (s *Service) ListAIUsers() []*user.User {
	return s.users.ListAIUsers()
}

func NewService(users *user.Service) *Service {
	return &Service{users: users}
}

func (s *Service) CreateAIPlayer(strategy string) (*user.User, error) {
	username := s.generateAIUsername()

	u := &user.User{
		Username: username,
		Password: "ai_password", // AI players don't need real passwords
		Balance:  1000,          // AI players start with more money
		Strategy: strategy,
		IsAI:     true,
	}

	userID := s.users.AddUser(u)
	u.ID = userID
	return u, nil
}

func (s *Service) MakeMove(strategy string, context *GameContext) game.Move {
	switch strategy {
	case "random":
		return s.randomMove()
	case "defensive":
		return s.defensiveMove()
	case "aggressive":
		return s.aggressiveMove()
	case "smart":
		return s.smartMove(context)
	default:
		return s.randomMove()
	}
}

type GameContext struct {
	OpponentHistory []string // Previous moves by opponent
	BetAmount       int      // Current bet amount
	Round           int      // Current round number
}

// randomMove returns a random move
func (s *Service) randomMove() game.Move {
	rand.Seed(time.Now().UnixNano())
	moves := game.ValidMoves()
	return moves[rand.Intn(len(moves))]
}

// defensiveMove tends to play defensively
func (s *Service) defensiveMove() game.Move {
	// Defensive AI tends to play rock (beats scissors)
	rand.Seed(time.Now().UnixNano())
	if rand.Float32() < 0.4 {
		return game.Rock
	}
	return s.randomMove()
}

// aggressiveMove tends to play aggressively
func (s *Service) aggressiveMove() game.Move {
	// Aggressive AI tends to play scissors (beats paper)
	rand.Seed(time.Now().UnixNano())
	if rand.Float32() < 0.4 {
		return game.Scissors
	}
	return s.randomMove()
}

// smartMove uses basic pattern recognition
func (s *Service) smartMove(context *GameContext) game.Move {
	if len(context.OpponentHistory) == 0 {
		return s.randomMove()
	}

	// Count opponent's most common move
	moveCount := make(map[string]int)
	for _, move := range context.OpponentHistory {
		moveCount[move]++
	}

	// Find most common move and counter it
	var mostCommon string
	maxCount := 0
	for move, count := range moveCount {
		if count > maxCount {
			maxCount = count
			mostCommon = move
		}
	}

	// Counter the most common move
	switch mostCommon {
	case "rock":
		return game.Paper
	case "paper":
		return game.Scissors
	case "scissors":
		return game.Rock
	default:
		return s.randomMove()
	}
}

func (s *Service) generateAIUsername() string {
	rand.Seed(time.Now().UnixNano())
	aiNames := []string{"AI_Alpha", "AI_Beta", "AI_Gamma", "AI_Delta", "AI_Epsilon"}
	return aiNames[rand.Intn(len(aiNames))]
}

func (s *Service) ShouldAcceptChallenge(betAmount int, strategy string) bool {
	switch strategy {
	case "aggressive":
		return betAmount <= 200 // Accept higher bets
	case "defensive":
		return betAmount <= 50 // Only accept small bets
	case "smart":
		return betAmount <= 100 // Moderate risk
	default:
		return betAmount <= 100
	}
}
