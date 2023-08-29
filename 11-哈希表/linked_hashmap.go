package hashmap

import (
	"DataStruct_Go/utils"
	"fmt"
)

type linkedHashMap struct {
	m     hashmap
	first *hashNode
	last  *hashNode
}

func (l *linkedHashMap) Size() int {
	return l.m.size
}

func (l *linkedHashMap) IsEmpty() bool {
	return l.m.size == 0
}

func (l *linkedHashMap) Put(key any, value any) any {
	l.m.Put(key, value)
	node := l.Get(key).(*hashNode)
	//创建节点并且维护prev和next
	if l.first == nil {
		l.first = node
		l.last = node
	} else {
		l.last.next = node
		node.prev = l.last
		l.last = node
	}
	return node.value
}

func (l *linkedHashMap) Get(key any) any {
	return l.m.Get(key)
}

func (l *linkedHashMap) Remove(key any) any {
	return l.m.Remove(key)
}

func (l *linkedHashMap) ContainsKey(key any) bool {
	return l.m.ContainsKey(key)
}

func (l *linkedHashMap) ContainsValue(value any) bool {
	node := l.first
	for node != nil {
		if utils.Equal(value, node.value) {
			return true
		}
		node = node.next
	}
	return false
}
func (l *linkedHashMap) afterRemove(willNode, removedNode *hashNode) {
	//交换两个节点
	node1, node2 := willNode, removedNode
	if node1 != node2 {
		// 交换linkedWillNode和linkedRemovedNode在链表中的位置
		// 交换prev
		tmp := node1.prev
		node1.prev = node2.prev
		node2.prev = tmp
		if node1.prev == nil {
			l.first = node1
		} else {
			node1.prev.next = node1
		}
		if node2.prev == nil {
			l.first = node2
		} else {
			node2.prev.next = node2
		}

		// 交换next
		tmp = node1.next
		node1.next = node2.next
		node2.next = tmp
		if node1.next == nil {
			l.last = node1
		} else {
			node1.next.prev = node1
		}
		if node2.next == nil {
			l.last = node2
		} else {
			node2.next.prev = node2
		}
	}

	prev := node2.prev
	next := node2.next
	if prev == nil {
		l.first = next
	} else {
		prev.next = next
	}

	if next == nil {
		l.last = prev
	} else {
		next.prev = prev
	}
}

func (l *linkedHashMap) Traversal(visitors ...utils.Visitor) {
	visitor := utils.NewVisitor()
	if len(visitors) == 1 {
		visitor = visitors[0]
	}
	node := l.first
	for node != nil {
		if visitor.Visit(fmt.Sprintf("[%v:%v]", node.key, node.value)) {
			return
		}
		node = node.next
	}
}

func (l *linkedHashMap) Clear() {
	l.m.Clear()
	l.first = nil
	l.last = nil
}
