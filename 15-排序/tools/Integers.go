package tools

import (
	"math/rand"
)

// Random 从min到max随机生成count个数
func Random(count, min, max int) []int {
	if count <= 0 || min > max {
		return nil
	}
	array := make([]int, count)
	delta := max - min + 1
	//rand.Seed(time.Now().UnixMilli())
	for i := 0; i < count; i++ {
		array[i] = min + rand.Intn(delta)
	}
	return array
}

// Combine 合并数组array1和array2，并返回一个新的数组
func Combine(array1, array2 []int) []int {
	if array1 == nil || array2 == nil {
		return nil
	}
	return append(array1, array2...)
}

// Same 返回一个长度为count的数组，前unsameCount位是unsameCount到1，后面都是unsameCount+1
func Same(count, unsameCount int) []int {
	if count <= 0 || unsameCount > count {
		return nil
	}
	array := make([]int, count)
	for i := 0; i < unsameCount; i++ {
		array[i] = unsameCount - i
	}
	for i := unsameCount; i < count; i++ {
		array[i] = unsameCount + 1
	}
	return array
}

// HeadTailAscOrder 返回一个长度为count的数组且头尾升序，中间disorderCount位无序
func HeadTailAscOrder(min, max, disorderCount int) []int {
	array := AscOrder(min, max)
	if disorderCount > len(array) {
		return array
	}

	begin := (len(array) - disorderCount) >> 1
	reverse(array, begin, begin+disorderCount)
	return array
}

// CenterAscOrder 从min到max中返回一个长度为max - min + 1的数组且头尾无序，中间disorderCount位升序
func CenterAscOrder(min, max, disorderCount int) []int {
	array := AscOrder(min, max)
	if disorderCount > len(array) {
		return array
	}
	left := disorderCount >> 1
	reverse(array, 0, left)

	right := disorderCount - left
	reverse(array, len(array)-right, len(array))
	return array
}

// HeadAscOrder 从min到max中返回一个长度为max - min + 1的数组且头部升序，后disorderCount位无序
func HeadAscOrder(min, max, disorderCount int) []int {
	array := AscOrder(min, max)
	if disorderCount > len(array) {
		return array
	}
	reverse(array, len(array)-disorderCount, len(array))
	return array
}

// TailAscOrder 从min到max中返回一个长度为max - min + 1的数组且尾部升序，前disorderCount位无序
func TailAscOrder(min, max, disorderCount int) []int {
	array := AscOrder(min, max)
	if disorderCount > len(array) {
		return array
	}
	reverse(array, 0, disorderCount)
	return array
}

// AscOrder 从min到max中返回一个长度为max - min + 1的升序数组
func AscOrder(min, max int) []int {
	if min > max {
		return nil
	}
	array := make([]int, max-min+1)
	for i := 0; i < len(array); i++ {
		array[i] = min
		min++
	}
	return array
}

// DescOrder 从min到max中返回一个长度为max - min + 1的降序数组
func DescOrder(min, max int) []int {
	if min > max {
		return nil
	}
	array := make([]int, max-min+1)
	for i := 0; i < len(array); i++ {
		array[i] = max
		max++
	}
	return array
}

// reverse 反转一个数组，索引范围是[begin, end)
func reverse(array []int, begin, end int) {
	count := (end - begin) >> 1
	sum := begin + end - 1
	for i := begin; i < begin+count; i++ {
		j := sum - i
		array[i], array[j] = array[j], array[i]
	}
}

// Copy 复制数组
func Copy(array []int) []int {
	tmp := make([]int, len(array))
	copy(tmp, array)
	return tmp
}

// IsAscOrder 判断数组是否是升序
func IsAscOrder(array []int) bool {
	if array == nil || len(array) == 0 {
		return false
	}
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}
