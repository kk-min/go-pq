package pq

import "container/heap"

type PQ[T comparable] struct {
	Items []*T
	less  LessFunction[T]
}

type LessFunction[T comparable] func(T, T) bool

func CreatePQ[T comparable](less LessFunction[T]) *PQ[T] {
	return &PQ[T]{[]*T{}, less}
}

func (pq *PQ[T]) Len() int {
	return len(pq.Items)
}

func (pq *PQ[T]) Swap(i, j int) {
	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
}

func (pq *PQ[T]) Less(i, j int) bool {
	return pq.less(*pq.Items[i], *pq.Items[j])
}

func (pq *PQ[T]) Push(item any) {
	pq.Items = append(pq.Items, item.(*T))
}

func (pq *PQ[T]) Put(item T) {
	heap.Push(pq, &item)
}

func (pq *PQ[T]) Pop() any {
	old := pq.Items
	n := len(old)
	item := old[n-1]
	pq.Items = old[0 : n-1]
	return item
}

func (pq *PQ[T]) Get() T {
	return *heap.Pop(pq).(*T)
}

func (pq *PQ[T]) Peek() T {
	return *pq.Items[0]
}
