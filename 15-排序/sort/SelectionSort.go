package sort

// SelectionSort 选择排序 （不稳定排序、原地算法）
//
//	 1 从序列中找出最大的那个元素，然后与最末尾的元素交换位置
//			✓ 执行完一轮后，最末尾的那个元素就是最大的元素
//		2 忽略 1 中曾经找到的最大元素，重复执行步骤 1
//	选择排序的交换次数要远远少于冒泡排序，平均性能优于冒泡排序
type SelectionSort struct {
	BaseSort
}

func (s *SelectionSort) Sort() {
	//s.selectionSort1()
	//s.selectionSort2()
	s.selectionSort3()
}

func (s *SelectionSort) selectionSort1() { //1 2 3 4 5 6
	for i := 0; i < len(s.Array)-1; i++ {
		min := i
		for j := i + 1; j < len(s.Array); j++ {
			if s.cmp(j, min) < 0 {
				min = j
			}
		}
		s.swap(i, min)
	}
}

// 优化二 从末尾开始遍历
func (s *SelectionSort) selectionSort2() { //6 5 4 3 2 1
	for end := len(s.Array) - 1; end > 0; end-- {
		for begin := 0; begin < end; begin++ {
			if s.cmp(begin, end) > 0 {
				s.swap(begin, end)
			}
		}
	}
}

// 优化二 从末尾开始遍历 记录最大值index
func (s *SelectionSort) selectionSort3() { //6 5 4 3 2 1
	for end := len(s.Array) - 1; end > 0; end-- {
		max := 0
		for begin := 1; begin <= end; begin++ {
			if s.cmp(begin, max) > 0 {
				max = begin
			}
		}
		s.swap(max, end)
	}
}

// 优化三 使用堆来选择最大值
