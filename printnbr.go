package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	c := 1
	if n < 0 {
		c = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * c
		if f != 0 {
			PrintNbr(f)
		}
		k := (n % 10 * c) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}
