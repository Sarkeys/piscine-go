package piscine

func SplitWhiteSpaces(str string) []string {
	a := ""
	var b []string
	for i, v := range str {
		if i == len(str)-1 && string(v) != " " {
			a += string(v)
			b = append(b, a)
		} else if string(v) != " " {
			a += string(v)
		} else {
			if len(a) >= 1 {
				b = append(b, a)
			}
			a = ""
		}
	}
	return b
}
