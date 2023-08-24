package main

import (
	"DataStruct_Go/utils"
	"fmt"
)

func main() {
	n := 5
	//fmt.Println(fib1(5))
	//fmt.Println(fib2(5))
	utils.TimeTool("fib1", func() {
		fmt.Println(fib1(n))
	})
	utils.TimeTool("fib2", func() {
		fmt.Println(fib3(n))
	})
}

/* 0 1 2 3 4 5
 * 0 1 1 2 3 5 8 13 ....
 */

// O(2^n)
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

// O(n)
func fib2(n int) int {
	if n <= 1 {
		return n
	}

	first, second := 0, 1
	for i := 0; i < n-1; i++ {
		sum := first + second
		first = second
		second = sum
	}
	return second
}
func fib3(n int) int {
	if n <= 1 {
		return n
	}

	first, second := 0, 1
	for i := n - 1; i > 1; i-- {
		second += first
		first = second - first
	}
	return second
}

func test1(n int) {
	// 汇编指令

	// 1
	if n > 10 {
		fmt.Println("n > 10")
	} else if n > 5 { // 2
		fmt.Println("n > 5")
	} else {
		fmt.Println("n <= 5")
	}

	// 1 + 4 + 4 + 4
	for i := 0; i < 4; i++ {
		fmt.Println("test")
	}

	// 140000
	// O(1)
	// O(1)
}

func test2(n int) {
	// O(n)
	// 1 + 3n
	for i := 0; i < n; i++ {
		fmt.Println("test")
	}
}

func test3(n int) {
	// 1 + 2n + n * (1 + 3n)
	// 1 + 2n + n + 3n^2
	// 3n^2 + 3n + 1
	// O(n^2)

	// O(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("test")
		}
	}
}

func test4(n int) {
	// 1 + 2n + n * (1 + 45)
	// 1 + 2n + 46n
	// 48n + 1
	// O(n)
	for i := 0; i < n; i++ {
		for j := 0; j < 15; j++ {
			fmt.Println("test")
		}
	}
}

func test5(n int) {
	// 8 = 2^3
	// 16 = 2^4

	// 3 = log2(8)
	// 4 = log2(16)

	// 执行次数 = log2(n)
	// O(logn)
	for n/2 > 0 {
		n = n / 2
		fmt.Println("test")
	}
}

func test6(n int) {
	// log5(n)
	// O(logn)
	for n/5 > 0 {
		n = n / 2
		fmt.Println("test")
	}
}

func test7(n int) {
	// 1 + 2*log2(n) + log2(n) * (1 + 3n)

	// 1 + 3*log2(n) + 2 * nlog2(n)
	// O(nlogn)
	for i := 1; i < n; i = i * 2 {
		// 1 + 3n
		for j := 0; j < n; j++ {
			fmt.Println("test")
		}
	}
}

func test10(n int) {
	// O(n)
	a, b := 10, 20
	c := a + b
	array := make([]int, n)
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i] + c)
	}
}
