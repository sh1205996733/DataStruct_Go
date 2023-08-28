package sort

/**
二分搜索 有序数组 （logN）
假设在 [begin, end) 范围内搜索某个元素 v，mid == (begin + end) / 2
	如果 v < m，去 [begin, mid) 范围内二分搜索
	如果 v > m，去 [mid + 1, end) 范围内二分搜索
	如果 v == m，直接返回 mid
*/

// IndexOf 查找v在有序数组array中的位置
func IndexOf(array []int, v int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	begin, end := 0, len(array)
	for begin < end {
		mid := (begin + end) >> 1
		if v < array[mid] {
			end = mid
		} else if v > array[mid] {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// Search 查找v在有序数组array中待插入位置
// 在元素 v 的插入过程中，可以先二分搜索出合适的插入位置，然后再将元素 v 插入
// 要求二分搜索返回的插入位置:第1个大于 v 的元素位置
//
//	 	0 1 2 3 4 5  6  7
//		2 4 8 8 8 12 14
//		如果 v 是 5，返回 2
//		如果 v 是 1，返回 0
//		如果 v 是 15，返回 7
//		如果 v 是 8，返回 5
//
// 假设在 [begin, end) 范围内搜索某个元素 v，mid == (begin + end) / 2
//
//	如果 v < m，去 [begin, mid) 范围内二分搜索
//	如果 v ≥ m，去 [mid + 1, end) 范围内二分搜索
func Search(array []int, v int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	begin, end := 0, len(array)
	for begin < end {
		mid := (begin + end) >> 1
		if v < array[mid] {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}
