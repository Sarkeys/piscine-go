package piscine

func ToUpper(s string) string {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if c[i] > 96 && c[i] < 123 {
			c[i] = c[i] - 32
		}
	}
	return string(c)
}
