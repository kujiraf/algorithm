package list

import "fmt"

// List represents a single linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	head *Node
	len  int
}

// Node is an node of a linked list.
type Node struct {
	data interface{}
	next *Node
}

// Append adds a node to the end of list.
func (l *List) Append(n *Node) {
	if l.head == nil {
		l.head = n
		l.len++
		return
	}

	current := l.head
	for i := 0; i < l.len; i++ {
		if current.next == nil {
			current.next = n
			l.len++
			return
		}
		current = current.next
	}
}

// Prepend adds new node to the top of the list.
func (l *List) Prepend(n *Node) {
	if l.head == nil {
		l.head = n
		l.len++
		return
	}

	tmp := l.head
	l.head = n
	l.head.next = tmp
	l.len++
}

// Print prints every list node.
func (l List) Print() {
	s := l.listToSlice()
	fmt.Printf("%v\n", s)
}

// ToString returns string of list.
func (l List) ToString() string {
	s := l.listToSlice()
	return fmt.Sprint(s)
}

func (l List) listToSlice() []interface{} {
	s := make([]interface{}, 0, l.len)
	current := l.head
	for i := 0; i < l.len; i++ {
		s = append(s, current.data)
		current = current.next
	}
	return s
}
