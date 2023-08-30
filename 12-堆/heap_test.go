package heap

import (
	"fmt"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap(nil)
	heap.Add(72)
	heap.Add(43)
	heap.Add(50)
	heap.Add(38)
	heap.Add(10)
	heap.Add(90)
	heap.Add(65)
	fmt.Println(heap)
	heap.Remove()
	fmt.Println(heap)

	fmt.Println(heap.Replace(70))
	fmt.Println(heap)

	fmt.Println(heap.Replace(60))
	fmt.Println(heap)
}

func TestHeapify(t *testing.T) {
	data := []any{88, 44, 53, 41, 16, 6, 70, 18, 85, 98, 81, 23, 36, 43, 37}
	heap := NewBinaryHeap(data)
	fmt.Println(heap)

	data[0] = 10
	data[1] = 20
	fmt.Println(heap)
}

func TestMinBinaryHeap(t *testing.T) { //小顶堆
	data := []any{88, 44, 53, 41, 16, 6, 70, 18, 85, 98, 81, 23, 36, 43, 37}
	heap := NewBinaryHeap(data)
	fmt.Println(heap)
}

func TestTopK(t *testing.T) {
	// 新建一个小顶堆
	heap := NewBinaryHeap(nil)

	// 找出最大的前k个数
	k := 3
	data := []int{51, 30, 39, 92, 74, 25, 16, 93,
		91, 19, 54, 47, 73, 62, 76, 63, 35, 18,
		90, 6, 65, 49, 3, 26, 61, 21, 48}
	for i := 0; i < len(data); i++ {
		if heap.Size() < k { // 前k个数添加到小顶堆
			heap.Add(data[i]) // logk
		} else if data[i] > heap.Get().(int) { // 如果是第k + 1个数，并且大于堆顶元素
			heap.Replace(data[i]) // logk
		}
	}
	// O(nlogk)
	fmt.Println(heap)
}
