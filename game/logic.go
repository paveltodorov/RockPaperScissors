package game

func DecideWinner(choice1, choice2 string) int {
	if choice1 == choice2 {
		return -1
	}
	if (choice1 == "rock" && choice2 == "scissor") ||
		(choice1 == "paper" && choice2 == "rock") ||
		(choice1 == "scissor" && choice2 == "paper") {
		return 0
	}
	return 1
}
