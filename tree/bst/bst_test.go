package bst

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	ifErr := func(tcase string, actual interface{}, expected int) {
		if actual != expected {
			t.Errorf("%s: got %d, want %d", tcase, actual.(int), expected)
		}
	}

	b := &bst{data: 5}
	ifErr("root", b.data, 5)
	b.insert(4)
	ifErr("root.left", b.L.data, 4)
	b.insert(3)
	ifErr("root.left.left", b.L.L.data, 3)
	b.insert(7)
	ifErr("root.right", b.R.data, 7)
	b.insert(8)
	ifErr("root.right.right", b.R.R.data, 8)
	b.insert(3)
	ifErr("root.left.left", b.L.L.data, 3)
	b.insert(78)
	ifErr("root.right.right.right", b.R.R.R.data, 78)
	b.insert(56)
	ifErr("root.right.right.right.left", b.R.R.R.L.data, 56)
	b.insert(4)
	b.insert(8)
}

func TestInsertV2(t *testing.T) {
	ifErr := func(tcase string, actual interface{}, expected int) {
		if actual != expected {
			if actual == nil {
				t.Errorf("%s: got nil, want %d", tcase, expected)
			} else {
				t.Errorf("%s: got %d, want %d", tcase, actual.(int), expected)
			}
		}
	}

	b := &bst{data: 5}
	ifErr("root", b.data, 5)
	b.insertV2(4)
	ifErr("root.left", b.L.data, 4)
	b.insertV2(3)
	ifErr("root.left.left", b.L.L.data, 3)
	b.insertV2(7)
	ifErr("root.right", b.R.data, 7)
	b.insertV2(8)
	ifErr("root.right.right", b.R.R.data, 8)
	b.insertV2(3)
	ifErr("root.left.left", b.L.L.data, 3)
	b.insertV2(78)
	ifErr("root.right.right.right", b.R.R.R.data, 78)
	b.insertV2(56)
	ifErr("root.right.right.right.left", b.R.R.R.L.data, 56)
	b.insertV2(4)
	b.insertV2(8)
}

func setupBst() *bst {
	b := &bst{data: 5}
	b.insertV2(4)
	b.insertV2(3)
	b.insertV2(7)
	b.insertV2(8)
	b.insertV2(3)
	b.insertV2(78)
	b.insertV2(56)
	b.insertV2(4)
	b.insertV2(8)
	return b
}

func TestInorder(t *testing.T) {
	b := setupBst()
	exp := "[3 4 5 7 8 56 78]"
	if act := fmt.Sprint(b.inorder()); act != exp {
		t.Errorf("inorder assersion err: got %s, want %s", act, exp)
	}
}

func TestSearch(t *testing.T) {
	var isTrue func(b *bst, d ...int)
	isTrue = func(b *bst, d ...int) {
		for _, v := range d {
			if !b.search(v) {
				t.Error("assertion error : ", v)
			}
		}
	}

	b := setupBst()
	isTrue(b, 5, 4, 3, 7, 8, 78, 56)
	if b.search(100) {
		t.Error("assertion error : ", 100)
	}
}

func TestRemove(t *testing.T) {

	b := setupBst()
	var is func(exp ...int)
	is = func(exp ...int) {
		act := fmt.Sprint(b.inorder())
		e := fmt.Sprint(exp)
		if act != e {
			t.Errorf("remove assersion err: got %s, want %s", act, e)
		}
	}

	b = &bst{data: 1}
	isDebug = true
	b.insert(1)
	b.remove(1)
	is()

	b.insert(5)
	b.insert(4)
	b.insert(6)
	is(4, 5, 6)

	b.remove(5)
	is(4, 6)

	b.remove(4)
	is(6)

	// ---------------------------------
	// is(3, 4, 5, 7, 8, 56, 78)

	// b.remove(4)
	// is(3, 5, 7, 8, 56, 78)

	// b.remove(7)
	// b.remove(8)
	// is(4, 5, 56, 78)

	// b.remove(78)
	// b.remove(0)
	// is(4, 5, 56)

	// b.remove(4)
	// b.remove(5)
	// b.remove(56)
	// is()
}
