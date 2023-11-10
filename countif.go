package piscine

func CountIf(f func(string) bool, tab []string) int {
	a := map2(f, tab)
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] {
			count++
		}
	}
	return count
}

func map3(f func(string) bool, a []string) []bool {
	b := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = f(a[i])
	}
	return b
}
