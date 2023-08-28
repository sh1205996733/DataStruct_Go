package sort

// RadixSort 基数排序 [稳定排序、非原地算法]
type RadixSort struct {
	BaseSort
}

func (s RadixSort) Sort() {
	s.radixSort0()
	//s.radixSort()
}

/**
 * 基数排序非常适合用于整数排序（尤其是非负整数），因此本课程只演示对非负整数进行基数排序
 * 执行流程：依次对个位数、十位数、百位数、千位数、万位数...进行排序（从低位到高位）
 * 	个位数、十位数、百位数的取值范围都是固定的0~9，可以使用计数排序对它们进行排序
 */
func (s *RadixSort) radixSort0() {
	// 找出最大值
	max := s.Array[0]
	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
		}
	} // O(n)

	// 个位数: array[i] / 1 % 10 = 3
	// 十位数：array[i] / 10 % 10 = 9
	// 百位数：array[i] / 100 % 10 = 5
	// 千位数：array[i] / 1000 % 10 = ...
	counts := make([]int, 10)
	output := make([]int, len(s.Array))
	for divider := 1; divider <= max; divider *= 10 { //123
		//1 10 100 1000
		s.countSort(divider, counts, output) //对每一位进行计数排序
	}
}

func (s *RadixSort) countSort(divider int, counts, output []int) {
	for i := 0; i < len(counts); i++ {
		counts[i] = 0
	}
	// 统计每个整数出现的次数
	for i := 0; i < len(s.Array); i++ {
		counts[s.Array[i]/divider%10]++ //个位数
	} // O(n)
	// 每个次数累加上其前面的所有次数
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	// 从后往前遍历元素，将它放到有序数组中的合适位置
	for i := len(s.Array) - 1; i >= 0; i-- {
		index := s.Array[i] / divider % 10
		output[counts[index]-1] = s.Array[i]
		counts[index]--
	}
	// 将有序数组赋值到array
	for i := 0; i < len(s.Array); i++ {
		s.Array[i] = output[i]
	}
}

// 基数排序 – 另一种思路的实现
func (s *RadixSort) radixSort() {
	// 找出最大值
	max := s.Array[0]
	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
		}
	} // O(n)
	//桶数组
	buckets := make([][]int, 10)
	for i := range buckets {
		buckets[i] = make([]int, len(buckets))
	}
	//每个桶的元素数量
	bucketSize := make([]int, len(buckets))
	for divider := 1; divider <= max; divider *= 10 { //123
		for i := 0; i < len(s.Array); i++ {
			no := s.Array[i] / divider % 10
			//if buckets[no] == nil {
			//	buckets[no] = make([]int, len(buckets))
			//}
			buckets[no][bucketSize[no]] = s.Array[i]
			bucketSize[no]++
		}
		index := 0
		for i := 0; i < len(buckets); i++ {
			for j := 0; j < bucketSize[i]; j++ {
				s.Array[index] = buckets[i][j]
				index++
			}
			bucketSize[i] = 0
		}
	}
}
