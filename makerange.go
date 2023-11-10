package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	a := make([]int, max-min)
	for i := 0; i < len(a); i++ {
		a[i] = min + i
	}
	return a
}
