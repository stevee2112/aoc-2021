package main

import (
	"fmt"
	"strings"
	"stevee2112/aoc-2021/util"
)

func main() {

	// Part 1
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

	// Part 2

	gameState := map[string]int{}
	initialState := makeStateSlug(10, 0, 1, 0)

	gameState[initialState] = 1

	// All possible sums for 3 rolls
	sumMap := map[int]int{}
	// Die Sums
	for i := 1;i <= 3;i++ {
		for j := 1;j <= 3;j++ {
			for k := 1;k <= 3;k++ {
				sum := i + j + k
				sumMap[sum]++
			}
		}
	}

	turns := 50 // Should be more than enough
	totalWins := map[int]int{}

	playerTurn := 0
	for i := 0;i < turns;i++ {
		wins := 0
		gameState, wins = takeTurn(playerTurn + 1, gameState, sumMap)
		totalWins[playerTurn + 1] += wins
		playerTurn = 1 - playerTurn
	}

	mostWins := totalWins[1]

	if totalWins[2] > mostWins {
		mostWins = totalWins[2]
	}

	fmt.Printf("Part 2: %d\n", mostWins)
}

func takeTurn(player int, gameState map[string]int, sumMap map[int]int) (map[string]int, int) {

	wins := 0
	targetScore := 21

	newGameState := map[string]int{}
	for state, stateCount := range gameState {

		if player == 1 {
			player1At, player1Score, player2At, player2Score := parseStateSlug(state)

			for sum, sumCount := range sumMap {
				move :=  move(player1At, sum)
				score := player1Score + move

				if score >= targetScore {
					wins += (sumCount * stateCount)
				} else {
					newGameState[makeStateSlug(move, score, player2At, player2Score)] += sumCount * stateCount
				}
			}
		}

		if player == 2 {
			player1At, player1Score, player2At, player2Score := parseStateSlug(state)

			for sum, sumCount := range sumMap {
				move :=  move(player2At, sum)
				score := player2Score + move

				if score >= targetScore {
					wins += (sumCount * stateCount)
				} else {
					newGameState[makeStateSlug(player1At, player1Score, move, score)] += sumCount * stateCount
				}
			}
		}
	}

	return newGameState, wins
}

func printGameState(gameState map[string]int) {
	for state, count := range gameState {
		player1At, player1Score, player2At, player2Score := parseStateSlug(state)
		fmt.Printf(
			"Player 1 at: %d, score: %d - Player 2 at %d, score: %d (count %d)\n",
			player1At, player1Score, player2At, player2Score, count,
		)
	}
}

func makeStateSlug(player1At int, player1Score int, player2At int, player2Score int) string {
	return fmt.Sprintf("%d:%d:%d:%d", player1At, player1Score, player2At, player2Score)
}

func parseStateSlug(state string) (player1At int, player1Score int, player2At int, player2Score int) {

	parts := strings.Split(state,":")

	return util.Atoi(parts[0]), util.Atoi(parts[1]), util.Atoi(parts[2]), util.Atoi(parts[3])
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
