package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"runtime"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialAttackAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && isSpecialAttackAvailable {
			return "SPECIAL_ATTACK"
		}

		fmt.Println("User input invalid. Please try again!")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}
	return userInput, nil
}
