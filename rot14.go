package piscine

func Rot14(s string) string {
	data := []rune(s)
	for i := 0; i < len(data); i++ {
		if data[i] > 64 && data[i] < 91 {
			data[i] += 14
			if data[i] > 90 {
				data[i] -= 26
			}
		} else if data[i] > 96 && data[i] < 123 {
			data[i] += 14
			if data[i] > 122 {
				data[i] -= 26
			}
		} else {
			continue
		}
	}
	s = string(data)
	return s
}
