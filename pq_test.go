package pq

import (
	"testing"
)

type CustomItem struct {
	Key string
	Val int
}

func TestPQCreationInt(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ Creation failed")
	}
}

func TestPQPutInt(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ Creation failed")
	}

	q.Put(5)
	q.Put(3)
	q.Put(1)
	q.Put(1)

	if len(q.Items) != 4 {
		t.Error("PQ PutInt failed")
	}

}

func TestPQGetInt(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ Creation failed")
	}

	q.Put(5)
	q.Put(3)

	if q.Get() != 3 {
		t.Error("PQ GetInt failed")
	}
}

func TestPQPeekInt(t *testing.T) {
	q := CreatePQ[int](func(a, b int) bool { return a < b })
	if q == nil {
		t.Error("PQ Creation failed")
	}

	q.Put(5)
	q.Put(3)

	if q.Peek() != 3 {
		t.Error("PQ Peek failed")
	}
}

func TestPQCreationCustom(t *testing.T) {
	q := CreatePQ[CustomItem](func(a, b CustomItem) bool { return a.Val < b.Val })
	if q == nil {
		t.Error("PQ Creation failed")
	}
}

func TestPQPutCustom(t *testing.T) {
	q := CreatePQ[CustomItem](func(a, b CustomItem) bool { return a.Val < b.Val })
	if q == nil {
		t.Error("PQ Creation failed")
	}

	q.Put(CustomItem{"a", 5})
	q.Put(CustomItem{"b", 3})
	q.Put(CustomItem{"c", 1})

	if len(q.Items) != 3 {
		t.Error("PQ PutCustom failed")
	}
}

func TestPQGetCustom(t *testing.T) {
	q := CreatePQ[CustomItem](func(a, b CustomItem) bool { return a.Val < b.Val })
	if q == nil {
		t.Error("PQ Creation failed")
	}

	q.Put(CustomItem{"a", 5})
	q.Put(CustomItem{"b", 3})
	q.Put(CustomItem{"c", 1})

	head := q.Get()
	if head.Key != "c" || head.Val != 1 {
		t.Error("PQ GetCustom failed")
	}
}

func TestPQPeekCustom(t *testing.T) {
	q := CreatePQ[CustomItem](func(a, b CustomItem) bool { return a.Val < b.Val })
	if q == nil {
		t.Error("PQ Creation failed")
	}

	q.Put(CustomItem{"a", 5})
	q.Put(CustomItem{"b", 3})
	q.Put(CustomItem{"c", 1})

	head := q.Peek()
	if head.Key != "c" || head.Val != 1 {
		t.Error("PQ PeekCustom failed")
	}
}
