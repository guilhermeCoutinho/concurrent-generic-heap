# concurrent-generic-heap
This project is study oriented. The goal was to implement a generic data type and thread-safe heap to hopefully understand better the heap data structure as well as solving underlying problems with concurrency.

The final implementation consists of a unsafe heap and a safe heap which is kinda of a wrapper of the unsafe version.

caveat: Although any data type can be used,  all elements pushed to the heap must be of the same type.

## Unsafe heap examples
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

## Heap Visualization

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

## Safe heap
The chosen design was to use channels to issue commands to a centralized goroutine that would call the correct unsafe operation in a safe way.
Its like this centralized goroutine is listening to heap operation requests. To return values from it, inside the request tehres another channel that will serve as a callback for the result of operations.
This is how it works
```
func (s *SafeHeapImpl) startListening() {
	for {
		select {
		case push := <-s.pushChan:
			s.heap.Push(push)

		case req := <-s.popChan:
			val, err := s.heap.Pop()
			req.result <- operationResult{val: val, err: err}

		case req := <-s.peekChan:
			val, err := s.heap.Peek()
			req.result <- operationResult{val: val, err: err}

		case req := <-s.removeChan:
			err := s.heap.Remove(req.val)
			req.result <- operationResult{val: req.val, err: err}
		}
	}
}
```
