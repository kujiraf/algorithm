package tree

import "fmt"

var isDebug bool

type bst struct {
	L    *bst
	R    *bst
	data interface{}
}

func (b *bst) insert(d int) {

	current := b
	for current != nil {
		if current.data.(int) > d {
			if current.L == nil {
				current.L = &bst{data: d}
				return
			}
			current = current.L
		} else if current.data.(int) < d {
			if current.R == nil {
				current.R = &bst{data: d}
				return
			}
			current = current.R
		} else if current.data == d {
			return
		}
	}
}

func (b *bst) insertV2(d int) {

	if b.data.(int) > d {
		if b.L == nil {
			b.L = &bst{data: d}
		} else {
			b.L.insertV2(d)
		}
	}

	if b.data.(int) < d {
		if b.R == nil {
			b.R = &bst{data: d}
		} else {
			b.R.insertV2(d)
		}
	}

}

func (b *bst) inorder() []int {
	if b.data == nil {
		b.debugf("(inorder) no data\n")
		return nil
	}

	ctx := make([]int, 0)

	var appendInt func(node *bst)
	appendInt = func(node *bst) {
		if node == nil {
			return
		}

		appendInt(node.L)
		ctx = append(ctx, node.data.(int))
		appendInt(node.R)
	}

	appendInt(b)
	return ctx
}

func (b *bst) search(d int) bool {

	if b == nil {
		return false
	}

	if b.data.(int) > d {
		return b.L.search(d)
	}

	if b.data.(int) < d {
		return b.R.search(d)
	}

	return true
}

func (b *bst) remove(d int) {
	if b == nil {
		return
	}

}

func (b *bst) debugf(format string, a ...interface{}) {
	if isDebug {
		fmt.Print("[DEBUG] ", fmt.Sprintf(format, a...))
	}
}
