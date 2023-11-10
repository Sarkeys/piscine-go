package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	text1 := os.Args[1]
	text2 := os.Args[2]
	check := make(map[rune]bool, len(text2))
	for _, r := range text2 {
		check[r] = true
	}
	for _, r := range text1 {
		if check[r] {
			z01.PrintRune(r)
			check[r] = false
		}
	}
}
