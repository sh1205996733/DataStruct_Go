package recursion

//尾调用 一个函数的最后一个动作是调用函数,如果最后一个动作是调用自身，称为尾递归（Tail Recursion），是尾调用的特殊情况

// 求n阶乘
func facttorial0(n int) int {
	if n <= 1 {
		return n
	}
	return facttorial0(n-1) * n //不是尾调用,因为它最后1个动作是乘法
}

// 尾递归示例1 – 阶乘
func facttorial(n int) int {
	return facttorialN(n, 1)
}

func facttorialN(n int, ret int) int {
	if n <= 1 {
		return ret
	}
	return facttorialN(n-1, ret*n)
}

// 尾递归示例2 – 斐波那契数列
func fib(n int) int {
	return fibN(n, 1, 1)
}

func fibN(n int, first, second int) int {
	if n <= 2 {
		return second
	}
	return fibN(n-1, second, first+second)
}
