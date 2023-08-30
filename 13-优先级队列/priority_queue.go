package priority_queue

import (
	heap "DataStruct_Go/12-堆"
	"DataStruct_Go/utils"
)

type priorityQueue struct {
	heap utils.Heap
}

func NewPriorityQueue() *priorityQueue {
	return &priorityQueue{
		heap: heap.NewBinaryHeap(nil),
	}
}

func (p *priorityQueue) Size() int {
	return p.heap.Size()
}

func (p *priorityQueue) IsEmpty() bool {
	return p.heap.IsEmpty()
}

func (p *priorityQueue) Clear() {
	p.heap.Clear()
}

func (p *priorityQueue) EnQueue(element any) {
	p.heap.Add(element)
}

func (p *priorityQueue) DeQueue() any {
	return p.heap.Remove()
}

func (p *priorityQueue) Front() any {
	return p.heap.Get()
}
