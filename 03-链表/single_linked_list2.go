package linked_list

import (
	"DataStruct_Go/utils"
	"fmt"
	"strings"
)

// SingleLinkedListV2 单链表 带虚拟头节点
type SingleLinkedListV2 struct {
	BaseList
	head *node // 头节点
}

// Get 最好：O(1) 最坏：O(n) 平均：O(n)
func (s *SingleLinkedListV2) Get(index int) any {
	return s.node(index).element
}

// Add 最好：O(1) 最坏：O(n) 平均：O(n)
func (s *SingleLinkedListV2) Add(element any) {
	s.AddAtIndex(s.size, element)
}

// AddAtIndex 最好：O(1) 最坏：O(n) 平均：O(n)
func (s *SingleLinkedListV2) AddAtIndex(index int, element any) {
	s.RangeCheckForAdd(index)
	prev := s.head
	if index > 0 {
		prev = s.node(index - 1)
	}
	prev.next = NewNode(nil, element, prev.next)
	s.size++
}

// Set 最好：O(1) 最坏：O(n) 平均：O(n)
func (s *SingleLinkedListV2) Set(index int, element any) any {
	n := s.node(index)
	oldEle := n.element
	n.element = element
	return oldEle
}

// Remove 最好：O(1) 最坏：O(n) 平均：O(n)
func (s *SingleLinkedListV2) Remove(index int) any {
	s.RangeCheck(index)
	prev := s.head
	if index > 0 {
		prev = s.node(index - 1)
	}
	old := prev.next
	prev.next = old.next
	s.size--
	return old.element
}

func (s *SingleLinkedListV2) Contains(element any) bool {
	return s.IndexOf(element) != utils.ELEMENT_NOT_FOUND
}

func (s *SingleLinkedListV2) IndexOf(element any) int {
	if element == nil {
		cur := s.head.next
		for i := 0; i < s.size; i++ {
			if cur.element == nil {
				return i
			}
			cur = cur.next
		}
	} else {
		cur := s.head.next
		for i := 0; i < s.size; i++ {
			if cur.element == element {
				return i
			}
			cur = cur.next
		}
	}
	return utils.ELEMENT_NOT_FOUND
}

func (s *SingleLinkedListV2) Clear() {
	s.head = nil
	s.size = 0
}

func (s *SingleLinkedListV2) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("size=%v, [", s.size))
	node := s.head.next
	for i := 0; i < s.size; i++ {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(node.String())
		node = node.next
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *SingleLinkedListV2) node(index int) *node {
	s.RangeCheck(index)
	node := s.head.next
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
