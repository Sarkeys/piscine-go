package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		for j := 1; j < len(arguments)-1; j++ {
			if arguments[j] > arguments[j+1] {
				arguments[j+1], arguments[j] = arguments[j], arguments[j+1]
			}
		}
	}
	for i := 1; i < len(arguments); i++ {
		for _, r := range arguments[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
