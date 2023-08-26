package linked_list

import (
	"fmt"
)

type BaseList struct {
	size int //元素的数量
}

// Size 元素的数量
func (b *BaseList) Size() int {
	return b.size
}

// IsEmpty 是否为空
func (b *BaseList) IsEmpty() bool {
	return b.size == 0
}

// RangeCheck 越界检查
func (b *BaseList) RangeCheck(index int) {
	if index < 0 || index >= b.size {
		panic(fmt.Sprintf("Index:%d, Size:%d", index, b.size))
	}
}

// RangeCheckForAdd 添加越界检查
func (b *BaseList) RangeCheckForAdd(index int) {
	if index < 0 || index > b.size {
		panic("添加角标越界")
	}
}
