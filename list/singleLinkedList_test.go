package list

import (
	"testing"
)

var list = List{}

func setupEmptyList() {
	list = List{}
}

func setupList() {
	setupEmptyList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("aaa")
	list.Append("bbb")
	list.Append("ccc")
}

func setupIntList(nums ...int) {
	setupEmptyList()
	for _, v := range nums {
		list.Append(v)
	}
}

func TestAppend(t *testing.T) {
	setupList()
	expected := "[1 2 3 aaa bbb ccc]"
	if actual := list.ToString(); actual != expected {
		t.Errorf("list: got %s, want %s", actual, expected)
	}
	if len := list.len; len != 6 {
		t.Errorf("len: got %d, want %d", list.len, 6)
	}
}

func TestPrepend(t *testing.T) {
	setupEmptyList()
	list.Prepend(1)
	list.Append(2)
	list.Append(3)
	list.Prepend("aaa")
	list.Append("bbb")
	list.Prepend("ccc")
	expected := "[ccc aaa 1 2 3 bbb]"
	if actual := list.ToString(); actual != expected {
		t.Errorf("got %s, want %s", actual, expected)
	}
	if len := list.len; len != 6 {
		t.Errorf("len: got %d, want %d", list.len, 6)
	}
}

var removeListTest = []struct {
	name    string
	l       List
	arg     interface{}
	outList string
	outLen  int
}{
	{
		name:    "empty list",
		l:       List{},
		arg:     5,
		outList: "[]",
		outLen:  0,
	},
	{
		name:    "not matched",
		l:       func() List { setupList(); return list }(),
		arg:     5,
		outList: "[1 2 3 aaa bbb ccc]",
		outLen:  6,
	},
	{
		name:    "head matched",
		l:       func() List { setupList(); return list }(),
		arg:     1,
		outList: "[2 3 aaa bbb ccc]",
		outLen:  5,
	},
	{
		name:    "around middle data matched",
		l:       func() List { setupList(); return list }(),
		arg:     "aaa",
		outList: "[1 2 3 bbb ccc]",
		outLen:  5,
	},
	{
		name:    "last data matched",
		l:       func() List { setupList(); return list }(),
		arg:     "ccc",
		outList: "[1 2 3 aaa bbb]",
		outLen:  5,
	},
	{
		name:    "single element list",
		l:       List{head: &Node{data: "a"}, len: 1},
		arg:     "a",
		outList: "[]",
		outLen:  0,
	},
}

func TestRemoveFirst(t *testing.T) {
	for _, tt := range removeListTest {
		tt.l.RemoveFirst(tt.arg)
		if tt.l.ToString() != tt.outList {
			t.Errorf("%s[list]: arg %s, got %s, want %s", tt.name, tt.arg, tt.l.ToString(), tt.outList)
		}
		if len := tt.l.len; len != tt.outLen {
			t.Errorf("%s[len]: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

var reverseTest = []struct {
	name    string
	l       List
	outList string
	outLen  int
}{
	{
		name:    "list has several nodes",
		l:       func() List { setupList(); return list }(),
		outList: "[ccc bbb aaa 3 2 1]",
		outLen:  6,
	},
	{
		name:    "list has a node",
		l:       List{head: &Node{data: 1}, len: 1},
		outList: "[1]",
		outLen:  1,
	},
	{
		name:    "empty list",
		l:       List{},
		outList: "[]",
		outLen:  0,
	},
}

func TestReverseIterative(t *testing.T) {
	for _, tt := range reverseTest {
		tt.l.ReverseIterative()
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s[list]: got %s, want %s", tt.name, actual, tt.outList)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s[len]: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

func TestReverseRecursive(t *testing.T) {
	for _, tt := range reverseTest {
		tt.l.ReverseRecursive()
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s[list]: got %s, want %s", tt.name, actual, tt.outList)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s[len]: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

var reverseEvenTest = []struct {
	name    string
	l       List
	outList string
	outLen  int
}{
	{
		name:    "only even",
		l:       func() List { setupIntList(2, 4, 6, 8); return list }(),
		outList: "[8 6 4 2]",
		outLen:  4,
	},
	{
		name:    "an even between odds",
		l:       func() List { setupIntList(1, 3, 2, 5); return list }(),
		outList: "[1 3 2 5]",
		outLen:  4,
	},
	{
		name:    "odds and evens",
		l:       func() List { setupIntList(2, 4, 6, 8, 1, 5, 6, 7, 8, 2); return list }(),
		outList: "[8 6 4 2 1 5 6 7 2 8]",
		outLen:  10,
	},
}

func TestReverseEven(t *testing.T) {
	for _, tt := range reverseEvenTest {
		tt.l.ReverseEven()
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s[list]: got %s, want %s", tt.name, actual, tt.outList)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s[len]: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}
