package pq

import "container/heap"

type PQ[T comparable] struct {
	items []*T
	less  LessFunction[T]
}

type LessFunction[T comparable] func(T, T) bool

func CreatePQ[T comparable](less LessFunction[T]) *PQ[T] {
	return &PQ[T]{[]*T{}, less}
}

func (pq *PQ[T]) Len() int {
	return len(pq.items)
}

func (pq *PQ[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PQ[T]) Less(i, j int) bool {
	return pq.less(*pq.items[i], *pq.items[j])
}

func (pq *PQ[T]) Push(item any) {
	pq.items = append(pq.items, item.(*T))
}

func (pq *PQ[T]) Put(item T) {
	heap.Push(pq, &item)
}

func (pq *PQ[T]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}
