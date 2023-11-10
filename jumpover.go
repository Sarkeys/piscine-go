package piscine

func JumpOver(str string) string {
	strk := ""
	for i := 2; i < len(str); i += 3 {
		if len(str) > 2 {
			strk += string(str[i])
		}
	}
	return strk + "\n"
}
