package linked_list

import (
	"DataStruct_Go/utils"
	"fmt"
	"strings"
)

// DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	BaseList
	head *node // 头节点
	tail *node // 尾节点
}

func (d *DoubleLinkedList) Get(index int) any {
	return d.node(index).element
}

func (d *DoubleLinkedList) Add(element any) {
	d.AddAtIndex(d.size, element)
}

func (d *DoubleLinkedList) AddAtIndex(index int, element any) {
	d.RangeCheckForAdd(index)
	if index == d.size { //尾部
		oldLast := d.tail
		d.tail = NewNode(oldLast, element, nil)
		if oldLast == nil { // 这是链表添加的第一个元素
			d.head = d.tail
		} else {
			oldLast.next = d.tail
		}
	} else {
		next := d.node(index)
		prev := next.prev
		node := NewNode(prev, element, next)
		if prev == nil {
			d.head = node
		} else {
			prev.next = node
		}
		next.prev = node

	}
	d.size++
}

func (d *DoubleLinkedList) Set(index int, element any) any {
	n := d.node(index)
	oldEle := n.element
	n.element = element
	return oldEle
}

func (d *DoubleLinkedList) Remove(index int) any {
	d.RangeCheck(index)
	node := d.node(index) // 获取指定位置的节点
	prev := node.prev     // 获取指定位置的上一个节点
	next := node.next     // 获取指定位置的下一个节点
	if prev == nil {      // 删除第一个 index == 0
		d.head = next
	} else {
		prev.next = next
	}
	if next == nil { //删除最后一个 index == size-1
		d.tail = prev
	} else {
		next.prev = prev
	}
	d.size--
	return node.element
}

func (d *DoubleLinkedList) Contains(element any) bool {
	return d.IndexOf(element) != utils.ELEMENT_NOT_FOUND
}

func (d *DoubleLinkedList) IndexOf(element any) int {
	if element == nil {
		node := d.head
		for i := 0; i < d.size; i++ {
			if node.element == nil {
				return i
			}
			node = node.next
		}
	} else {
		node := d.head
		for i := 0; i < d.size; i++ {
			if node.element == element {
				return i
			}
			node = node.next
		}
	}
	return utils.ELEMENT_NOT_FOUND
}

func (d *DoubleLinkedList) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *DoubleLinkedList) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("size=%v, [", d.size))
	node := d.head
	for i := 0; i < d.size; i++ {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(node.String())
		node = node.next
	}
	sb.WriteString("]")
	return sb.String()
}

// 获取指定位置上的元素
func (d *DoubleLinkedList) node(index int) *node {
	d.RangeCheck(index)
	if index < (d.size >> 1) {
		node := d.head
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	} else {
		node := d.tail
		for i := d.size - 1; i > index; i-- {
			node = node.prev
		}
		return node
	}
}
