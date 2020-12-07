package tree

type bst struct {
	root *bst
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
				current.L.root = current
				return
			}
			current = current.L
		} else if current.data.(int) < d {
			if current.R == nil {
				current.R = &bst{data: d}
				current.R.root = current
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
			b.L.root = b
			return b
		}

		if b.data.(int) < d {
			b.R = ins(b.R, d)
			b.R.root = b
			return b
		}

		return b
	}

	b = ins(b, d)
}

func (b *bst) inorder() []int {
	if b.data == nil {
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
