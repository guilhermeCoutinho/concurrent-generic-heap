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

**Interesting point**: One interesting thing I noticed, is that some methods are essentially unsafe in the scope of the caller and other methods are unsafe only in the scope of the heap instance. My current design choice was to keep the API simple, so that the user of the safe-heap doesnt need to know what is safe/unsafe in their scope.

For example:
- Size():  The size() function is unsafe for the caller scope because the result of this function will become obsolete as soon as another heap operation happens elsewhere. So the caller of Size() function would be required to lock the heap while its using the size value to do whatever it wants.
- Push():  The push() function is only unsafe within the instance of the heap. The caller will call push and this will lock the caller untill push is resolved internally, but the caller doesnt need to do anything special about it. The caller just has to be aware that push() is a locking operation and it could slow down critical paths or do some syncronization with waitGroups for example.

Im curious to see how other libs threat this, if they just leave that problem to good documentation or enforce the caller to threat these things somehow.

This is how we centralize the execution of the internal operations (keeping a dedicated channel for every operation):
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
