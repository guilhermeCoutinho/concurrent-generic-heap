package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/guilhermeCoutinho/concurrent-generic-heap/heap"
	"github.com/guilhermeCoutinho/concurrent-generic-heap/safeheap"
)

func main() {
	exampleHeapSafe()
	exampleHeapNotSafe()
}

func exampleHeapSafe() {
	minHeapUnsafe := heap.NewHeap(10, func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	minHeapSafe := safeheap.NewSafeHeap(minHeapUnsafe, 100)

	doRandomConcurrentOperations(minHeapSafe)
	fmt.Println(minHeapSafe.TestIfHeapified())
}

func exampleHeapNotSafe() {
	minHeap := heap.NewHeap(10, func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	maxHeap := heap.NewHeap(10, func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})

	stringHeap := heap.NewHeap(10, func(a, b interface{}) bool {
		return a.(string) > b.(string)
	})

	for i := 0; i < 10; i++ {
		minHeap.Push(rand.Intn(100))
		maxHeap.Push(rand.Intn(100))
		stringHeap.Push(randomWord())
	}

	fmt.Println(minHeap.TestIfHeapified())
	fmt.Println(maxHeap.TestIfHeapified())
	fmt.Println(stringHeap.TestIfHeapified())
}

func doRandomConcurrentOperations(heap safeheap.SafeHeap) {
	var wg sync.WaitGroup
	goRoutines := 100
	wg.Add(goRoutines)
	for i := 0; i < goRoutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				heap.Push(rand.Intn(1000))

				if rand.Intn(100) < 50 {
					heap.Pop()
				}

				randomNumber := rand.Intn(1000)
				heap.Remove(randomNumber)
			}
		}()
	}
	wg.Wait()
}
