package piscine

func AlphaCount(s string) int {
	a := []rune(s)
	len := 0
	for i := range a {
		if a[i] >= 'A' && a[i] <= 'Z' || a[i] >= 'a' && a[i] <= 'z' {
			len++
		}
	}
	return len
}
