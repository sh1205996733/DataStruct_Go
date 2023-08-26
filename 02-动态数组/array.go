package array

import (
	"DataStruct_Go/utils"
	"fmt"
	"strings"
)

// ArrayList 动态数组
type ArrayList struct {
	size     int   //元素的数量
	elements []any //所有的元素
}

const (
	DEFAULT_CAPACITY = 10
)

func NewArrayList(params ...int) utils.List {
	var capition int
	if len(params) != 1 || params[0] < DEFAULT_CAPACITY {
		capition = DEFAULT_CAPACITY
	}
	elements := make([]any, capition)
	return &ArrayList{
		elements: elements,
	}
}

// Size 元素的数量
func (l *ArrayList) Size() int {
	return l.size
}

// IsEmpty 是否为空
func (l *ArrayList) IsEmpty() bool {
	return l.size == 0
}

// rangeCheck 越界检查
func (l *ArrayList) rangeCheck(index int) {
	if index < 0 || index >= l.size {
		panic(fmt.Sprintf("Index:%d, Size:%d", index, l.size))
	}
}

// rangeCheckForAdd 添加越界检查
func (l *ArrayList) rangeCheckForAdd(index int) {
	if index < 0 || index > l.size {
		panic("添加角标越界")
	}
}

// Contains 是否包含某个元素
func (l *ArrayList) Contains(element any) bool {
	return l.IndexOf(element) != utils.ELEMENT_NOT_FOUND
}

// Add 添加元素到最后面
func (l *ArrayList) Add(element any) {
	l.AddAtIndex(l.size, element)
}

// Get 返回index位置对应的元素
func (l *ArrayList) Get(index int) any {
	l.rangeCheck(index)
	return l.elements[index]
}

// Set 设置index位置的元素 原来的元素ֵ
func (l *ArrayList) Set(index int, element any) any {
	l.rangeCheck(index)
	old := l.elements[index]
	l.elements[index] = element
	return old
}

// AddAtIndex 往index位置添加元素
func (l *ArrayList) AddAtIndex(index int, element any) {
	l.rangeCheckForAdd(index)
	l.ensureCapacity(l.size + 1) // 扩容
	//1,2,3,4,5
	for i := l.size; i > index; i-- { //从size开始，从右到左，前一个往后挪
		l.elements[i] = l.elements[i-1]
	}
	l.elements[index] = element
	l.size++
}

// ensureCapacity 保证要有capacity的容量
func (l *ArrayList) ensureCapacity(capacity int) {
	oldCapacity := len(l.elements) //原数组长度
	if capacity <= oldCapacity {
		return
	}
	// 新容量为旧容量的1.5倍
	newCapacity := oldCapacity>>2 + oldCapacity
	newElements := make([]any, newCapacity)
	for i := 0; i < l.size; i++ {
		newElements[i] = l.elements[i]
	}
	l.elements = newElements
	fmt.Println(oldCapacity, "扩容为", newCapacity)
}

// Remove 删除index位置对应的元素
func (l *ArrayList) Remove(index int) any {
	l.rangeCheck(index)
	old := l.elements[index]            // 1,2,3,4,5
	for i := index; i < l.size-1; i++ { //从index开始，从左到右，后一个往前挪，最后一个不用挪，因为size--时候访问不到
		l.elements[i] = l.elements[i+1]
	}
	l.size--
	l.elements[l.size] = nil
	return old
}

// IndexOf 查看元素的索引(返回第一次出现的索引位置)
// LastIndexOf ：返回最后一次出现的索引位置，只需倒序循环即可
func (l *ArrayList) IndexOf(element any) int {
	if element == nil { // 1
		for i := 0; i < l.size; i++ {
			if l.elements[i] == nil {
				return i
			}
		}
	} else {
		for i := 0; i < l.size; i++ {
			if element == l.elements[i] {
				return i // n
			}
		}
	}
	return utils.ELEMENT_NOT_FOUND
}

// Clear 清除所有元素
func (l *ArrayList) Clear() {
	for i := 0; i < l.size; i++ {
		l.elements[i] = nil
	}
	l.size = 0
}

func (l *ArrayList) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("size=%d, [", l.size))
	for i := 0; i < l.size; i++ {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", l.elements[i]))
		//if i != l.size-1 {
		//	str.WriteString(", ")
		//}
	}
	sb.WriteString("]")
	return sb.String()
}
