package sort

// CountingSort 计数排序 [稳定排序、非原地算法]
type CountingSort struct {
	BaseSort
}

func (s CountingSort) Sort() {
	//s.countingSort0()
	s.countingSort()
}

func (s CountingSort) countingSort0() {
	/**
	 * 计数排序的核心:统计每个整数在序列中出现的次数，进而推导出每个整数在有序序列中的索引
	 * ◼ 这个版本的实现存在以下问题
	 * 	   无法对负整数进行排序
	 * 极其浪费内存空间
	 *   是个不稳定的排序
	 */
	// 找出最大值
	max := s.Array[0]
	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
		}
	} // O(n)
	// 开辟内存空间，存储每个整数出现的次数
	counts := make([]int, max+1)
	// 统计每个整数出现的次数
	for i := 0; i < len(s.Array); i++ {
		counts[s.Array[i]]++
	} // O(n)
	// 根据整数的出现次数，对整数进行排序
	index := 0
	for i := 0; i < len(counts); i++ {
		for ; counts[i] > 0; counts[i]-- { // 出现n 就设置n次 s.Array[index]
			s.Array[index] = i
			index++
		}
	}
}

// 计数排序 – 改进思路
func (s CountingSort) countingSort() {
	// 找出最大值和最小值
	max, min := s.Array[0], s.Array[0]
	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
		}
		if s.Array[i] < min {
			min = s.Array[i]
		}
	} // O(n)
	// 开辟内存空间，存储每个整数出现的次数
	counts := make([]int, max-min+1)
	// 统计每个整数出现的次数
	for i := 0; i < len(s.Array); i++ {
		counts[s.Array[i]-min]++
	} // O(n)
	// 每个次数累加上其前面的所有次数
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	// 从后往前遍历元素，将它放到有序数组中的合适位置
	output := make([]int, len(s.Array))
	for i := len(s.Array) - 1; i >= 0; i-- {
		output[counts[s.Array[i]-min]-1] = s.Array[i]
		counts[s.Array[i]-min]--
	}
	// 将有序数组赋值到array
	for i := 0; i < len(s.Array); i++ {
		s.Array[i] = output[i]
	}
}
