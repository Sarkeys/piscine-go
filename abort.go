package piscine

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	Bubble(array)
	return array[2]
}
