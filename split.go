package piscine

func Split(s, sep string) []string {
	data := []rune(s)
	data2 := []rune(sep)
	f := 0
	for i := 0; i < len(sep); i++ {
		data = append(data, data2[i])
	}
	var array []string
	for j := 0; j < len(data); j++ {
		if data[j] == data2[0] {
			count := 0
			for k := 0; k < len(sep); k++ {
				if j+k < len(data) {
					if data[j+k] == data2[k] {
						count++
					}
				}
			}
			if count == len(sep) {
				temp := data[f:j]
				f = j + len(sep)
				if string(temp) != "" {
					array = append(array, string(temp))
				}
			}
		}
	}
	return array
}
