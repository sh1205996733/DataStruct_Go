package tools

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	fmt.Println(Random(10, 2, 5))
}
func TestSame(t *testing.T) {
	fmt.Println(Same(10, 6))
}
func TestCombine(t *testing.T) {
	a1 := Same(10, 6)
	a2 := Same(10, 3)
	fmt.Println(Combine(a1, a2))
}
func TestHeadTailAscOrder(t *testing.T) {
	fmt.Println(HeadTailAscOrder(2, 7, 4))
}
func TestCenterAscOrder(t *testing.T) {
	fmt.Println(CenterAscOrder(2, 20, 7))
}
func TestHeadAscOrder(t *testing.T) {
	fmt.Println(HeadAscOrder(2, 20, 7))
}
func TestTailAscOrder(t *testing.T) {
	fmt.Println(TailAscOrder(2, 20, 7))
}
func TestCopy(t *testing.T) {
	a := Random(10, 2, 5)
	fmt.Println(Copy(a))
}
