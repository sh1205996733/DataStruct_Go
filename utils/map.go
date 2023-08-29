package utils

type Map interface {
	Size() int
	IsEmpty() bool
	Clear()
	Put(any, any) any
	Get(any) any
	Remove(any) any
	ContainsKey(any) bool
	ContainsValue(any) bool
	Traversal(...Visitor) // 遍历map
}
