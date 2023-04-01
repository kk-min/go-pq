package pq

import (
	"reflect"
	"testing"
)

func TestPQCreation(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ creation failed")
	}
}

func TestPQPutInt(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ creation failed")
	}

	q.Put(5)
	q.Put(3)

	temp := []int{}
	for _, v := range q.Items {
		temp = append(temp, *v)
	}

	if !reflect.DeepEqual(temp, []int{3, 5}) {
		t.Error("PQ Put failed: ", q.Items)
	}
}

func TestPQGetInt(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ creation failed")
	}

	q.Put(5)
	q.Put(3)

	if q.Get() != 3 {
		t.Error("PQ Get failed")
	}
}

func TestPQGetPeek(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ creation failed")
	}

	q.Put(5)
	q.Put(3)

	if q.Peek() != 3 {
		t.Error("PQ Get failed")
	}
}
