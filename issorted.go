package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	on, neg := true, true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			neg = false
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			on = false
		}
	}
	return on || neg
}
