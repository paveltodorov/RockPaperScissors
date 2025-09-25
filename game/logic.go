package game

type GameResult int

const (
	Tie GameResult = iota
	Player1Wins
	Player2Wins
)

func DecideWinner(move1, move2 Move) GameResult {
	if move1 == move2 {
		return Tie
	}

	if (move1 == Rock && move2 == Scissors) ||
		(move1 == Paper && move2 == Rock) ||
		(move1 == Scissors && move2 == Paper) {
		return Player1Wins
	}

	return Player2Wins
}

func DecideWinnerString(choice1, choice2 string) GameResult {
	move1, ok1 := ParseMove(choice1)
	move2, ok2 := ParseMove(choice2)

	if !ok1 || !ok2 {
		return Tie // Invalid moves result in tie
	}

	return DecideWinner(move1, move2)
}
