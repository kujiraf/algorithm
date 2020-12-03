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
	outRev  string
	outLen  int
}{
	{
		name:    "an element list",
		l:       func() Dlist { setupDlistByAppend(1); return dList }(),
		outList: "[1]",
		outRev:  "[1]",
		outLen:  1,
	},
	{
		name:    "several elements list",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3); return dList }(),
		outList: "[1 2 3]",
		outRev:  "[3 2 1]",
		outLen:  3,
	},
}

func TestDlistAppend(t *testing.T) {
	for _, tt := range appendTest {
		// append済み
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s: got %s, want %s", tt.name, actual, tt.outList)
		}
		if rev := tt.l.ToStringReverse(); rev != tt.outRev {
			t.Errorf("%s: got %s, want %s", tt.name, rev, tt.outRev)
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
	outRev  string
	outLen  int
}{
	{
		name:    "an element list",
		l:       func() Dlist { setupDlistByPrepend(1); return dList }(),
		outList: "[1]",
		outRev:  "[1]",
		outLen:  1,
	},
	{
		name:    "several elements list",
		l:       func() Dlist { setupDlistByPrepend(1, 2, 3); return dList }(),
		outList: "[1 2 3]",
		outRev:  "[3 2 1]",
		outLen:  3,
	},
}

func TestDlistPrepend(t *testing.T) {
	for _, tt := range prependTest {
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s: got %s, want %s", tt.name, actual, tt.outList)
		}
		if rev := tt.l.ToStringReverse(); rev != tt.outRev {
			t.Errorf("%s: got %s, want %s", tt.name, rev, tt.outRev)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

var removeTest = []struct {
	name    string
	l       Dlist
	in      interface{}
	outList string
	outRev  string
	outLen  int
}{
	{
		name:    "remove an element",
		l:       func() Dlist { setupDlistByAppend(1); return dList }(),
		in:      1,
		outList: "[]",
		outRev:  "[]",
		outLen:  0,
	},
	{
		name:    "not mutched",
		l:       func() Dlist { setupDlistByAppend(1); return dList }(),
		in:      5,
		outList: "[1]",
		outRev:  "[1]",
		outLen:  1,
	},
	{
		name:    "mutched at the top of the list",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3, 4, 5); return dList }(),
		in:      1,
		outList: "[2 3 4 5]",
		outRev:  "[5 4 3 2]",
		outLen:  4,
	},
	{
		name:    "mutched at the middle of the list",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3, 4, 5); return dList }(),
		in:      3,
		outList: "[1 2 4 5]",
		outRev:  "[5 4 2 1]",
		outLen:  4,
	},
	{
		name:    "mutched several elements",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3, 4, 5, 1, 2); return dList }(),
		in:      1,
		outList: "[2 3 4 5 1 2]",
		outRev:  "[2 1 5 4 3 2]",
		outLen:  6,
	},
	{
		name:    "mutched at the end of the list",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3, 4, 5); return dList }(),
		in:      5,
		outList: "[1 2 3 4]",
		outRev:  "[4 3 2 1]",
		outLen:  4,
	},
}

func TestRemove(t *testing.T) {
	for _, tt := range removeTest {
		tt.l.Remove(tt.in)
		if actual := tt.l.ToString(); actual != tt.outList {
			t.Errorf("%s: got %s, want %s", tt.name, actual, tt.outList)
		}
		if rev := tt.l.ToStringReverse(); rev != tt.outRev {
			t.Errorf("%s: got %s, want %s", tt.name, rev, tt.outRev)
		}
		if tt.l.len != tt.outLen {
			t.Errorf("%s: got %d, want %d", tt.name, tt.l.len, tt.outLen)
		}
	}
}

type revTable struct {
	name    string
	l       Dlist
	outList string
	outRev  string
	outLen  int
}

var reverseTestD = []revTable{
	{
		name:    "an element",
		l:       func() Dlist { setupDlistByAppend(1); return dList }(),
		outList: "[1]",
		outRev:  "[1]",
		outLen:  1,
	},
	{
		name:    "several elements",
		l:       func() Dlist { setupDlistByAppend(1, 2, 3); return dList }(),
		outList: "[3 2 1]",
		outRev:  "[1 2 3]",
		outLen:  3,
	},
	{
		name:    "no element",
		l:       func() Dlist { setupDlistByAppend(); return dList }(),
		outList: "[]",
		outRev:  "[]",
		outLen:  0,
	},
}

func TestReverse(t *testing.T) {

	var test func(table []revTable, f func(tt revTable), execName string)
	test = func(table []revTable, f func(tt revTable), execName string) {
		for _, tt := range table {
			f(tt) // ここだけ差し替えたい
			if actual := tt.l.ToString(); actual != tt.outList {
				t.Errorf("%s(%s): got %s, want %s", tt.name, execName, actual, tt.outList)
			}
			if rev := tt.l.ToStringReverse(); rev != tt.outRev {
				t.Errorf("%s(%s): got %s, want %s", tt.name, execName, rev, tt.outRev)
			}
			if tt.l.len != tt.outLen {
				t.Errorf("%s(%s): got %d, want %d", tt.name, execName, tt.l.len, tt.outLen)
			}
		}
	}

	// 次のテストでも同じスライスを使うため、shallow copyする
	copied := make([]revTable, 0, len(reverseTestD))
	copy(copied, reverseTestD)
	test(copied, func(tt revTable) { tt.l.ReverseIterative() }, "iterative")

	copied = make([]revTable, 0, len(reverseTestD))
	copy(copied, reverseTestD)
	test(copied, func(tt revTable) { tt.l.ReverseRecursive() }, "recursive")
}
