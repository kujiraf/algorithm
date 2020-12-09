package heap

import (
	"fmt"
	"testing"
)

var h *MinHeap

func setupMinHeap() {
	h = &MinHeap{}
}

func assert(t *testing.T, expected ...int) {
	e := fmt.Sprint(expected)
	if a := fmt.Sprint(h.array); a != e {
		t.Errorf("got %s, want %s", a, e)
	}
}

func TestInsert(t *testing.T) {

	setupMinHeap()
	h.Insert(5)
	assert(t, 5)
	h.Insert(7)
	h.Insert(3)
	h.Insert(4)
	h.Insert(1)
	assert(t, 1, 3, 5, 7, 4)
}

func TestExtract(t *testing.T) {
	setupMinHeap()
	h.Insert(5)
	h.Insert(8)
	h.Insert(10)
	h.Insert(9)
	h.Insert(20)
	h.Insert(11)
	h.Insert(15)
	assert(t, 5, 8, 10, 9, 20, 11, 15)

	h.Extract()
	assert(t, 8, 9, 10, 15, 20, 11)
	h.Extract()
	assert(t, 9, 11, 10, 15, 20)
	h.Extract()
	assert(t, 10, 11, 20, 15)
	h.Extract()
	assert(t, 11, 15, 20)
	h.Extract()
	assert(t, 15, 20)
	h.Extract()
	assert(t, 20)
	h.Extract()
	assert(t)
}
