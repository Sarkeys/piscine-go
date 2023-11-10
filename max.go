package piscine

func Max(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	return a[len(a)-1]
}
