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

func (l *Dlist) Prepend(d interface{}) {
	node := &Dnode{data: d}

	l.head.prev = node
	node.next = l.head
	l.head = node
	l.len++
}

// Print prints every list node.
func (l Dlist) Print() {
	s := l.listToSlice()
	fmt.Printf("%v\n", s)
}

func (l *Dlist) Remove(d interface{}) {
	if l.len == 0 {
		return
	}

	if l.head.data == d {
		if l.len == 1 {
			l.head = nil
			l.len--
			return
		}
		l.head = l.head.next
		l.head.prev = nil
		l.len--
		return
	}

	var prev *Dnode
	current := l.head
	for current != nil {
		if current.data == d {
			if prev != nil {
				prev.next = current.next
			}
			if current.next != nil {
				current.next.prev = prev
			}
			l.len--
			return
		}
		prev = current
		current = current.next
	}
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
			log.Fatalf("to string error: list size over error :%v\n", s)
		}
	}
	return s
}

func (l Dlist) ToStringReverse() string {
	current := l.head
	var end *Dnode
	count := 0
	for current != nil {
		end = current
		current = current.next
		count++
		if count > l.len {
			log.Fatalf("to string reverse error: list size over error. len=%d\n", l.len)
		}
	}

	count = 0
	current = end
	s := make([]interface{}, 0, l.len)
	for current != nil {
		s = append(s, current.data)
		current = current.prev
	}

	return fmt.Sprint(s)
}
