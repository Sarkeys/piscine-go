package piscine

func BasicAtoi(s string) int {
	b := 0
	c := 0
	a := []rune(s)
	for _, word := range a {
		for i := '0'; i < word; i++ {
			c++
		}
		b = b*10 + c
		c = 0
	}
	return b
}
