package sort

import "math/rand"

// QuickSort 快速排序 (不稳定排序、原地算法)
type QuickSort struct {
	BaseSort
}

func (s *QuickSort) Sort() {
	s.quickSort(0, len(s.Array))
}

// 对 [begin, end) 范围的元素进行快速排序
func (s *QuickSort) quickSort(begin, end int) {
	if end-begin < 2 { // 只有一个元素直接返回
		return
	}
	// 确定轴点位置 O(n)
	pivot := s.pivotIndex(begin, end)
	s.quickSort(begin, pivot)
	s.quickSort(pivot+1, end)
}

/**
 * 构造出 [begin, end) 范围的轴点元素
 *  ① 从序列中选择一个轴点元素（pivot）
 *  	✓ 假设每次选择 0 位置的元素为轴点元素
 *  ② 利用pivot 将序列分割成 2 个子序列
 *  	✓ 将小于 pivot 的元素放在pivot前面（左侧）
 *  	✓ 将大于 pivot的元素放在pivot后面（右侧）
 *  	 ✓ 等于pivot的元素放哪边都可以
 *  ③ 对子序列进行 ① ② 操作
 *  	✓直到不能再分割（子序列中只剩下1个元素）
 * 快速排序的本质逐渐将每一个元素都转换成轴点元素
 * 如果序列中的所有元素都与轴点元素相等，利用目前的算法实现，轴点元素可以将序列分割成 2 个均匀的子序列
 * 如果cmp 位置的判断分别改为 ≤、≥ 轴点元素分割出来的子序列极度不均匀，导致出现最坏时间复杂度 O(n2)
 * @return 轴点元素的最终位置
 */
// 如果序列中的所有元素都与轴点元素相等，利用目前的算法实现，轴点元素可以将序列分割成 2 个均匀的子序列
// 思考：cmp 位置的判断分别改为 ≤、≥ 会起到什么效果？
// 轴点元素分割出来的子序列极度不均匀 导致出现最坏时间复杂度 O(n^2)
func (s *QuickSort) pivotIndex(begin, end int) int {
	// 随机选择一个元素跟begin位置进行交换  (为了降低最坏情况的出现概率，一般采取的做法是 随机选择轴点元素)
	s.swap(begin, rand.Int()%(end-begin)+begin)
	// 备份begin位置的元素
	pivot := s.Array[begin]
	// end指向最后一个元素
	end--
	// 将大于 pivot 的元素放在pivot后面（右侧）
	// 将小于 pivot 的元素放在pivot前面（左侧）
	// 等于pivot的元素放哪边都可以
	for begin < end {
		for begin < end { // 先从end处遍历
			if s.cmpVal(pivot, s.Array[end]) < 0 { // 右边元素 > 轴点元素,end--,继续从end处遍历
				end--
			} else { // 右边元素 <= 轴点元素,将array[end]覆盖到array[begin],begin++,退出循环，开始从begin处遍历
				s.Array[begin] = s.Array[end]
				begin++
				break
			}
		}
		for begin < end { // 从begin处遍历
			if s.cmpVal(s.Array[begin], pivot) < 0 { // 轴点元素 > 左边元素,begin++,继续从begin处遍历
				begin++
			} else { // 轴点元素 <= 左边元素,将array[begin]覆盖到array[end],end--,退出循环，开始从end处遍历
				s.Array[end] = s.Array[begin]
				end--
				break
			}
		}
	}
	// begin == end
	s.Array[begin] = pivot
	return begin
}
