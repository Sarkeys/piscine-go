package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		n := l.Head
		for n.Next != nil {
			n = n.Next
		}
		return n.Data
	}
}
