package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	data := []rune(arguments[0])
	for i := 2; i < len(data); i++ {
		z01.PrintRune(data[i])
	}
	z01.PrintRune('\n')
}
