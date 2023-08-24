package printer

import (
	"reflect"
	"strings"
)

// 去重
func repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

// 去除重复空格
func blank(length int) string {
	if length <= 0 {
		return ""
	}
	return strings.Repeat(" ", length)
}

// 获取类型
func typeOf(btNode interface{}, funcName string) interface{} {
	return reflect.ValueOf(btNode).MethodByName(funcName).Call(nil)[0].Interface()
}
