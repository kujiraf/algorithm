package list

import (
	"testing"
)

var sList = Slist{}

func setupEmptyList() {
	sList = Slist{}
}

func setupList() {
	setupEmptyList()
	sList.Append(1)
	sList.Append(2)
	sList.Append(3)
	sList.Append("aaa")
	sList.Append("bbb")
	sList.Append("ccc")
}

func setupIntList(nums ...int) {
	setupEmptyList()
	for _, v := range nums {
		sList.Append(v)
	}
}

func TestAppend(t *testing.T) {
	setupList()
	expected := "[1 2 3 aaa bbb ccc]"
	if actual := sList.ToString(); actual != expected {
		t.Errorf("list: got %s, want %s", actual, expected)
	}
	if len := sList.len; len != 6 {
		t.Errorf("len: got %d, want %d", sList.len, 6)
	}
}

func TestPrepend(t *testing.T) {
	setupEmptyList()
	sList.Prepend(1)
	sList.Append(2)
	sList.Append(3)
	sList.Prepend("aaa")
	sList.Append("bbb")
	sList.Prepend("ccc")
	expected := "[ccc aaa 1 2 3 bbb]"
	if actual := sList.ToString(); actual != expected {
		t.Errorf("got %s, want %s", actual, expected)
	}
	if len := sList.len; len != 6 {
		t.Errorf("len: got %d, want %d", sList.len, 6)
	}
}

var removeListTest = []struct {
	name    string
	l       Slist
	arg     interface{}
	outList string
	outLen  int
}{
	{
		name:    "empty list",
		l:       Slist{},
		arg:     5,
		outList: "[]",
		outLen:  0,
	},
	{
		name:    "not matched",
		l:       func() Slist { setupList(); return sList }(),
		arg:     5,
		outList: "[1 2 3 aaa bbb ccc]",
		outLen:  6,
	},
	{
		name:    "head matched",
		l:       func() Slist { setupList(); return sList }(),
		arg:     1,
		outList: "[2 3 aaa bbb ccc]",
		outLen:  5,
	},
	{
		name:    "around middle data matched",
		l:       func() Slist { setupList(); return sList }(),
		arg:     "aaa",
		outList: "[1 2 3 bbb ccc]",
		outLen:  5,
	},
	{
		name:    "last data matched",
		l:       func() Slist { setupList(); return sList }(),
		arg:     "ccc",
		outList: "[1 2 3 aaa bbb]",
		outLen:  5,
	},
	{
		name:    "single element list",
		l:       Slist{head: &Snode{data: "a"}, len: 1},
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
	l       Slist
	outList string
	outLen  int
}{
	{
		name:    "list has several nodes",
		l:       func() Slist { setupList(); return sList }(),
		outList: "[ccc bbb aaa 3 2 1]",
		outLen:  6,
	},
	{
		name:    "list has a node",
		l:       Slist{head: &Snode{data: 1}, len: 1},
		outList: "[1]",
		outLen:  1,
	},
	{
		name:    "empty list",
		l:       Slist{},
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
	l       Slist
	outList string
	outLen  int
}{
	{
		name:    "only even",
		l:       func() Slist { setupIntList(2, 4, 6, 8); return sList }(),
		outList: "[8 6 4 2]",
		outLen:  4,
	},
	{
		name:    "an even between odds",
		l:       func() Slist { setupIntList(1, 3, 2, 5); return sList }(),
		outList: "[1 3 2 5]",
		outLen:  4,
	},
	{
		name:    "odds and evens",
		l:       func() Slist { setupIntList(2, 4, 6, 8, 1, 5, 6, 7, 8, 2); return sList }(),
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
