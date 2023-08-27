package utils

import (
	"fmt"
	"reflect"
)

// Max 获取left、right最大值
func Max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

// If 条件返回 三元运算
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// ToString 转换成字符串
func ToString(value any) string {
	return fmt.Sprintf("%v", value)
}

// ValueOf 返回value的类型
func ValueOf(value any) reflect.Value {
	return reflect.ValueOf(value)
}
