package utils

// If 条件返回 三元运算
func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}
