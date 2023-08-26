package linked_list

// SingleCircleLinkedList 单向循环链表:和单链表的区别在 添加/删除头结点时需要维护最后一个的next
type SingleCircleLinkedList struct {
	SingleLinkedList
}

func (s *SingleCircleLinkedList) Add(element any) {
	s.AddAtIndex(s.size, element)
}

func (s *SingleCircleLinkedList) AddAtIndex(index int, element any) {
	s.RangeCheckForAdd(index)
	if index == 0 {
		newHead := NewNode(nil, element, s.head)
		// 拿到最后一个节点
		last := newHead
		if s.head != nil { // size > 0
			last = s.node(s.Size() - 1)
			last.next = newHead // 维护最后一个的next为新的first
		}
		s.head = newHead

	} else {
		prev := s.node(index - 1)
		prev.next = NewNode(nil, element, prev.next)
	}
	s.size++
}

func (s *SingleCircleLinkedList) Remove(index int) any {
	s.RangeCheck(index)
	node := s.head
	if index == 0 {
		if s.size == 1 { // 只有一个元素
			s.head = nil
		} else {
			s.head = s.head.next
			last := s.node(s.size - 1)
			last.next = s.head
		}
	} else {
		prev := s.node(index - 1) // 获取前一个节点
		node = prev.next
		prev.next = node.next
	}
	s.size--
	return node.element
}
