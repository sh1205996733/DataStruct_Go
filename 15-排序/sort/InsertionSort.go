package sort

// InsertionSort 插入排序 (稳定排序、原地算法)
/**
 	◼ 插入排序的时间复杂度与逆序对的数量成正比关系
		逆序对的数量越多，插入排序的时间复杂度越高
	◼ 最坏、平均时间复杂度:O(n2)
	◼ 最好时间复杂度:O(n)
	◼ 空间复杂度:O(1)
	◼ 属于稳定排序
	◼ 当逆序对的数量极少时，插入排序的效率特别高
	甚至速度比 O nlogn 级别的快速排序还要快
	◼ 数据量不是特别大的时候，插入排序的效率也是非常好的
*/
type InsertionSort struct {
	BaseSort
}

func (i *InsertionSort) Sort() {
	//i.insertionSort1()
	//i.insertionSort2()
	i.insertionSort3()
}

/**
 *  插入排序非常类似于扑克牌的排序
 *  执行流程:
 *  ① 在执行过程中，插入排序会将序列分为2部分
 *		✓ 头部是已经排好序的，尾部是待排序的
 *	② 从头开始扫描每一个元素
 *		✓ 每当扫描到一个元素，就将它插入到头部合适的位置，使得头部数据依然保持有序
 */
func (i *InsertionSort) insertionSort1() { //从第一个元素开始，假如第一个是有序的
	for begin := 1; begin < len(i.Array); begin++ {
		cur := begin
		for cur > 0 && i.cmp(cur, cur-1) < 0 {
			i.swap(cur, cur-1)
			cur--
		}
	}
}

/**
 * 减少交换次数
 * 思路是将【交换】转为【挪动】
 *  ① 先将待插入的元素备份
 *  ② 头部有序数据中比待插入元素大的，都朝尾部方向挪动1个位置
 *  ③ 将待插入元素放到最终的合适位置
 */
func (i *InsertionSort) insertionSort2() { //从第一个元素开始，假如第一个是有序的
	for begin := 1; begin < len(i.Array); begin++ {
		cur := begin
		v := i.Array[cur]
		for cur > 0 && i.cmpVal(v, i.Array[cur-1]) < 0 {
			i.Array[cur] = i.Array[cur-1]
			cur--
		}
		i.Array[cur] = v
	}
}

/**
 * 利用二分搜索找到 index 位置元素的待插入位置
 */
func (i *InsertionSort) insertionSort3() { //从第一个元素开始，假如第一个是有序的
	for begin := 1; begin < len(i.Array); begin++ {
		i.insert(begin, i.search(begin))
	}
}

// 将source位置的元素插入到dest位置
func (i *InsertionSort) insert(source, dest int) {
	t := i.Array[source]
	for cur := source; cur > dest; cur-- {
		i.Array[cur] = i.Array[cur-1]
	}
	i.Array[dest] = t
}

// 利用二分搜索找到 index 位置元素的待插入位置
// 已经排好序数组的区间范围是 [0, index) 2 3 4 5  1
func (i *InsertionSort) search(index int) int {
	begin, end := 0, index
	for begin < end {
		mid := (begin + end) >> 1
		if i.cmp(index, mid) < 0 {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}
