package piscine

func BasicAtoi2(s string) int {
	b := 0
	a := []rune(s)
	n := len(s)
	for i := 0; i < n; i++ {
		if a[i] < '0' || a[i] > '9' {
			return 0
		} else {
			b *= 10
			b += int(a[i]) - '0'
		}
	}
	return b
}
