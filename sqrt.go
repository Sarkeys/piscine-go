package piscine

func Sqrt(nb int) int {
	for i := 1; i*i < nb+1; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
