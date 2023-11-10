package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var a []int
	for i := min; i < max; i++ {
		a = append(a, i)
	}
	return a
}
