package recursion

// 斐波那契数列

import "math"

// 时间复杂度O(2^n) 空间复杂度 O(n)
func fib0(n int) int {
	if n <= 2 {
		return 1
	}
	return fib0(n-1) + fib0(n-2)
}

// 用数组存放计算过的结果，避免重复计算 时间复杂度O(n) 空间复杂度 O(n)
func fib1(n int) int {
	if n <= 2 {
		return 1
	}
	array := make([]int, n+1)
	array[1], array[2] = 1, 1
	var fib func(int, []int) int
	fib = func(n int, array []int) int {
		if array[n] == 0 {
			array[n] = fib(n-1, array) + fib(n-2, array)
		}
		return array[n]
	}
	return fib(n, array)
}

// 去除递归调用 时间复杂度O(n) 空间复杂度 O(n)
func fib2(n int) int {
	if n <= 2 {
		return 1
	}
	array := make([]int, n+1)
	array[1], array[2] = 1, 1
	for i := 3; i <= n; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array[n]
}

// 由于每次运算只需要用到数组中的 2 个元素，所以可以使用滚动数组来优化 时间复杂度O(n) 空间复杂度 O(1)
func fib3(n int) int {
	if n <= 2 {
		return 1
	}
	array := make([]int, 2)
	array[0], array[1] = 1, 1
	for i := 3; i <= n; i++ {
		array[i%2] = array[(i-1)%2] + array[(i-2)%2]
	}
	return array[n%2]
}

// 由于每次运算只需要用到数组中的 2 个元素，所以可以使用滚动数组来优化 时间复杂度O(n) 空间复杂度 O(1)
// 位运算取代模运算 时间复杂度O(n) 空间复杂度 O(1)
func fib4(n int) int {
	if n <= 2 {
		return 1
	}
	array := make([]int, 2)
	array[0], array[1] = 1, 1
	for i := 3; i <= n; i++ {
		array[i&1] = array[(i-1)&1] + array[(i-2)&1]
	}
	return array[n&1]
}

// 去除数组  时间复杂度O(n) 空间复杂度 O(1)
func fib5(n int) int {
	if n <= 2 {
		return 1
	}
	first, second := 1, 1
	for i := 3; i <= n; i++ {
		second = first + second
		first = second - first
	}
	return second
}

// 特征方程  时间复杂度、空间复杂度取决于 pow 函数（至少可以低至O(logn) ）
func fib6(n int) int {
	c := math.Sqrt(5)
	return (int)((math.Pow((1+c)/2, float64(n)) - math.Pow((1-c)/2, float64(n))) / c)
}
