package heap

import (
	"DataStruct_Go/utils"
	"fmt"
	"strings"
)

// 二叉堆（最大堆）
type binaryHeap struct {
	size     int
	elements []any
}

const DEFAULT_CAPACITY = 10

func NewBinaryHeap(elements []any) *binaryHeap {
	if elements == nil || len(elements) == 0 {
		return &binaryHeap{
			elements: make([]any, DEFAULT_CAPACITY),
		}
	}
	//不能直接 this.elements = elements,防止其他变量修改，只能复制过去
	size := len(elements)
	newElements := make([]any, utils.Max(DEFAULT_CAPACITY, size))
	for i := 0; i < size; i++ {
		newElements[i] = elements[i]
	}
	binaryHeap := &binaryHeap{
		size:     size,
		elements: newElements,
	}
	binaryHeap.heapify()
	return binaryHeap
}

// 批量建堆
func (h *binaryHeap) heapify() {
	// 自上而下的上滤  时间复杂度nlog(n)
	//for i := 1; i < h.size; i++ {
	//	h.siftUp(i)
	//}
	// 自下而上的下滤 时间复杂度n (从第一个非叶子节点开始)
	for i := (h.size >> 1) - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *binaryHeap) Size() int {
	return h.size
}

func (h *binaryHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *binaryHeap) Clear() {
	for i := 0; i < h.size; i++ {
		h.elements[i] = nil
	}
	h.size = 0
}

func (h *binaryHeap) Add(element any) {
	h.elementNotNullCheck(element)
	h.ensureCapacity(h.size + 1)
	h.elements[h.size] = element
	h.size++
	h.siftUp(h.size - 1)
}

func (h *binaryHeap) ensureCapacity(capacity int) {
	oldCapacity := len(h.elements)
	if capacity <= oldCapacity {
		return
	}
	// 新的容量为原来的1.5倍
	newCapacity := oldCapacity + oldCapacity>>1
	newElements := make([]any, newCapacity)
	for i, element := range h.elements {
		newElements[i] = element
	}
	h.elements = newElements
}

func (h *binaryHeap) Get() any {
	h.emptyCheck()
	return h.elements[0]
}

func (h *binaryHeap) Remove() any {
	h.emptyCheck()
	old := h.elements[0]
	h.elements[0] = h.elements[h.size-1]
	h.elements[h.size-1] = nil
	h.size--
	h.siftDown(0)
	return old
}

func (h *binaryHeap) Replace(element any) any {
	h.elementNotNullCheck(element)
	var old any
	if h.size == 0 { //数组为空，直接新增
		h.elements[0] = element
		h.size++
	} else {
		old = h.elements[0]
		h.elements[0] = element
		h.siftDown(0)
		return old
	}
	return old
}

// 上滤 让index位置的元素上滤(节点和父节点比较，大于父节点则交换位置，节点上去，父节点下来，一直往上比较，直到父节点为空(index==0)或者小于等于父节点)
func (h *binaryHeap) siftUp(index int) {
	element := h.elements[index]
	for index > 0 { //没有父节点
		pid := (index - 1) >> 1 //父节点的索引
		parent := h.elements[pid]
		if utils.Compare(element, parent) <= 0 { // 小于等于父节点直接break
			break
		}
		//大于父节点则交换位置
		// 将父元素存储在index位置
		h.elements[index] = parent
		// 重新赋值index
		index = pid
	}
	h.elements[index] = element
}

// 下滤 	让index位置的元素下滤(节点和最大的子节点比较，小于最大子节点则交换位置，最大子节点上去，节点下来，一直往下比较，直到叶子节点为空或者大于最大子父节点)
func (h *binaryHeap) siftDown(index int) {
	element := h.elements[index]
	half := h.size >> 1
	// 第一个叶子节点的索引 == 非叶子节点的数量
	// index < 第一个叶子节点的索引
	// 必须保证index位置是非叶子节点
	for index < half {
		// index的节点有2种情况
		// 1.只有左子节点
		// 2.同时有左右子节点

		// 默认为左子节点跟它进行比较
		left := (index << 1) + 1 //计算左子节点的索引 2*i +1
		// 右子节点
		right := left + 1                                           //计算右子节点的索引 2*i +2 或者 leftIndex + 1
		if utils.Compare(h.elements[right], h.elements[left]) > 0 { //找出左右子树最大的
			left = right
		}

		child := h.elements[left]

		if utils.Compare(element, child) >= 0 { // 大于等于最大子节点直接break
			break
		}
		//小于最大子节点则交换位置
		// 将最大子节点存储在index位置
		h.elements[index] = child
		// 重新赋值index
		index = left
	}
	h.elements[index] = element
}

func (h *binaryHeap) emptyCheck() {
	if h.size == 0 {
		panic("Heap is empty")
	}
}

func (h *binaryHeap) elementNotNullCheck(element any) {
	if element == nil {
		panic("element must not be null")
	}
}

func (h *binaryHeap) String() string {
	sb := strings.Builder{}
	cnt := h.size
	start, end := 0, 1
	for cnt > 0 {
		if end > h.size {
			end = h.size
		}
		sb.WriteString(fmt.Sprintf("%v\n", h.elements[start:end]))
		count := end - start
		start = end
		end += count * 2
		cnt -= count
	}
	return sb.String()
}
