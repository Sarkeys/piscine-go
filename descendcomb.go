package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			z01.PrintRune(rune(i/10) + 48)
			z01.PrintRune(rune(i%10) + 48)
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10 + 48))
			z01.PrintRune(rune(j%10 + 48))
			if !(i/10 == 0 && i%10 == 1 && j/10 == 0 && j%10 == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
