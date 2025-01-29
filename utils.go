package main

import "fmt"

func GetWelcome() string {
	name := ""
	fmt.Println("Welcome to THE CASINO")
	fmt.Printf("Please enter your name: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return GetWelcome()
	}
	fmt.Printf("Welcome %s Lets Play!!!", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet or 0 to quit (balance: $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Insufficient funds. Please try again.")
		} else {
			break
		}
	}
	return bet

}