package main

import (
	"github.com/jumaniyozov/gonster/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
		winner = "Ww"
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := (currentRound % 3) == 0

	interaction.ShowAvailableActions(isSpecialRound)

	return ""
}

func endGame() {}
