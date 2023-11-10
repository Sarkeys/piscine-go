package piscine

func StrLen(s string) int {
	a := []rune(s)
	cout := 0
	for range a {
		cout++
	}
	return cout
}
