package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	current := l.Head
	var prev *NodeL
	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next, current = current.Next, prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
