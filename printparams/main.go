package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		for _, r := range arguments[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
