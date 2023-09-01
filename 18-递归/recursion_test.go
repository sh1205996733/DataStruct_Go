package recursion

import (
	"DataStruct_Go/utils"
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	n := 9
	utils.Times("fib0", func() {
		fmt.Println(fib0(n))
	})
	utils.Times("fib1", func() {
		fmt.Println(fib1(n))
	})
	utils.Times("fib2", func() {
		fmt.Println(fib2(n))
	})
	utils.Times("fib3", func() {
		fmt.Println(fib3(n))
	})
	utils.Times("fib4", func() {
		fmt.Println(fib4(n))
	})
	utils.Times("fib5", func() {
		fmt.Println(fib5(n))
	})
	utils.Times("fib6", func() {
		fmt.Println(fib6(n))
	})
}

func TestHanio(t *testing.T) {
	hanio(3, "A", "B", "C")
}

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(9))
}

func TestFacttorial(t *testing.T) {
	fmt.Println(facttorial(5))
}
