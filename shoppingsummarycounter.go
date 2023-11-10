package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	a := make(map[string]int)
	s := ""
	for _, char := range str {
		if char == ' ' {
			a[s]++
			s = ""
		} else {
			s += string(char)
		}
	}
	a[s]++
	return a
}
