package piscine

func Index(s string, toFind string) int {
	count := 0
	if len(toFind) == 0 || len(toFind) > len(s) {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			for j := 0; j < len(toFind); j++ {
				if i+j < len(s) {
					if s[i+j] == toFind[j] {
						count++
					}
				}
			}
			if count == len(toFind) {
				return i
			}
		}
	}
	return -1
}
