package stack

import (
	linked_list "DataStruct_Go/03-链表"
	"DataStruct_Go/utils"
)

// Stack 栈 先进后出(FILO) 基于链表和数组都可以
type Stack struct {
	list utils.List
}

func NewStack() *Stack {
	return &Stack{
		//list: array.NewArrayList(),
		list: &linked_list.SingleLinkedList{},
	}
}
func (s *Stack) Clear() {
	s.list.Clear()
}

func (s *Stack) Size() int {
	return s.list.Size()
}

func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack) Push(element any) {
	s.list.Add(element)
}

func (s *Stack) Pop() any {
	return s.list.Remove(s.Size() - 1)
}

func (s *Stack) Peek() any {
	return s.list.Get(s.Size() - 1)
}
