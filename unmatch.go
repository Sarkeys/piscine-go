package piscine

func Unmatch(a []int) int {
	for _, char := range a {
		count := 0
		for _, char2 := range a {
			if char == char2 {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return char
		}
	}
	return -1
}
