package utils

// Comparable 比较器
type Comparable interface {
	CompareTo(interface{}) (int, error)
}
