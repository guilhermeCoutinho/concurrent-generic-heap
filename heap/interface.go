package heap

type Heap interface {
	Push(item interface{})               // O(lg n)
	Pop() (interface{}, error)           // O(lg n)
	Peek() (interface{}, error)          // O(1)
	Remove(val interface{}) error        // O(n)
	RemoveAt(i int) (interface{}, error) // O(lg n)
	Find(v interface{}) (int, error)     // O(n)
	Update(i int, val interface{})       // O (lg n)
	IsEmpty() bool                       // O(1)
	Size() int                           // O(1)

	TestIfHeapified() bool
	Print()
}

func NewHeap(size int, compareFunc func(a, b interface{}) bool) Heap {
	return &HeapImpl{
		arr:         make([]interface{}, size),
		compareFunc: compareFunc,
		size:        0,
	}
}
