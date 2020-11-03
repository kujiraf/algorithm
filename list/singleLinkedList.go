package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head *node
}

func (linkedList) hasNext(n *node) bool {
	if n == nil {
		return false
	}
	if n.next == nil {
		return false
	}
	return true
}

func (l linkedList) printList() {
	if l.head == nil {
		fmt.Print("[]\n")
		return
	}
	toPrint := l.head
	fmt.Print("[")
	for l.hasNext(toPrint) {
		fmt.Printf("%d, ", toPrint.data)
		toPrint = toPrint.next
	}
	// 最後のnodeはnextを持たないので、forループ内で出力されない
	fmt.Printf("%d]\n", toPrint.data)
}

func (l *linkedList) prepend(n *node) {
	if l.head == nil {
		l.head = n
		return
	}

	// 既存のデータをnextに退避
	next := l.head
	// headに新しいパラメータを入れる
	l.head = n
	// 追加した新しいパラメータのnextに既存のデータを入れる
	l.head.next = next
}

func (l *linkedList) appendNums(nums ...int) {
	for i := 0; i < len(nums); i++ {
		l.append(&node{data: nums[i]})
	}
}

func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
		return
	}

	end := l.head
	for l.hasNext(end) {
		end = end.next
	}
	end.next = n
}

func (l *linkedList) remove(target int) {
	if l.head == nil {
		fmt.Printf("no data\n")
		return
	}

	if l.head.data == target {
		fmt.Printf("remove(first data!):%d \n", target)
		l.head = l.head.next
		return
	}

	previousToDelete := l.head
	for l.hasNext(previousToDelete) {
		if previousToDelete.next.data == target {
			fmt.Printf("remove:%d \n", target)
			previousToDelete.next = previousToDelete.next.next
			return
		}
		previousToDelete = previousToDelete.next
	}
	fmt.Printf("no remove target:%d \n", target)
}

func (l *linkedList) reverse() {
	if l.head == nil {
		return
	}

	var prev *node
	current := l.head
	for l.hasNext(current) {
		// nextに現在のcurrent.nextを退避
		next := current.next
		// current.nextをprevに向ける(逆向きにする)
		current.next = prev
		// ここからは次の要素にインデックスを移動する
		// prevにはcurrentを入れる(このcurrentはnextが逆向きになっている)
		prev = current
		// currentにはnextを入れる(退避していたcurrent.next)
		current = next
	}
	// 元の最後の要素なので、current.nextに逆向きにリンクしたチェーンを入れる
	current.next = prev
	// 最後の要素のnextにprevを入れたので、これをもとのheadと入れ替える
	l.head = current
}

func (l *linkedList) recursiveReverse() {
	if l.head == nil || l.head.next == nil {
		return
	}
	l.recursive(l.head, nil)
}

func (l *linkedList) recursive(current *node, prev *node) {
	if !l.hasNext(current) {
		current.next = prev
		l.head = current
		return
	}
	next := current.next
	current.next = prev
	prev = current
	current = next
	l.recursive(current, prev)
}

/*
	連続する偶数だけリバースする関数

	Example

	1, 4, 6, 8, 9 => 1, 8, 6, 4, 9

	1, 4, 6, 8, 9, 1, 4, 6, 8, 9 => 1, 8, 6, 4, 9, 1, 8, 6, 4, 9

	1, 2, 3, 4, 5, 6 => 1, 2, 3, 4, 5, 6

	1, 3, 5 => 1, 3, 5
*/
func (l *linkedList) reverseEven() {
	if l.head == nil {
		return
	}

	current := l.head
	l.head = recursiveReverseEven(current, nil)
}

func recursiveReverseEven(head *node, prev *node) *node {
	if head == nil {
		return nil
	}

	current := head
	for current != nil && current.data%2 == 0 {
		next := current.next
		current.next = prev
		prev, current = current, next
	}

	if current != head {
		// 入れ替えが発生したのでcurrentとheadが異なる
		head.next = current
		recursiveReverseEven(current, nil)
		return prev
	}
	head.next = recursiveReverseEven(head.next, head)
	return head
}

func main() {
	mylist := linkedList{}

	// 43 18 48 1 5 9
	mylist.prepend(&node{data: 48})
	mylist.prepend(&node{data: 18})
	mylist.prepend(&node{data: 43})
	mylist.append(&node{data: 1})
	mylist.append(&node{data: 5})
	mylist.append(&node{data: 9})

	// mylist.remove(5)
	// mylist.remove(43)
	// mylist.remove(18)
	// mylist.remove(9)
	// mylist.remove(1)
	// mylist.remove(48)
	// mylist.remove(99)

	// mylist.printList()
	// mylist.reverse()
	// mylist.printList()
	// mylist.recursiveReverse()
	// mylist.printList()

	listForReverseEven := linkedList{}
	listForReverseEven.appendNums(1, 4, 6, 8, 9, 1, 4, 6, 8, 9)
	// listForReverseEven.appendNums(2, 4, 6, 1, 2, 4, 6)
	listForReverseEven.printList()
	listForReverseEven.reverseEven()
	listForReverseEven.printList()
}
