package hashtable

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

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

type HashTable struct {
	array [ArraySize]*bucket
}

// Insert
func (m *HashTable) Insert(k string) {

}

// Search
// Delete

func getHashCode(s string) (int64, error) {
	b := []byte(s)
	str := hex.EncodeToString(b)
	i, err := strconv.ParseInt(str, 16, 32)
	if err != nil {
		return 0, err
	}
	return i % 10, nil
}

// insert
func (b *bucket) insert(k string) {
	if b.head == nil {
		b.head = &node{key: k}
		b.len++
		return
	}

	if b.search(k) != nil {
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
func (b bucket) search(k string) *node {
	if b.head == nil {
		return nil
	}

	current := b.head
	for current != nil {
		if current.key == k {
			return current
		}
		current = current.next
	}
	return nil
}

// delete

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
