package sort

import "sort"

// BucketSort 桶排序 [稳定排序、非原地算法]
type BucketSort struct {
	BaseSort
}

func (s BucketSort) Sort() {
	s.bucketSort()
}

/*
*
  - ◼ 执行流程(类型HashMap的结构)
    ① 创建一定数量的桶（比如用数组、链表作为桶）
    ② 按照一定的规则（不同类型的数据，规则不同），将序列中的元素均匀分配到对应的桶
    ③ 分别对每个桶进行单独排序
    ④ 将所有非空桶的元素合并成有序序列
    ◼ 元素在桶中的索引
    元素值 * 元素数量
*/
func (s BucketSort) bucketSort() {
	//base := len(s.Array)
	base := 0.1
	//桶数组
	buckets := make([][]int, len(s.Array))
	for i := 0; i < len(s.Array); i++ {
		bucketIndex := int(float64(s.Array[i]) * base) //将序列中的元素均匀分配到对应的桶,不同数据类型的这里都要取整获取索引
		if buckets[bucketIndex] == nil {
			buckets[bucketIndex] = make([]int, 0)
		}
		buckets[bucketIndex] = append(buckets[bucketIndex], s.Array[i])
	}
	index := 0
	for i := 0; i < len(buckets); i++ {
		if buckets[i] == nil {
			continue
		}
		sort.Slice(buckets[i], func(x, y int) bool { //分别对每个桶进行单独排序
			return buckets[i][x] < buckets[i][y]
		})
		for _, d := range buckets[i] { //将所有非空桶的元素合并成有序序列
			s.Array[index] = d
			index++
		}
	}
}
