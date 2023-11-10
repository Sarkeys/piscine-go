package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	data := []rune(os.Args[1])
	for i := 0; i < len(data); i++ {
		if data[i] != ' ' {
			z01.PrintRune(data[i])
			if i != len(data)-1 && (data[i+1] == ' ' && i+1 < len(data)-'\t') {
				z01.PrintRune('\t')
			}
		}
	}
	z01.PrintRune('\n')
}
