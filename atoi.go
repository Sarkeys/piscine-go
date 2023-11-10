package piscine

func Atoi(s string) int {
	var i int = 0
	if len(s) == 0 || s == "-" || s == "+" {
		return i
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || c == '-' || c == '+') {
			return i
		}
	}
	if s[0] == '-' && s[1] == '-' {
		return i
	}
	if s[0] == '-' && len(s) != 1 {
		s = s[1:]
		for _, c := range s {
			if !(c >= '0' && c <= '9') {
				return 0
			}
			i = i*10 + int(c-'0')
		}
		i = -i
	} else if s[0] == '+' && len(s) != 1 {
		s = s[1:]
		for _, c := range s {
			if !(c >= '0' && c <= '9') {
				return 0
			}
			i = i*10 + int(c-'0')
		}
	} else {
		for _, c := range s {
			if !(c >= '0' && c <= '9') {
				return 0
			}
			i = i*10 + int(c-'0')
		}
	}
	return i
}
