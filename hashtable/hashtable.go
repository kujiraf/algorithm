package hashtable

import (
	"fmt"
	"log"
)

// ArraySize is HashTable's array size.
const ArraySize = 7

type bucket struct {
	head *node
	len  int
}

type node struct {
	key  string
	prev *node
	next *node
}

// HashTable has buckets with the indexes.
type HashTable struct {
	array [ArraySize]*bucket
}

// New create new HashTable instance.
func New() *HashTable {
	t := &HashTable{}
	for i := range t.array {
		t.array[i] = &bucket{}
	}
	return t
}

// Insert inserts new key. If the key is already exists, then this func does nothing.
func (t *HashTable) Insert(k string) {
	index := t.hash(k)
	b := t.array[index]
	if _, exist := b.search(k); !exist {
		b.insert(k)
	}
}

// Search searchs the node witch has the key. If it not found, return empty str and false.
func (t *HashTable) Search(k string) (string, bool) {
	index := t.hash(k)
	b := t.array[index]
	n, exist := b.search(k)
	if !exist {
		return "", false
	}
	return n.key, true
}

// Delete deletes the node witch has the key.
func (t *HashTable) Delete(k string) {
	index := t.hash(k)
	b := t.array[index]
	b.delete(k)
}

// ToString returns the 2d slice of string. It shows all HashTable elements.
func (t *HashTable) ToString() [][]string {
	str := make([][]string, ArraySize)
	for i, b := range t.array {
		cur := b.head
		j := 0
		str[i] = make([]string, b.len)
		for cur != nil {
			str[i][j] = cur.key
			cur = cur.next
			j++
		}
	}
	return str
}

// func getHashCode(s string) (int64, error) {
// 	b := []byte(s)
// 	str := hex.EncodeToString(b)
// 	i, err := strconv.ParseInt(str, 16, 32)
// 	if err != nil {
// 		return -1, err
// 	}
// 	return i % ArraySize, nil
// }

func (t HashTable) hash(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v)
	}

	return sum % ArraySize
}

// insert
func (b *bucket) insert(k string) {
	if b.head == nil {
		b.head = &node{key: k}
		b.len++
		return
	}

	next := b.head
	b.head = &node{key: k}
	b.head.next = next
	if next != nil {
		next.prev = b.head
	}
	b.len++
}

// search
func (b bucket) search(k string) (*node, bool) {
	if b.head == nil {
		return nil, false
	}

	current := b.head
	for current != nil {
		if current.key == k {
			return current, true
		}
		current = current.next
	}
	return nil, false
}

// delete
func (b *bucket) delete(k string) {
	n, exist := b.search(k)
	if !exist {
		return
	}

	next := n.next
	if n == b.head {
		b.head = next
		if next != nil {
			next.prev = nil
		}
		b.len--
		return
	}

	prev := n.prev
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	b.len--
	n = nil
}

func (b bucket) toString() string {
	if b.len == 0 {
		return ""
	}

	current := b.head
	count := 0
	s := make([]string, 0, b.len)
	for current != nil {
		s = append(s, current.key)
		current = current.next
		count++
		if count > b.len {
			log.Fatalf("to slice error. slice=%v, len=%d\n", s, b.len)
		}
	}
	return fmt.Sprint(s)
}

func (b bucket) toReverseString() string {
	if b.len == 0 {
		return ""
	}

	var prev *node
	current := b.head
	for current != nil {
		prev = current
		current = current.next
	}

	count := 0
	s := make([]string, 0, b.len)
	current = prev
	for current != nil {
		s = append(s, current.key)
		count++
		if count > b.len {
			log.Fatalf("to slice error. slice=%v, len=%d\n", s, b.len)
		}
		current = current.prev
	}
	return fmt.Sprint(s)
}
