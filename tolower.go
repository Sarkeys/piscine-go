package piscine

func ToLower(s string) string {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if c[i] > 64 && c[i] < 91 {
			c[i] = c[i] + 32
		}
	}
	return string(c)
}
