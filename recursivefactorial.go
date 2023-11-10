package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		nb = 1
	} else if nb < 0 || nb > 21 {
		nb = 0
	} else {
		nb *= RecursiveFactorial(nb - 1)
	}
	return nb
}
