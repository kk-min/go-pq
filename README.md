# Go Priority Queue 
_Copyright © 2023 kk-min_

Generic priority queue implementation for Go.

## Usage


```go
type CustomItem struct{
	Key string
	Val int
}

// Create a PriorityQueue of type CustomItem by passing in a "less" comparator function to CreatePQ:
q := pq.CreatePQ[CustomItem](func(a, b CustomItem) bool { return a.Val < b.Val })

// Put items in the Priority Queue:
q.Put(CustomItem{"a", 5})
q.Put(CustomItem{"b", 1})

// Peek the head of the queue:
q.Peek() // CustomItem{"b", 1}

// Get (Dequeue) the head of the queue:
q.Get() // CustomItem{"b", 1}
```
