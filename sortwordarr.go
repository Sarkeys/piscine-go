package piscine

func SortWordArr(a []string) {
	swap := true
	for swap {
		swap = false
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				swap = true
			}
		}
	}
}
