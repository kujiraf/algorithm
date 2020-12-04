package hashtable

import (
	"testing"
)

type table struct {
	name   string
	bucket *bucket
	f      func(in *bucket)
	out    string
	outRev string
	len    int
}

var insTest = []table{
	{
		name:   "append once",
		bucket: &bucket{},
		f:      func(in *bucket) { in.insert("k") },
		out:    "[k]",
		outRev: "[k]",
		len:    1,
	},
	{
		name:   "append several data that has defferent key from each other",
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
		name:   "override all datas",
		bucket: &bucket{},
		f: func(in *bucket) {
			in.insert("k")
			in.insert("k")
			in.insert("k")
		},
		out:    "[k]",
		outRev: "[k]",
		len:    1,
	},
	{
		name:   "override not all datas",
		bucket: &bucket{},
		f: func(in *bucket) {
			in.insert("k1")
			in.insert("k2")
			in.insert("k1")
			in.insert("k3")
			in.insert("k2")
		},
		out:    "[k3 k2 k1]",
		outRev: "[k1 k2 k3]",
		len:    3,
	},
}

func TestPrivateInsertSearch(t *testing.T) {
	copied := make([]table, len(insTest))
	copy(copied, insTest)

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
