package interaction

import "fmt"

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
