package list

import (
	"fmt"
	"log"
)

type Dlist struct {
	head *Dnode
	len  int
}

type Dnode struct {
	data interface{}
	prev *Dnode
	next *Dnode
}

func (l *Dlist) Append(d interface{}) {
	current := l.head
	node := &Dnode{data: d}
	if current == nil {
		l.head = node
		l.len++
		return
	}

	prev := current
	current = current.next
	for prev != nil {
		if current == nil {
			current = node
			prev.next = current
			current.prev = prev
			l.len++
			return
		}
		prev = current
		current = current.next
	}
}

// Print prints every list node.
func (l Dlist) Print() {
	s := l.listToSlice()
	fmt.Printf("%v\n", s)
}

// ToString returns string of list.
func (l Dlist) ToString() string {
	s := l.listToSlice()
	return fmt.Sprint(s)
}

func (l Dlist) listToSlice() []interface{} {
	current := l.head
	s := make([]interface{}, 0, l.len)
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
