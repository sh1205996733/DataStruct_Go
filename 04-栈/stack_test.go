package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(11)
	stack.Push(22)
	stack.Push(33)
	stack.Push(44)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
