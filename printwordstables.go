package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, r := range a {
		s := r
		for _, l := range s {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
