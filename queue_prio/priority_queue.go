package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Item struct {
	Name      string
	priority  int       // The priority of the item in the queue.
	createdAt time.Time // In case items have same priority
}

func NewItem(name string, priority int) *Item {
	return &Item{Name: name, priority: priority, createdAt: time.Now()}
}

type PriorityQueue []*Item

func (a PriorityQueue) Len() int      { return len(a) }
func (a PriorityQueue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PriorityQueue) Less(i, j int) bool {
	if a[i].priority == a[j].priority {
		return a[i].createdAt.Before(a[j].createdAt)
	}
	return a[i].priority > a[j].priority
}
func (a *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*a = append(*a, item)
}
func (a *PriorityQueue) Pop() interface{} {
	n := len(*a)
	item := (*a)[n-1]
	*a = (*a)[0 : n-1]
	return item
}

func main() {
	pq := PriorityQueue{
		NewItem("Vân", 6),
		NewItem("Hòa", 4),
		NewItem("Lán", 9),
	}
	heap.Init(&pq)
	heap.Push(&pq, NewItem("Hương", 6))

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%v\n", item.Name)
	}
}
