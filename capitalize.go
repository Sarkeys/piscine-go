package piscine

func Capitalize(s string) string {
	c := []rune(s)
	if c[0] >= 'a' && c[0] <= 'z' {
		c[0] -= 32
	}
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if c[i+1] >= 'A' && c[i+1] <= 'Z' {
				c[i+1] += 32
			}
			if (c[i] < 'a' || c[i] > 'z') && (c[i] < 'A' || c[i] > 'Z') && (c[i] < '0' || c[i] > '9') {
				if c[i+1] >= 'a' && c[i+1] <= 'z' {
					c[i+1] -= 32
				}
			}
		}
	}
	return string(c)
}
