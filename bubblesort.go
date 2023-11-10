package piscine

func Bubble(array []int) []int {
	size := 5
	for i := 0; i < size-1; i++ {
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
			i = -1
		}
	}
	return array
}
