package utils

import (
	"reflect"
	"strings"
)

// Comparable 比较器
type Comparable interface {
	CompareTo(interface{}) (int, error)
}

// Compare 返回值等于0，代表e1和e2相等；返回值大于0，代表e1大于e2；返回值小于于0，代表e1小于e2
func Compare(e1, e2 any) int {
	if e1 == nil {
		return -1
	}
	if e2 == nil {
		return 1
	}
	if reflect.TypeOf(e1) != reflect.TypeOf(e2) {
		panic("类型不一致！")
	} else {
		var cmp int
		var err error
		switch v := e1.(type) {
		case Comparable:
			cmp, err = (e1.(Comparable)).CompareTo(e2)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			cmp = v.(int) - e2.(int)
		case string:
			cmp = strings.Compare(v, e2.(string))
		case float64, float32:
			cmp = int(v.(float64) - e2.(float64))
		default:
			panic("类型不能比较")
		}
		if err != nil {
			panic("missing method CompareTo")
		} else {
			return cmp
		}
	}
	return 0
}
