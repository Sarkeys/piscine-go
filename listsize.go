package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	} else {
		n := l.Head
		count := 1
		for n.Next != nil {
			n = n.Next
			count++
		}
		return count
	}
}
