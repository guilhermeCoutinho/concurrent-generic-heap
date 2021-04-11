package heap

import "fmt"

type HeapImpl struct {
	arr         []interface{}
	compareFunc func(a, b interface{}) bool

	size int
}

func (h *HeapImpl) IsEmpty() bool {
	return h.size == 0
}

func (h *HeapImpl) Size() int {
	return h.size
}

func (h *HeapImpl) Update(i int, val interface{}) {
	h.arr[i] = val
	h.heapify(i)
}

func (h *HeapImpl) Push(item interface{}) {
	if len(h.arr) <= h.size {
		h.arr = append(h.arr, item)
	} else {
		h.arr[h.size] = item
	}

	h.heapifyUp(h.size, false)
	h.size++
}

func (h *HeapImpl) Pop() (interface{}, error) {
	if h.size == 0 {
		return nil, EmptyHeap
	}
	if h.size == 1 {
		h.size--
		return h.arr[0], nil
	}

	root := h.arr[0]
	h.size--
	h.arr[0] = h.arr[h.size]
	h.heapifyDown(0)
	return root, nil
}

func (h *HeapImpl) Peek() (interface{}, error) {
	if h.size == 0 {
		return nil, EmptyHeap
	}
	return h.arr[0], nil
}

func (h *HeapImpl) Find(val interface{}) (int, error) {
	for i := 0; i < h.size; i++ {
		fmt.Println(h.arr[i], val)
		if h.arr[i] == val {
			return i, nil
		}
	}
	return -1, ElementNotFound
}

func (h *HeapImpl) Remove(val interface{}) error {
	i, err := h.Find(val)
	if err != nil {
		return err
	}
	_, err = h.RemoveAt(i)
	return err
}

func (h *HeapImpl) RemoveAt(i int) (interface{}, error) {
	if i >= h.size || i < 0 {
		return nil, IndexOutOfRange
	}
	val := h.arr[i]
	h.heapifyUp(i, true)
	h.Pop()
	return val, nil
}

func (h *HeapImpl) TestIfHeapified() bool {
	for i := 0; i < h.size; i++ {
		l, r := h.childIdx(i)
		if l < h.size && h.compare(l, i) {
			return false
		}

		if r < h.size && h.compare(r, i) {
			return false
		}
	}
	return true
}

func (h *HeapImpl) Print() {
	h.printRecursive(0, "", "")
}

func (h *HeapImpl) printRecursive(i int, prefix, childrenPrefix string) {
	if i >= h.size {
		return
	}
	fmt.Println(prefix, h.arr[i])
	l, r := h.childIdx(i)
	if l < h.size {
		h.printRecursive(l, childrenPrefix+"├─── ", childrenPrefix+"│    ")
	}
	if r < h.size {
		h.printRecursive(r, childrenPrefix+"├─── ", childrenPrefix+"│    ")
	}
}

func (h *HeapImpl) heapify(i int) {
	parent := h.parentIdx(i)
	swapUp := parent >= 0 && h.compare(i, parent)

	if swapUp {
		h.heapifyUp(i, false)
		return
	}

	l, r := h.childIdx(i)
	swapLeft := l < h.size && h.compare(l, i)
	swapRight := r < h.size && h.compare(r, i)

	if swapLeft || swapRight {
		h.heapifyDown(i)
	}
}

func (h *HeapImpl) heapifyUp(i int, ignoreCompare bool) {
	if i == 0 {
		return
	}
	j := h.parentIdx(i)
	if h.compare(i, j) || ignoreCompare {
		h.swap(i, j)
		h.heapifyUp(j, ignoreCompare)
	}
}

func (h *HeapImpl) compare(i, j int) bool {
	return h.compareFunc(h.arr[i], h.arr[j])
}

func (h *HeapImpl) heapifyDown(i int) {
	l, r := h.childIdx(i)
	smallest := i

	if l < h.size && h.compare(l, smallest) {
		smallest = l
	}

	if r < h.size && h.compare(r, smallest) {
		smallest = r
	}

	if i != smallest {
		h.swap(i, smallest)
		h.heapifyDown(smallest)
	}
}

func (h *HeapImpl) swap(i, j int) {
	temp := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = temp
}

func (h *HeapImpl) parentIdx(i int) int {
	return (i - 1) / 2
}

func (h *HeapImpl) childIdx(i int) (left int, right int) {
	left = i*2 + 1
	right = i*2 + 2
	return
}
