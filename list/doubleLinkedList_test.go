package list

import (
	"testing"
)

var dList Dlist

func setupEmptyDlist() {
	dList = Dlist{}
}

func setupDlistByAppend(args ...interface{}) {
	setupEmptyDlist()
	for _, v := range args {
		dList.Append(v)
	}
}

func setupDlistByPrepend(args ...interface{}) {
	setupEmptyDlist()
	for _, v := range args {
		dList.Append(v)
	}
}

var appendTest = []struct {
	name    string
	l       Dlist
	outList string
	outLen  int
}{
	{
		name:    "an element list",
		l:       func() Dlist { setupDlistByAppend(1); return dList }(),
		outList: "[1]",
		outLen:  1,
	},
	{
		name:    "several elements list",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3); return dList }(),
		outList: "[1 2 3]",
		outLen:  3,
	},
}

func TestDlistAppend(t *testing.T) {
	for _, tt := range appendTest {
		// append済み
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s: got %s, want %s", tt.name, actual, tt.outList)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

var prependTest = []struct {
	name    string
	l       Dlist
	outList string
	outLen  int
}{
	{
		name:    "an element list",
		l:       func() Dlist { setupDlistByPrepend(1); return dList }(),
		outList: "[1]",
		outLen:  1,
	},
	{
		name:    "several elements list",
		l:       func() Dlist { setupDlistByPrepend(1, 2, 3); return dList }(),
		outList: "[1 2 3]",
		outLen:  3,
	},
}

func TestDlistPrepend(t *testing.T) {
	for _, tt := range prependTest {
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s: got %s, want %s", tt.name, actual, tt.outList)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

func TestDlistPrev(t *testing.T) {
	var is func(n Dnode, i int) bool
	is = func(n Dnode, i int) bool {
		if n.data.(int) == i {
			return true
		}
		t.Errorf("node:%d, expected:%d", n.data, i)
		return false
	}

	setupDlistByAppend(1, 2, 3)
	current := dList.head
	is(*current, 1)
	is(*current.next, 2)
	is(*current.next.next, 3)
	is(*current.next.next.prev, 2)
	is(*current.next.next.prev.prev, 1)

	setupDlistByPrepend(1, 2, 3)
	current = dList.head
	is(*current, 1)
	is(*current.next, 2)
	is(*current.next.next, 3)
	is(*current.next.next.prev, 2)
	is(*current.next.next.prev.prev, 1)
}
