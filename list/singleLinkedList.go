package list

import (
	"fmt"
	"log"
)

// Slist represents a single linked list.
// The zero value for Slist is an empty list ready to use.
type Slist struct {
	head *Snode
	len  int
}

// Snode is an node of a linked list.
type Snode struct {
	data interface{}
	next *Snode
}

// Append adds a node to the end of list.
func (l *Slist) Append(d interface{}) {
	if l.head == nil {
		l.head = &Snode{data: d}
		l.len++
		return
	}

	current := l.head
	for current != nil {
		if current.next == nil {
			current.next = &Snode{data: d}
			l.len++
			return
		}
		current = current.next
	}
}

// Prepend adds new node to the top of the list.
func (l *Slist) Prepend(d interface{}) {
	if l.head == nil {
		l.head = &Snode{data: d}
		l.len++
		return
	}

	newHead := &Snode{
		data: d,
		next: l.head,
	}
	l.head = newHead
	l.len++
}

// RemoveFirst removes only the first data from the list if args data is found in it.
func (l *Slist) RemoveFirst(d interface{}) {
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
func (l *Slist) ReverseIterative() {
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
func (l *Slist) ReverseRecursive() {
	if l.len <= 1 {
		return
	}

	prev := l.head
	current := prev.next
	prev.next = nil

	var recF func(prev *Snode, current *Snode) *Snode
	recF = func(prev *Snode, current *Snode) *Snode {
		if current == nil {
			return prev
		}

		next := current.next
		current.next = prev
		return recF(current, next)
	}

	l.head = recF(prev, current)
}

// ReverseEven sorts consecutive even numbers in reverse order.
func (l *Slist) ReverseEven() {
	if l.len <= 1 {
		return
	}

	var isEven func(arg interface{}) bool
	isEven = func(arg interface{}) bool {
		n, ok := arg.(int)
		if !ok {
			return false
		}
		if n%2 != 0 {
			return false
		}
		return true
	}

	var recRev func(head *Snode, prev *Snode) *Snode
	recRev = func(head *Snode, prev *Snode) *Snode {
		if head == nil {
			return nil
		}

		current := head
		// ここの条件が理解できない。なぜcurrentがevenというだけで逆にするの？（current != headの場合を見ればわかる）
		for current != nil && isEven(current.data) {
			next := current.next
			current.next = prev
			prev = current
			current = next
		}

		// current != headの場合(=headが偶数だった場合)、currentはoddまたはnilになっていて、head.nextをここに向ける
		if current != head {
			// ここの処理があることで、奇数に挟まれた偶数の向きを変えても矯正される
			// 例: input[head/current:even -> odd]
			// 上記のforループにて[<-head:even current:odd]となり
			// head.next = currentにて[head:even -> current:odd]となる
			// forループの最初に逆向きにしたheadのnextを、ここで最終的に戻しているということ。
			head.next = current
			recRev(current, nil)
			// current != headなので、headはprevに入っている
			return prev
		}
		// current == headの場合(=headが奇数だった場合)、一つノードをシフトして再帰的に呼び出し
		head.next = recRev(head.next, head)
		return head
	}

	l.head = recRev(l.head, nil)
}

// Print prints every list node.
func (l Slist) Print() {
	s := l.listToSlice()
	fmt.Printf("%v\n", s)
}

// ToString returns string of list.
func (l Slist) ToString() string {
	s := l.listToSlice()
	return fmt.Sprint(s)
}

func (l Slist) listToSlice() []interface{} {
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
