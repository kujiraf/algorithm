package tree

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

	b := bst{}
	b.insert(5)
	ifErr("root", b.data, 5)
	b.insert(4)
	ifErr("root.left", b.L.data, 4)
	ifErr("root.left.root", b.L.root.data, 5)
	b.insert(3)
	ifErr("root.left.left", b.L.L.data, 3)
	ifErr("root.left.left.root", b.L.L.root.data, 4)
	b.insert(7)
	ifErr("root.right", b.R.data, 7)
	ifErr("root.right.root", b.R.root.data, 5)
	b.insert(8)
	ifErr("root.right.right", b.R.R.data, 8)
	ifErr("root.right.right.root", b.R.R.root.data, 7)
	b.insert(3)
	ifErr("root.left.left", b.L.L.data, 3)
	ifErr("root.left.left.root", b.L.L.root.data, 4)
	b.insert(78)
	ifErr("root.right.right.right", b.R.R.R.data, 78)
	ifErr("root.right.right.right.root", b.R.R.R.root.data, 8)
	b.insert(56)
	ifErr("root.right.right.right.left", b.R.R.R.L.data, 56)
	ifErr("root.right.right.right.left.root", b.R.R.R.L.root.data, 78)
	b.insert(4)
	b.insert(8)
}

func TestInsertV2(t *testing.T) {
	ifErr := func(tcase string, actual interface{}, expected int) {
		if actual != expected {
			t.Errorf("%s: got %d, want %d", tcase, actual.(int), expected)
		}
	}

	b := bst{}
	b.insertV2(5)
	ifErr("root", b.data, 5)
	b.insertV2(4)
	ifErr("root.left", b.L.data, 4)
	ifErr("root.left.root", b.L.root.data, 5)
	b.insertV2(3)
	ifErr("root.left.left", b.L.L.data, 3)
	ifErr("root.left.left.root", b.L.L.root.data, 4)
	b.insertV2(7)
	ifErr("root.right", b.R.data, 7)
	ifErr("root.right.root", b.R.root.data, 5)
	b.insertV2(8)
	ifErr("root.right.right", b.R.R.data, 8)
	ifErr("root.right.right.root", b.R.R.root.data, 7)
	b.insertV2(3)
	ifErr("root.left.left", b.L.L.data, 3)
	ifErr("root.left.left.root", b.L.L.root.data, 4)
	b.insertV2(78)
	ifErr("root.right.right.right", b.R.R.R.data, 78)
	ifErr("root.right.right.right.root", b.R.R.R.root.data, 8)
	b.insertV2(56)
	ifErr("root.right.right.right.left", b.R.R.R.L.data, 56)
	ifErr("root.right.right.right.left.root", b.R.R.R.L.root.data, 78)
	b.insertV2(4)
	b.insertV2(8)
}

func setupBst() *bst {
	b := bst{}
	b.insertV2(5)
	b.insertV2(4)
	b.insertV2(3)
	b.insertV2(7)
	b.insertV2(8)
	b.insertV2(3)
	b.insertV2(78)
	b.insertV2(56)
	b.insertV2(4)
	b.insertV2(8)
	return &b
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
	b = &bst{}
	b.search(5)
}
