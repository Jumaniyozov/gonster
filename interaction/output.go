package interaction

import "fmt"

type RoundData struct {
	Action string
	PlayerAttackDmg int
	PlayerHealValue int
	MonsterAttackDmg int
	PlayerHealth int
	MonsterHealth int
}

func PrintGreeting() {
	fmt.Println("...GONSTER...")
	fmt.Println("Game is initializing...")
	fmt.Println("Let's go!")
}

func ShowAvailableActions(isSpecialAttackAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("1: Attack Monster")
	fmt.Println("2: Heal")
	if isSpecialAttackAvailable {
		fmt.Println("3: Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v points.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player health %v.\n", roundData.PlayerHealth)
	fmt.Printf("Monster health %v.\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}
