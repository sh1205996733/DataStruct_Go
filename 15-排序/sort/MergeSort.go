package sort

// MergeSort 归并排序 (稳定排序、非原地算法)
type MergeSort struct {
	BaseSort
}

func (s *MergeSort) Sort() {
	s.mergeSort()
}

/**
 * 执行流程
 * ① 不断地将当前序列平均分割成2个子序列 ✓ 直到不能再分割（序列中只剩1个元素）
 * ② 不断地将2个子序列合并成一个有序序列 ✓ 直到最终只剩下1个有序序列
 */
func (s *MergeSort) mergeSort() {
	s.divide(0, len(s.Array))
}

// 对 [begin, end) 范围的数据进行归并排序
func (s *MergeSort) divide(begin, end int) {
	//不断地将当前序列平均分割成2个子序列 ✓ 直到不能再分割（序列中只剩1个元素）
	if end-begin < 2 { //只剩1个元素
		return
	}
	mid := (end + begin) >> 1
	s.divide(begin, mid)
	s.divide(mid, end)
	//不断地将2个子序列合并成一个有序序列 ✓ 直到最终只剩下1个有序序列
	s.merge(begin, mid, end)
}

// 将 [begin, mid) 和 [mid, end) 范围的序列合并成一个有序序列
func (s *MergeSort) merge(begin, mid, end int) {
	li, le := 0, mid-begin
	ri, re := mid, end
	ai := begin
	leftArray := make([]int, le)
	// 申请额外空间 备份左边数组
	for i := li; i < le; i++ {
		leftArray[i] = s.Array[begin+i]
	}
	for li < le { // 如果左边还没有结束,左边结束直接返回
		//cmp(array[ri], leftArray[li]) <= 0 改为<=就失去稳定性
		if ri < re && s.cmpVal(s.Array[ri], leftArray[li]) < 0 { // 如果右边没有结束并且array[ri]<leftArray[li],右边先结束执行else
			s.Array[ai] = s.Array[ri]
			ri++
		} else {
			s.Array[ai] = leftArray[li]
			li++
		}
		ai++
	}
}
