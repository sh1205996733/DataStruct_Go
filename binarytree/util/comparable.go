package util

// 比较器
type Comparable interface {
	CompareTo(interface{}) (int, error)
}
