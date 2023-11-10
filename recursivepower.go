package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		nb = 1
	}
	if power < 0 {
		nb = 0
	} else if power > 0 {
		nb *= RecursivePower(nb, power-1)
	}
	return nb
}
