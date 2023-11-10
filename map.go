package piscine

func Map(f func(int) bool, a []int) []bool {
	Bool := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		Bool[i] = f(a[i])
	}
	return Bool
}
