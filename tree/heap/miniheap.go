package heap

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.minHeapifyUp(len(h.array) - 1)
}

func (h *MinHeap) Extract() (extracted int) {
	extracted = h.array[0]

	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.minHeapifyDown(0)

	return
}

func (h *MinHeap) minHeapifyUp(i int) {
	for h.array[parent(i)] > h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h *MinHeap) minHeapifyDown(i int) {
	len := len(h.array)
	if i >= len {
		return
	}
	li, ri := left(i), right(i)
	if li >= len && ri >= len {
		return
	}

	var childIndex int
	if ri >= len {
		childIndex = li
	} else {
		l, r := h.array[li], h.array[ri]
		if r < l {
			childIndex = ri
		} else {
			childIndex = li
		}
	}

	if h.array[childIndex] < h.array[i] {
		h.swap(i, childIndex)
		h.minHeapifyDown(childIndex)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *MinHeap) swap(i1 int, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
