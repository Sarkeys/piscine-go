package piscine

func IsLower(s string) bool {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if c[i] < 97 || c[i] > 122 {
			return false
		}
	}
	return true
}
