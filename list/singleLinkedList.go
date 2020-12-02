package list

import (
	"fmt"
	"log"
)

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

// ReverseIterative reverses the order of the elements in the list
func (l *List) ReverseIterative() {
	if l.len <= 1 {
		return
	}

	prev := l.head
	current := prev.next
	// headのnextをnilにする
	prev.next = nil
	for current != nil {
		// 次のnodeをnextに退避する
		next := current.next
		// currentの向き先を逆にする
		current.next = prev

		// 次のnodeの処理に移る
		prev = current
		current = next
	}
	l.head = prev
}

// ReverseRecursive is same as ReverseIterative
func (l *List) ReverseRecursive() {
	if l.len <= 1 {
		return
	}

	prev := l.head
	current := prev.next
	prev.next = nil

	var recF func(prev *Node, current *Node) *Node
	recF = func(prev *Node, current *Node) *Node {
		if current == nil {
			return prev
		}

		next := current.next
		current.next = prev
		return recF(current, next)
	}

	l.head = recF(prev, current)
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
	count := 0
	for current != nil {
		s = append(s, current.data)
		current = current.next
		count++
		if count > l.len {
			log.Fatalf("list size over error :%v\n", s)
		}
	}
	return s
}
