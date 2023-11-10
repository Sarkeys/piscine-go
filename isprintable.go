package piscine

func IsPrintable(s string) bool {
	c := []rune(s)
	b := true
	for i := 0; i < len(c); i++ {
		b = b && (c[i] > 31 && c[i] < 127)
	}
	return b
}
