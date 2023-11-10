package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return FindNextPrime(nb + 1)
	} else if nb == 2 {
		return nb
	}
	for i := 2; i*i < nb+1; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
