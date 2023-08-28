package sort

// HeapSort 堆排序 [稳定排序、原地算法]
type HeapSort struct {
	heapSize int
	BaseSort
}

func (s HeapSort) Sort() {
	s.heapSort()
}

func (s HeapSort) heapSort() {
	// 原地建堆
	s.heapSize = len(s.Array)
	// 自下而上的下滤 时间复杂度n (从第一个非叶子节点开始)
	for i := (s.heapSize >> 1) - 1; i >= 0; i-- {
		s.siftDown(i)
	}

	for s.heapSize > 1 {
		// 交换堆顶元素和尾部元素
		s.heapSize--
		s.swap(0, s.heapSize)
		// 对0位置进行siftDown（恢复堆的性质）
		s.siftDown(0)
	}
}

func (s HeapSort) siftDown(index int) {
	element := s.Array[index]
	half := s.heapSize >> 1
	for index < half {
		// 默认为左子节点跟它进行比较
		left := (index << 1) + 1 //计算左子节点的索引 2*i +1
		child := s.Array[left]
		// 右子节点
		right := left + 1 //计算右子节点的索引 2*i +2 或者 leftIndex + 1
		// 右子节点比左子节点大
		if right < s.heapSize && s.cmpVal(s.Array[right], child) > 0 {
			child = s.Array[right]
			left = right
		}

		if s.cmpVal(element, child) >= 0 {
			break // 大于等于最大子节点直接break
		}
		s.Array[index] = child
		// 重新赋值index
		index = left
	}
	s.Array[index] = element
}
