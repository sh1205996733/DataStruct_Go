package utils

type List interface {
	Get(int) any         // 获取index位置的元素
	Add(any)             // 添加元素到尾部
	AddAtIndex(int, any) // 在index位置插入一个元素
	Set(int, any) any    // 设置index位置的元素
	Remove(int) any      // 删除index位置的元素
	Contains(any) bool   // 是否包含某个元素
	IndexOf(any) int     // 查看元素的索引
	IsEmpty() bool       // 是否为空
	Clear()              // 清除所有元素
	Size() int           // 元素的数量
}

const ELEMENT_NOT_FOUND = -1
