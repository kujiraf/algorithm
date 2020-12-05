package hashtable

import (
	"fmt"
	"testing"
)

type testTable struct {
	name   string
	bucket *bucket
	f      func(in *bucket)
	in     string
	out    string
	outRev string
	len    int
}

var testPriIns = []testTable{
	{
		name:   "insert an element",
		bucket: &bucket{},
		f:      func(in *bucket) { in.insert("k") },
		out:    "[k]",
		outRev: "[k]",
		len:    1,
	},
	{
		name:   "insert several elements",
		bucket: &bucket{},
		f: func(in *bucket) {
			in.insert("k1")
			in.insert("k2")
			in.insert("k3")
		},
		out:    "[k3 k2 k1]",
		outRev: "[k1 k2 k3]",
		len:    3,
	},
	{
		name:   "some elements are same",
		bucket: &bucket{},
		f: func(in *bucket) {
			in.insert("k1")
			in.insert("k2")
			in.insert("k1")
			in.insert("k3")
			in.insert("k2")
		},
		out:    "[k2 k3 k1 k2 k1]",
		outRev: "[k1 k2 k1 k3 k2]",
		len:    5,
	},
}

func TestPriInsert(t *testing.T) {
	copied := make([]testTable, len(testPriIns))
	copy(copied, testPriIns)

	for _, tt := range copied {
		tt.f(tt.bucket)
		if a := tt.bucket.toString(); a != tt.out {
			t.Errorf("%s: got %s, want %s", tt.name, a, tt.out)
		}
		if rev := tt.bucket.toReverseString(); rev != tt.outRev {
			t.Errorf("%s: got %s, want %s", tt.name, rev, tt.outRev)
		}
		if l := tt.bucket.len; l != tt.len {
			t.Errorf("%s: got %d, want %d", tt.name, l, tt.len)
		}
	}
}

var b *bucket

func setupBucket(keys ...string) {
	b = &bucket{}
	for _, v := range keys {
		b.insert(v)
	}
}

var testPriSearch = []testTable{
	{
		name:   "no elements",
		bucket: func() *bucket { setupBucket(); return b }(),
		in:     "k",
		out:    "",
	},
	{
		name:   "an element matched",
		bucket: func() *bucket { setupBucket("k"); return b }(),
		in:     "k",
		out:    "k",
	},
	{
		name:   "no matches",
		bucket: func() *bucket { setupBucket("k1", "k2"); return b }(),
		in:     "k",
		out:    "",
	},
	{
		name:   "first of the bucket matched",
		bucket: func() *bucket { setupBucket("i", "j", "k"); return b }(),
		in:     "k",
		out:    "k",
	},
	{
		name:   "middle of the bucket matched",
		bucket: func() *bucket { setupBucket("i", "k", "j"); return b }(),
		in:     "k",
		out:    "k",
	},
	{
		name:   "last of the bucket matched",
		bucket: func() *bucket { setupBucket("k", "j", "i"); return b }(),
		in:     "k",
		out:    "k",
	},
}

func TestPriSearch(t *testing.T) {
	for _, tt := range testPriSearch {
		b := tt.bucket
		n, exist := b.search(tt.in)
		if !exist {
			if tt.out != "" {
				t.Errorf("%s: got %v, want %s", tt.name, n, tt.out)
			}
		}
		if exist {
			if a := n.key; a != tt.out {
				t.Errorf("%s: got %s, want %s", tt.name, a, tt.out)
			}
		}
	}
}

func TestInsert(t *testing.T) {
	table := New()
	table.Insert("abc")
	table.Insert("b")
	table.Insert("cccccc")
	table.Insert("a")
	table.Insert("aa")
	table.Insert("d ")
	table.Insert("d")
	table.Insert("d")
	table.Insert("eeeee")
	e := "[[b abc] [eeeee] [d] [] [] [aa] [d  a cccccc]]"
	if a := table.ToString(); fmt.Sprint(a) != e {
		t.Errorf("Insert assertion err: got %s, want %s", a, e)
	}
}

// For Insert
// var testdata1 = []table{
// 	{
// 		name:   "append once",
// 		bucket: &bucket{},
// 		f:      func(in *bucket) { in.insert("k") },
// 		out:    "[k]",
// 		outRev: "[k]",
// 		len:    1,
// 	},
// 	{
// 		name:   "append several data that has defferent key from each other",
// 		bucket: &bucket{},
// 		f: func(in *bucket) {
// 			in.insert("k1")
// 			in.insert("k2")
// 			in.insert("k3")
// 		},
// 		out:    "[k1 k2 k3]",
// 		outRev: "[k3 k2 k1]",
// 		len:    3,
// 	},
// 	{
// 		name:   "override all datas",
// 		bucket: &bucket{},
// 		f: func(in *bucket) {
// 			in.insert("k")
// 			in.insert("k")
// 			in.insert("k")
// 		},
// 		out:    "[k]",
// 		outRev: "[k]",
// 		len:    1,
// 	},
// 	{
// 		name:   "override not all datas",
// 		bucket: &bucket{},
// 		f: func(in *bucket) {
// 			in.insert("k1")
// 			in.insert("k2")
// 			in.insert("k1")
// 			in.insert("k3")
// 			in.insert("k2")
// 		},
// 		out:    "[k1 k2 k3]",
// 		outRev: "[k3 k2 k1]",
// 		len:    3,
// 	},
// }
