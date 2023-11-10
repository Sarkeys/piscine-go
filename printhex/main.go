package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsValidArgument(s string) bool {
	for _, r := range s {
		if r > '9' || r < '0' {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	if !IsValidArgument(os.Args[1]) {
		return
	}
	var answer []rune
	hex := "0123456789abcdef"
	dec := "0123456789"
	k := 1
	input := os.Args[1]
	mydec := 0
	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < len(dec); j++ {
			if input[i] == dec[j] {
				mydec += j * k
				k *= 10
			}
		}
	}
	for mydec > 0 {
		answer = append(answer, rune(hex[mydec%16]))
		mydec /= 16
	}
	for i := len(answer) - 1; i >= 0; i-- {
		z01.PrintRune(answer[i])
	}
	z01.PrintRune('\n')
}
