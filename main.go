package main

import (
	"github.com/jumaniyozov/gonster/actions"
	"github.com/jumaniyozov/gonster/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int
	var playerAttackValue int
	var playerHealValue int
	var monsterAttackValue int

	if userChoice == "ATTACK" {
		playerAttackValue = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackValue = actions.AttackMonster(true)
	}

	monsterAttackValue = actions.AttackPlayer()

	playerHealth, monsterHealth = actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackValue,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackValue,
	}

	interaction.PrintRoundStatistics(&roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
}
