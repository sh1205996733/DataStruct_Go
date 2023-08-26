package linked_list

// DoubleCircleLinkedList 双向循环链表:和双向链表的区别在 添加/删除头结点时需要维护最后一个的next
type DoubleCircleLinkedList struct {
	DoubleLinkedList
	current *node // 当前节点
}

func (d *DoubleCircleLinkedList) Add(element any) {
	d.AddAtIndex(d.size, element)
}

func (d *DoubleCircleLinkedList) AddAtIndex(index int, element any) {
	d.RangeCheckForAdd(index)
	if index == d.size { // 往最后面添加元素
		oldLast := d.tail
		d.tail = NewNode(oldLast, element, d.head)
		if d.size == 0 { //第一个元素
			d.head = d.tail
			d.head.next = d.head
			d.head.prev = d.head
		} else {
			oldLast.next = d.tail
			d.head.prev = d.tail
		}
	} else {
		next := d.node(index)
		prev := next.prev
		node := NewNode(prev, element, next)
		prev.next = node
		next.prev = node
		if next == d.head { // index = 0
			d.head = node
		}
	}
	d.size++
}

func (d *DoubleCircleLinkedList) Remove(index int) any {
	d.RangeCheck(index)
	return d.removeNode(d.node(index))
}

// removeNode 删除节点
func (d *DoubleCircleLinkedList) removeNode(node *node) any {
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev
		if node == d.head { //删除头节点
			d.head = next
		}
		if node == d.tail { //删除尾节点
			d.tail = prev
		}
	}
	d.size--
	return node.element
}

// RemoveCurrent 删除当前元素，并将current指向下一个元素
func (d *DoubleCircleLinkedList) RemoveCurrent() any {
	if d.current == nil {
		return nil
	}
	next := d.current.next
	element := d.removeNode(d.current)
	if d.size == 0 {
		d.current = nil
	} else {
		d.current = next
	}
	return element
}

// Reset 重置当前节点
func (d *DoubleCircleLinkedList) Reset() {
	d.current = d.head
}

// Next 当前节点的下一个
func (d *DoubleCircleLinkedList) Next() any {
	if d.current == nil {
		return nil
	}
	d.current = d.current.next
	return d.current.element
}
