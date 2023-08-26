package queue

import (
	"fmt"
	"strings"
)

// 循环双端队列 (就是一个循环数组)
type circleDeque struct {
	elements []any
	size     int
	front    int
}

func NewCircleDeque() *circleDeque {
	return &circleDeque{
		elements: make([]any, DEFAULT_CAPACITY),
	}
}

func (q *circleDeque) Size() int {
	return q.size
}

func (q *circleDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *circleDeque) Clear() {
	for i := 0; i < q.size; i++ {
		q.elements[q.index(i)] = nil
	}
	q.size = 0
	q.front = 0
}

// 计算真实索引
func (q *circleDeque) index(index int) int {
	index += q.front
	if index < 0 {
		return index + len(q.elements)
	}
	if index < len(q.elements) {
		return index
	}
	// 仅适用于index<2*elements.length
	return index - len(q.elements)
}

// EnQueueFront 头部入队
func (q *circleDeque) EnQueueFront(element any) {
	q.ensureCapacity(q.size + 1)
	q.front = q.index(-1)
	q.elements[q.front] = element
	q.size++
}

// EnQueueRear 尾部入队
func (q *circleDeque) EnQueueRear(element any) {
	q.ensureCapacity(q.size + 1)
	q.elements[q.index(q.size)] = element
	q.size++
}

// DeQueueFront 头部出队
func (q *circleDeque) DeQueueFront() any {
	oldVal := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = q.index(1)
	q.size--
	return oldVal
}

// DeQueueRear 尾部出队
func (q *circleDeque) DeQueueRear() any {
	rear := q.index(q.size - 1)
	oldVal := rear
	q.elements[rear] = nil
	q.size--
	return oldVal
}

// Front 获取队列的头元素
func (q *circleDeque) Front() any {
	return q.elements[q.front]
}

// Rear 获取队列的尾元素
func (q *circleDeque) Rear() any {
	return q.elements[q.index(q.size-1)]
}

// 保证要有capacity的容量
func (q *circleDeque) ensureCapacity(capacity int) {
	oldCapacity := len(q.elements)
	if capacity <= oldCapacity {
		return
	}
	// 新容量为旧容量的1.5倍
	newCapacity := (oldCapacity >> 1) + oldCapacity
	newElements := make([]any, newCapacity)
	for i := 0; i < q.size; i++ {
		newElements[i] = q.elements[q.index(i)]
	}
	q.elements = newElements
	// 重置front
	q.front = 0
}

func (q *circleDeque) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("capcacity=%d  size=%d front=%d, [", len(q.elements), q.size, q.front))
	for i := 0; i < len(q.elements); i++ {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", q.elements[i]))
	}
	sb.WriteString("]")
	return sb.String()
}
