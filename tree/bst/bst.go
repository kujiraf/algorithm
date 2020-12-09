package bst

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

	if b.data == d {

	}
	if b.data.(int) > d {
		rm(b, b.L, d)
	}
	if b.data.(int) < d {
		rm(b, b.R, d)
	}
}

func rm(root *bst, node *bst, d int) {
	if node == nil {
		return
	}

	if node.data.(int) > d {
		rm(node, node.L, d)
	}
	if node.data.(int) < d {
		rm(node, node.R, d)
	}

	// 削除対象のノードに、node.R.L.L.L...で右ノードの最小値を設定する
	tmp := mini(node.R, d)
	// 探し出した最小値をノードに設定する
	node.data = tmp.data
	// 持ってきたデータは削除対象なので、削除する
	rm(node, node.R, tmp.data.(int))
}

func mini(node *bst, d int) *bst {
	if node.L == nil {
		return node
	}

	for node.L != nil {
		node = node.L
	}
	return node.L
}

func (b *bst) debugf(format string, a ...interface{}) {
	if isDebug {
		fmt.Print("[DEBUG] ", fmt.Sprintf(format, a...))
	}
}
