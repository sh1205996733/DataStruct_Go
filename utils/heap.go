package utils

type Heap interface {
	Size() int       // 元素的数量
	IsEmpty() bool   // 是否为空
	Clear()          // 清空
	Add(any)         // 添加元素
	Get() any        // 获得堆顶元素
	Remove() any     // 删除堆顶元素
	Replace(any) any // 删除堆顶元素的同时插入一个新元素
}
