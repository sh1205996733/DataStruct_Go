package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	//queue := NewQueue()
	queue := NewQueueStack()
	queue.EnQueue(11)
	queue.EnQueue(22)
	queue.EnQueue(33)
	queue.EnQueue(44)

	for !queue.IsEmpty() {
		fmt.Println(queue.DeQueue())
	}

}
func TestDeque(t *testing.T) {
	deque := NewDeque()
	deque.EnQueueFront(11)
	deque.EnQueueFront(22)
	deque.EnQueue(33)
	deque.EnQueue(44)

	/* 尾  44  33   11  22 头 */

	for !deque.IsEmpty() {
		fmt.Println(deque.DeQueueRear())
	}
}
func TestCircleQueue(t *testing.T) {
	queue := NewCircleQueue()
	// 0 1 2 3 4 5 6 7 8 9
	for i := 0; i < 10; i++ {
		queue.EnQueue(i)
	}
	// null null null null null 5 6 7 8 9
	for i := 0; i < 5; i++ {
		queue.DeQueue()
	}
	// 15 16 17 18 19 5 6 7 8 9
	for i := 15; i < 20; i++ {
		queue.EnQueue(i)
	}
	fmt.Println(queue)
	for !queue.IsEmpty() {
		fmt.Println(queue.DeQueue())
	}
}

func TestCircleDeque(t *testing.T) {
	queue := NewCircleDeque()
	// 头5 4 3 2 1  100 101 102 103 104 105 106 8 7 6 尾

	// 头 8 7 6  5 4 3 2 1  100 101 102 103 104 105 106 107 108 109 null null 10 9 尾
	for i := 0; i < 10; i++ {
		queue.EnQueueFront(i + 1)
		queue.EnQueueRear(i + 100)
	}
	// 头 null 7 6  5 4 3 2 1  100 101 102 103 104 105 106 null null null null null null null 尾
	for i := 0; i < 3; i++ {
		queue.DeQueueFront()
		queue.DeQueueRear()
	}

	// 头 11 7 6  5 4 3 2 1  100 101 102 103 104 105 106 null null null null null null 12 尾
	queue.EnQueueFront(11)
	queue.EnQueueFront(12)
	fmt.Println(queue)
	for !queue.IsEmpty() {
		fmt.Println(queue.DeQueueFront())
	}
}
