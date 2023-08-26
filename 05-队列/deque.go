package queue

import (
	array "DataStruct_Go/02-动态数组"
	"DataStruct_Go/utils"
)

// 双端队列 基于链表和数组都可以
type deque struct {
	list utils.List
}

func NewDeque() *deque {
	return &deque{
		list: array.NewArrayList(),
		//list: &linked_list.SingleLinkedList{},
	}
}

func (q *deque) Size() int {
	return q.list.Size()
}

func (q *deque) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *deque) Clear() {
	q.list.Clear()
}

// EnQueueFront 头部入队
func (q *deque) EnQueueFront(element any) {
	q.list.AddAtIndex(0, element)
}

// EnQueue 尾部入队
func (q *deque) EnQueue(element any) {
	q.list.Add(element)
}

// DeQueueFront 头部出队
func (q *deque) DeQueueFront() any {
	return q.list.Remove(0)
}

// DeQueueRear 尾部出队
func (q *deque) DeQueueRear() any {
	return q.list.Remove(q.list.Size() - 1)
}

// Front 获取队列的头元素
func (q *deque) Front() any {
	return q.list.Get(0)
}

// Rear 获取队列的尾元素
func (q *deque) Rear() any {
	return q.list.Get(q.list.Size() - 1)
}
