package piscine

func LastRune(s string) rune {
	d := []rune(s)
	return d[len(s)-1]
}
