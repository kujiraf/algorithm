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

	prev := l.head
	current := prev.next
	for current != nil {
		if current.data == d {
			prev.next = current.next
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

func (l *Dlist) ReverseIterative() {
	if l.len <= 1 {
		return
	}

	var prev *Dnode
	current := l.head
	for current != nil {
		next := current.next
		current.prev = current.next
		current.next = prev

		prev = current
		current = next
	}
	l.head = prev
}

func (l *Dlist) ReverseRecursive() {
	if l.len <= 1 {
		return
	}

	var rec func(prev *Dnode, current *Dnode) *Dnode
	rec = func(prev *Dnode, current *Dnode) *Dnode {
		if current == nil {
			return prev
		}
		next := current.next
		current.prev = current.next
		current.next = prev

		return rec(current, next)
	}

	l.head = rec(nil, l.head)
}

func (l *Dlist) BubleSort() {
	if l.len <= 1 {
		return
	}

	var sort func(prev *Dnode, current *Dnode, index int)
	sort = func(prev *Dnode, current *Dnode, index int) {
		if current == nil || index < 1 {
			return
		}

		if prev.data.(int) > current.data.(int) {
			prev.data, current.data = current.data, prev.data
		}

		index--
		sort(current, current.next, index)
	}

	i := l.len
	for i < 2 {
		index := i
		sort(l.head, l.head.next, index-1)
		i--
	}
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
