package game

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

func (m Move) String() string {
	switch m {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	default:
		return "unknown"
	}
}

func ParseMove(s string) (Move, bool) {
	switch s {
	case "rock":
		return Rock, true
	case "paper":
		return Paper, true
	case "scissors":
		return Scissors, true
	default:
		return Rock, false
	}
}

func ValidMoves() []Move {
	return []Move{Rock, Paper, Scissors}
}

func ValidMoveStrings() []string {
	return []string{"rock", "paper", "scissors"}
}
