package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var n *NodeI
	if l == nil || l.Next == nil {
		return l
	}
	n = l
	for l != nil {
		next := l.Next
		for next != nil {
			if l.Data > next.Data {
				l.Data, next.Data = next.Data, l.Data
			}
			next = next.Next
		}
		l = l.Next
	}
	return n
}
