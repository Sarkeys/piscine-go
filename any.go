package piscine

func map2(f func(string) bool, a []string) []bool {
	b := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = f(a[i])
	}
	return b
}

func Any(f func(string) bool, a []string) bool {
	b := map2(f, a)
	for i := 0; i < len(b); i++ {
		if b[i] {
			return true
		}
	}
	return false
}
