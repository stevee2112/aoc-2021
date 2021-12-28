package main

import (
	"fmt"
)

func main() {

	playersAt := map[int]int{
		0: 10, // player 1
		1: 1, // player 2
	}

	playerScores := map[int]int{
		0: 0, // player 1
		1: 0, // player 2
	}

	diceAt := 1

	done := false
	currentPlayer := 0
	rollCounter := 0

	for !done {
		sum := 0

		for i := 0; i < 3;i++ {
			val := 0
			val, diceAt = roleDice(diceAt)
			sum += val
			rollCounter++
		}

		playersAt[currentPlayer] = move(playersAt[currentPlayer], sum)

		playerScores[currentPlayer] += playersAt[currentPlayer]

		if playerScores[currentPlayer] >= 1000 {
			done = true
		}

		// Switch players
		currentPlayer = 1 - currentPlayer
	}

	fmt.Printf("Part 1: %d\n", rollCounter * playerScores[currentPlayer])
	fmt.Printf("Part 2: %d\n", 0)
}

func move(at int, moves int) int {

	boardMax := 10
	new := at + moves

	if new > boardMax {
		new = new % boardMax
	}

	if new == 0 {
		new = boardMax
	}

	return new
}

func roleDice(diceAt int) (int, int) {

	diceMax := 100

	role := 0
	role += diceAt
	diceAt++
	if diceAt > diceMax {
		diceAt = diceAt % diceMax
	}

	if diceAt == 0 {
		diceAt = diceMax
	}

	return role, diceAt
}
