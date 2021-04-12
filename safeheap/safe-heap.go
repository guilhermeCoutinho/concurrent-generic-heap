package safeheap

import (
	"generic-heap/heap"
)

type heapOperation struct {
	val    interface{}
	result chan operationResult
}

type operationResult struct {
	val interface{}
	err error
}

type SafeHeapImpl struct {
	heap     heap.Heap
	pushChan chan interface{}

	popChan    chan heapOperation
	peekChan   chan heapOperation
	removeChan chan heapOperation
}

func NewSafeHeap(heap heap.Heap, size int) SafeHeap {
	safeHeap := &SafeHeapImpl{
		heap:     heap,
		pushChan: make(chan interface{}, size),

		popChan:    make(chan heapOperation, size),
		peekChan:   make(chan heapOperation, size),
		removeChan: make(chan heapOperation, size),
	}
	go safeHeap.startListening()
	return safeHeap
}

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

func (s *SafeHeapImpl) doOperation(val interface{}, chanOp chan heapOperation) operationResult {
	callBack := make(chan operationResult)
	chanOp <- heapOperation{
		val:    val,
		result: callBack,
	}
	return <-callBack
}

func (s *SafeHeapImpl) Push(item interface{}) {
	s.pushChan <- item
}

func (s *SafeHeapImpl) Pop() (interface{}, error) {
	result := s.doOperation(nil, s.popChan)
	return result.val, result.err
}

func (s *SafeHeapImpl) Peek() (interface{}, error) {
	result := s.doOperation(nil, s.peekChan)
	return result.val, result.err
}

func (s *SafeHeapImpl) Remove(item interface{}) error {
	result := s.doOperation(item, s.removeChan)
	return result.err
}

func (s *SafeHeapImpl) Print() {
	s.heap.Print()
}

func (s *SafeHeapImpl) TestIfHeapified() bool {
	return s.heap.TestIfHeapified()
}
