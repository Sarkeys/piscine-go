package piscine

func TrimAtoi(s string) int {
	Result := 0
	a := 1
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b := int(r - 48)
			Result = Result*10 + b
		} else if r == '-' && Result == 0 {
			a = -1
		}
	}
	return Result * a
}
