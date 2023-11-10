package piscine

import "github.com/01-edu/z01"

func Contains(r rune, data []rune) bool {
	for i := 0; i < len(data); i++ {
		if r == data[i] {
			return true
		}
	}
	return false
}

func IsValid(base string) bool {
	var data []rune
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if Contains(rune(base[i]), data) {
			return false
		}
		data = append(data, rune(base[i]))
		if base[i] == '+' || base[i] == '-' {
			return false
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if IsValid(base) {
		if nbr > 0 {
			var data []rune
			for nbr > 0 {
				data = append(data, rune(base[nbr%len(base)]))
				nbr /= len(base)
			}
			for i := len(data) - 1; i >= 0; i-- {
				z01.PrintRune(data[i])
			}
		}
		if nbr < 0 {
			var data []rune
			for nbr < 0 {
				data = append(data, rune(base[-(nbr%len(base))]))
				nbr /= len(base)
			}
			z01.PrintRune('-')
			for i := len(data) - 1; i >= 0; i-- {
				z01.PrintRune(data[i])
			}
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
