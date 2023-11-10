package piscine

import "github.com/01-edu/z01"

var (
	a [8]rune
	b [9]bool
	f [9]int
)

func EightQueens() {
	count := 0
	for i := 1; i <= 8; i++ {
		if b[i] {
			count++
		}
	}

	if count == 8 {
		for _, c := range a {
			z01.PrintRune(c)
		}
		z01.PrintRune(rune(10))
		return
	}

	for i := rune(49); i <= rune(56); i++ {
		d := 0
		for j := rune(49); j <= i; j++ {
			d++
		}
		if !b[d] {
			g := true
			for j := 1; j <= count; j++ {
				if d == f[j]-(count+1-j) || d == f[j]+(count+1-j) {
					g = false
					break
				}
			}
			if g {
				b[d] = true
				a[count] = i
				f[count+1] = d
				EightQueens()
				b[d] = false
			}
		}
	}
}
