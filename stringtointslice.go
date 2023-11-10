package piscine

func StringToIntSlice(str string) []int {
	num := []int{}
	if len(str) == 0 {
		num = nil
	}
	for _, char := range str {
		num = append(num, int(rune(char)))
	}
	return num
}
