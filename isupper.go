package piscine

func IsUpper(s string) bool {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if c[i] < 65 || c[i] > 90 {
			return false
		}
	}
	return true
}
