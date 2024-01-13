package main

type MinHeap struct {
	length int
	data   []int
}

func (h *MinHeap) insert(value int) {
	h.data[h.length] = value
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) delete() int {
	if h.length == 0 {
		return -1
	}

	res := h.data[0]
	h.length--

	if h.length == 0 {
		h.data = h.data[:0]
		return res
	}

	h.data[0] = h.data[h.length]

	h.heapifyDown(0)
	return res
}

func (h *MinHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	p := h.parent(index)
	pV := h.data[p]
	v := h.data[index]

	// find min child then compare, then swap
	if pV > v {
		h.data[index], h.data[p] = pV, v
		h.heapifyUp(p)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lIndex := h.leftChild(index)
	rIndex := h.rightChild(index)

	if index >= h.length || lIndex >= h.length {
		return
	}

	lV := h.data[lIndex]
	rV := h.data[rIndex]
	v := h.data[index]

	if lV > rV && v > rV {
		h.data[index], h.data[rIndex] = rV, v
		h.heapifyDown(rIndex)
	} else if rV > lV && v > lV {
		h.data[index], h.data[lIndex] = lV, v
		h.heapifyDown(lIndex)
	}
}

func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChild(index int) int {
	return index*2 + 1
}

func (h *MinHeap) rightChild(index int) int {
	return index*2 + 2
}
