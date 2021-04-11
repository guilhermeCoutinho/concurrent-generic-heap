# generic-heap
This project is study oriented. The end result is a generic implementation for binary heap with a nice to use interface. 

All elements must be of the same type, even though it could be any.

## Examples
3 examples of different heaps and how they are instantiated:

CommonCode
```
  preallocatedSize := 10
```

-   MinHeap:
```
	minHeap := heap.NewHeap(preallocatedSize, func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
```

-   MaxHeap:
```
	maxHeap := heap.NewHeap(preallocatedSize, func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})
```

- Lexicographic heap:
```
	stringHeap := heap.NewHeap(preallocatedSize, func(a, b interface{}) bool {
		return a.(string) > b.(string)
	})
```

## Visualization

The implementation also contains a nice visualization of the heap. For example, after 10 random words pushed to a lexicographic Heap:
```
 understood
├───  turkey
│    ├───  repair
│    │    ├───  brush
│    │    ├───  loving
│    ├───  homeless
│    │    ├───  frantic
├───  thin
│    ├───  actually
│    ├───  stroke
```
