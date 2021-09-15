package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN
	maxAttackValue := PLAYER_ATTACK_MAX

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue
	return dmgValue
}

func HealPlayer() int {

	healValue := generateRandBetween(PLAYER_HEAL_MIN, PLAYER_HEAL_MAX)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() int {
	dmgValue := generateRandBetween(MONSTER_ATTACK_MIN, MONSTER_ATTACK_MAX)
	currentPlayerHealth -= dmgValue

	return dmgValue
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max - min) + min
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}