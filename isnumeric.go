package piscine

func IsNumeric(s string) bool {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if c[i] < 48 || c[i] > 57 {
			return false
		}
	}
	return true
}
