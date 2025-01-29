package main

import (
	"fmt"
)

// 🏈 🏈 🏈
// 🎾 🎾 🎾
// ⚽️ ⚽️ ⚽️
// 🏐 🏐 🏐
// 🔥 🔥 🔥
// 🦀 🦀 🦀

func main() {

	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	arrayResult := GenerateSymbolArray(symbols)

	var balance uint
	balance = 200

	name := GetWelcome()
	fmt.Println("Game starts now!", name)
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := GetSpin(arrayResult, 3, 3)
		PrintSpin(spin)
		winningLines := CheckWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("Congratulations, you won $%d on line #%d\n", win, i+1)
			}

		}
	}
	fmt.Printf("You left with: $%d\n", balance)

}
