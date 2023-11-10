package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	player1 := "Player 1: "
	player2 := "Player 2: "
	player3 := "Player 3: "
	player4 := "Player 4: "
	printt(player1)
	lena(0, 3, deck)
	z01.PrintRune('\n')
	printt(player2)
	lena(3, 6, deck)
	z01.PrintRune('\n')
	printt(player3)
	lena(6, 9, deck)
	z01.PrintRune('\n')
	printt(player4)
	lena(9, 12, deck)
	z01.PrintRune('\n')
}

func printt(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func lena(a, b int, deck []int) {
	for i := a; i < b; i++ {
		fmt.Printf("%v", deck[i])
		if i != b-1 {
			fmt.Printf(", ")
		}
	}
}
