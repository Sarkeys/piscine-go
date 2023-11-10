package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		Nex := l.Head
		for Nex.Next != nil {
			Nex = Nex.Next
			// Nex = n
		}
		Nex.Next = n
	}
}
