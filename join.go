package piscine

func Join(strs []string, sep string) string {
	var a string
	for i := 0; i < len(strs); i++ {
		a += strs[i]
		if i < len(strs)-1 {
			a += sep
		}
	}
	return a
}
