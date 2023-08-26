package queue

import (
	array "DataStruct_Go/02-动态数组"
	stack "DataStruct_Go/04-栈"
	"DataStruct_Go/utils"
)

// 队列 FIFO 先进先出 基于链表和数组都可以
type queue struct {
	list utils.List
}

func NewQueue() *queue {
	return &queue{
		list: array.NewArrayList(),
		//list: &linked_list.SingleLinkedList{},
	}
}

func (q *queue) Size() int {
	return q.list.Size()
}

func (q *queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *queue) Clear() {
	q.list.Clear()
}

// EnQueue 入队
func (q *queue) EnQueue(element any) {
	q.list.Add(element)
}

// DeQueue 出队
func (q *queue) DeQueue() any {
	return q.list.Remove(0)
}

// Front 获取队列的头元素
func (q *queue) Front() any {
	return q.list.Get(0)
}

type queueStack struct {
	inStack  *stack.Stack
	outStack *stack.Stack
}

func NewQueueStack() *queueStack {
	return &queueStack{
		inStack:  stack.NewStack(),
		outStack: stack.NewStack(),
	}
}

func (q *queueStack) Size() int {
	return q.inStack.Size() + q.outStack.Size()
}

func (q *queueStack) IsEmpty() bool {
	return q.inStack.IsEmpty() && q.outStack.IsEmpty()
}

func (q *queueStack) Clear() {
	q.inStack.Clear()
	q.outStack.Clear()
}

// EnQueue 入队
func (q *queueStack) EnQueue(element any) {
	q.inStack.Push(element)
}

// DeQueue 出队
func (q *queueStack) DeQueue() any {
	if q.IsEmpty() {
		return nil
	}
	if q.outStack.Size() > 0 {
		return q.outStack.Pop()
	} else {
		for !q.inStack.IsEmpty() {
			q.outStack.Push(q.inStack.Pop())
		}
		return q.outStack.Pop()
	}
}

// Front 获取队列的头元素
func (q *queueStack) Front() any {
	if q.IsEmpty() {
		return nil
	}
	if q.outStack.Size() > 0 {
		return q.outStack.Peek()
	} else {
		for !q.inStack.IsEmpty() {
			q.outStack.Push(q.inStack.Pop())
		}
		return q.outStack.Peek()
	}
}
