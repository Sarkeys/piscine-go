package piscine

import "github.com/01-edu/z01"

func PrintCombw() {
	for i := 00; i < 99; i++ {
		for j := 0; j <= i+1; j++ {
			z01.PrintRune(rune(i/10) + 48)
			z01.PrintRune(rune(i%10) + 48)
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10 + 48))
			z01.PrintRune(rune(j%10 + 48))
			if !(i/10 == 9 && i%10 == 8 && j/10 == 9 && j%10 == 9) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
