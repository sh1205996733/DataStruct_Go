package search

import (
	"fmt"
	"testing"
)

func TestLowerBound0(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	target := 6
	fmt.Println(lower_bound1(nums, target))
	fmt.Println(lower_bound2(nums, target))
	fmt.Println(lower_bound3(nums, target))
}
