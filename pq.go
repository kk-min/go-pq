package pq

import "container/heap"

// Internally maintains a slice of pointers to items of type T.
type PQ[T comparable] struct {
	Items []*T
	less  LessFunction[T]
}

// Generic comparator function that is passed as a parameter when creating a new PQ.
type LessFunction[T comparable] func(T, T) bool

// Creates a new instance of a priority queue with the given comparator function.
func CreatePQ[T comparable](less LessFunction[T]) *PQ[T] {
	return &PQ[T]{[]*T{}, less}
}

// Returns the size of the priority queue.
func (pq *PQ[T]) Len() int {
	return len(pq.Items)
}

// Adds an item to the priority queue.
func (pq *PQ[T]) Put(item T) {
	heap.Push(pq, &item)
}

// Removes and returns the top item from the priority queue.
func (pq *PQ[T]) Get() T {
	return *heap.Pop(pq).(*T)
}

// Returns the top item from the priority queue without removing it.
func (pq *PQ[T]) Peek() T {
	return *pq.Items[0]
}

// For implementing the heap interface; **do not use**.
func (pq *PQ[T]) Swap(i, j int) {
	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
}

// For implementing the heap interface; **do not use**.
func (pq *PQ[T]) Less(i, j int) bool {
	return pq.less(*pq.Items[i], *pq.Items[j])
}

// For implementing the heap interface; **do not use**.
func (pq *PQ[T]) Push(item any) {
	pq.Items = append(pq.Items, item.(*T))
}

// For implementing the heap interface; **do not use**.
func (pq *PQ[T]) Pop() any {
	old := pq.Items
	n := len(old)
	item := old[n-1]
	pq.Items = old[0 : n-1]
	return item
}
