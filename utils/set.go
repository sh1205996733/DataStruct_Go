package utils

type Set interface {
	Size() int            // 元素的数量
	IsEmpty() bool        // 是否为空
	Clear()               // 清除所有元素
	Contains(any) bool    // 是否包含某个元素
	Add(any)              // 添加元素到尾部
	Remove(any)           // 删除元素
	Traversal(...Visitor) // 遍历Set
}
