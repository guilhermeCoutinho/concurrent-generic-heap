package safeheap

type SafeHeap interface {
	Push(item interface{})        // O(lg n)
	Pop() (interface{}, error)    // O(lg n)
	Peek() (interface{}, error)   // O(1)
	Remove(val interface{}) error // O(n)

	TestIfHeapified() bool
	Print()
}
