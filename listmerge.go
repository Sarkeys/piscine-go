package piscine

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head
	if l1 == nil || l2 == nil {
		return
	}
	if current == nil {
		l1.Head, l1.Tail = l2.Head, l2.Tail
		return
	}
	for current.Next != nil {
		current = current.Next
	}
	current.Next = l2.Head
}
