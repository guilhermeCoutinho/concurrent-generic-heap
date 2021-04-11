package main

import (
	"generic-heap/heap"
	"math/rand"
)

func main() {
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

	minHeap.Print()
	maxHeap.Print()
	stringHeap.Print()
}
