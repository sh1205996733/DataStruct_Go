package util

import (
	"fmt"
	"reflect"
)

// 获取left、right最大值
func Max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

// 条件返回 三元运算
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 转换成字符串
func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// 转换成字符串
func ValueOf(value interface{}) reflect.Value {
	return reflect.ValueOf(value)
}
