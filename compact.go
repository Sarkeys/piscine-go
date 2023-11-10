package piscine

func Compact(ptr *[]string) int {
	str := []string{}
	str = make([]string, 0)
	for _, char := range *ptr {
		if char != "" {
			str = append(str, char)
		}
	}
	*ptr = str
	return len(str)
}
