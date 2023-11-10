package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
	} else {
		var str string
		for i := 1; i < len(os.Args); i++ {
			str = str + " " + os.Args[i]
		}
		str1 := []rune(str)
		vowels := make([]rune, vowelsnum(str))
		indices := make([]int, vowelsnum(str))
		j := 0
		for i := 0; i < len(str1); i++ {
			if vowel(str1[i]) {
				vowels[j] = str1[i]
				indices[j] = i
				j++
			}
		}
		newvowels := revarr(vowels)
		for i, j := range indices {
			str1[j] = newvowels[i]
		}
		for i, r := range str1 {
			if i != 0 {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune('\n')
	}
}

func vowelsnum(s string) int {
	v := 0
	for _, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
			v = v + 1
		}
	}
	return v
}

func vowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

func revarr(str []rune) []rune {
	revstr := make([]rune, len(str))
	j := 0
	for i := len(str) - 1; i >= 0; i-- {
		revstr[j] = str[i]
		j++
	}
	return revstr
}
