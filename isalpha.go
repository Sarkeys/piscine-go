package piscine

func IsAlpha(s string) bool {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if (c[i] < 97 || c[i] > 122) && (c[i] < 65 || c[i] > 90) && (c[i] < 48 || c[i] > 57) {
			return false
		}
	}
	return true
}
