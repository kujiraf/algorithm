package tree

import "fmt"

var isDebug bool

type bst struct {
	L    *bst
	R    *bst
	data interface{}
}

func (b *bst) insert(d int) {
	if b.data == nil {
		b.data = d
		return
	}

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

	if b.data == nil {
		b.data = d
		return
	}

	var ins func(b *bst, d int) *bst
	ins = func(b *bst, d int) *bst {
		if b == nil {
			b = &bst{data: d}
			return b
		}

		if b.data.(int) > d {
			b.L = ins(b.L, d)
			return b
		}

		if b.data.(int) < d {
			b.R = ins(b.R, d)
			return b
		}

		return b
	}

	b = ins(b, d)
}

func (b *bst) inorder() []int {
	if b == nil || b.data == nil {
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
	if b.data == nil {
		return false
	}

	var srch func(node *bst, d int) bool
	srch = func(node *bst, d int) bool {

		if node == nil {
			return false
		}

		if node.data.(int) > d {
			return srch(node.L, d)
		}

		if node.data.(int) < d {
			return srch(node.R, d)
		}

		return true
	}

	return srch(b, d)
}

func (b *bst) remove(d int) {
	if b.data == nil {
		return
	}

	var rm func(tgt *bst, d int) *bst
	rm = func(tgt *bst, d int) *bst {

		if tgt == nil {
			return tgt
		}
		if tgt.data.(int) > d {
			return rm(tgt.L, d)
		}
		if tgt.data.(int) < d {
			return rm(tgt.R, d)
		}

		// tgt.Rがnilの場合、tgt.Lをそのまま持ってくる
		if tgt.R == nil {
			fmt.Println(tgt.inorder(), tgt.L.inorder())
			return tgt.L
		}

		// tgtのL,Rを一時保存しておく
		b.debugf("tgt :%d\n", tgt.inorder())
		tgtL := tgt.L
		b.debugf("tgtL:%d\n", tgtL.inorder())
		tgtR := tgt.R
		b.debugf("tgtR:%d\n", tgtR.inorder())
		// tgt.Rをtgtにシフトする
		tgt = tgtR
		// tgt.R.Lは、tgt.L.R.R.R...の終点にぶら下げる(tgt.R.L>tgt.R....が確定しているため)
		tgt.L = tgtL
		b.debugf("tgt :%d\n", tgt.inorder())
		if tgtR != nil && tgtR.L == nil {
			return tgt
		}

		right := tgtL.R
		for right != nil {
			right = right.R
		}
		right = tgtR.L
		return tgt
	}

	*b = *rm(b, d)
}

func (b *bst) debugf(format string, a ...interface{}) {
	if isDebug {
		fmt.Print("[DEBUG] ", fmt.Sprintf(format, a...))
	}
}
