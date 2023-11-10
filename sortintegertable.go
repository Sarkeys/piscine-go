package piscine

func SortIntegerTable(table []int) {
	for i := range table {
		if i == len(table)-1 {
			continue
		}
		for j := range table {
			if j == len(table)-1 {
				continue
			}
			if table[j] > table[j+1] {
				t := table[j]
				table[j] = table[j+1]
				table[j+1] = t
			}
		}
	}
}
