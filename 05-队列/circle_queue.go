package queue

// 循环队列 (就是一个循环数组)
type circleQueue struct {
	elements []any
	size     int
	front    int
}

const DEFAULT_CAPACITY = 10

func NewCircleQueue() *circleQueue {
	return &circleQueue{
		elements: make([]any, DEFAULT_CAPACITY),
	}
}

func (q *circleQueue) Size() int {
	return q.size
}

func (q *circleQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *circleQueue) Clear() {
	for i := 0; i < q.size; i++ {
		q.elements[q.index(i)] = nil
	}
	q.size = 0
	q.front = 0
}

// 计算真实索引
func (q *circleQueue) index(index int) int {
	index += q.front
	if index < len(q.elements) {
		return index
	}
	// 仅适用于index<2*elements.length
	return index - len(q.elements)
}

// EnQueue 尾部入队
func (q *circleQueue) EnQueue(element any) {
	q.ensureCapacity(q.size + 1)
	q.elements[q.index(q.size)] = element
	q.size++
}

// DeQueue 头部出队
func (q *circleQueue) DeQueue() any {
	oldVal := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = q.index(1)
	q.size--
	return oldVal
}

// Front 获取队列的头元素
func (q *circleQueue) Front() any {
	return q.elements[q.front]
}

// 保证要有capacity的容量
func (q *circleQueue) ensureCapacity(capacity int) {
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
