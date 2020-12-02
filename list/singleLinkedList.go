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
func (l *List) Append(d interface{}) {
	if l.head == nil {
		l.head = &Node{data: d}
		l.len++
		return
	}

	current := l.head
	for current != nil {
		if current.next == nil {
			current.next = &Node{data: d}
			l.len++
			return
		}
		current = current.next
	}
}

// Prepend adds new node to the top of the list.
func (l *List) Prepend(d interface{}) {
	if l.head == nil {
		l.head = &Node{data: d}
		l.len++
		return
	}

	newHead := &Node{
		data: d,
		next: l.head,
	}
	l.head = newHead
	l.len++
}

// RemoveFirst removes only the first data from the list if args data is found in it.
func (l *List) RemoveFirst(d interface{}) {
	if l.len == 0 {
		return
	}

	// the case the head data is equal d
	if l.head.data == d {
		l.head = l.head.next
		l.len--
		return
	}

	prev := l.head
	current := prev.next
	for current != nil {
		if current.data == d {
			// headを持っているのがprevなので、prev.nextをcurrent.nextにする
			// current = current.nextだと、currentはheadを持っていないので、リムーブできない。
			prev.next = current.next
			l.len--
			return
		}
		prev = current
		current = current.next
	}
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
	for current != nil {
		s = append(s, current.data)
		current = current.next
	}
	return s
}
