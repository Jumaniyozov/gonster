package main

import (
	"fmt"
	"github.com/jumaniyozov/gonster/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound % 3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {

	} else if userChoice == "HEAL" {

	} else {

	}

	return ""
}

func endGame() {}
