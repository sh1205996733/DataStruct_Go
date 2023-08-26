package utils

func Asserts(ret bool) {
	if !ret {
		panic("测试未通过")
	}
}
