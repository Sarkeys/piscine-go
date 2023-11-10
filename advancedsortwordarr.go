package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	swap := true
	for swap {
		swap = false
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) > 0 {
				a[i], a[i-1] = a[i-1], a[i]
				swap = true
			}
		}
	}
}
