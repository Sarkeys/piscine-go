package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	check := make(map[rune]bool)
	for _, char := range str1 {
		check[char] = true
	}
	for _, char := range str2 {
		check[char] = true
	}
	for _, char := range str1 {
		if check[char] {
			z01.PrintRune(char)
			check[char] = false
		}
	}
	for _, char := range str2 {
		if check[char] {
			z01.PrintRune(char)
			check[char] = false
		}
	}
}
