package set

import (
	linked_list "DataStruct_Go/03-链表"
	"DataStruct_Go/utils"
)

// listSet 线性集合
type listSet struct {
	list utils.List
}

func NewListSet() utils.Set {
	return &listSet{
		list: new(linked_list.SingleLinkedList),
	}
}

func (l *listSet) Size() int {
	return l.list.Size()
}

func (l *listSet) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *listSet) Clear() {
	l.list.Clear()
}

func (l *listSet) Contains(element any) bool {
	return l.list.Contains(element)
}

func (l *listSet) Add(element any) {
	index := l.list.IndexOf(element)
	if index != utils.ELEMENT_NOT_FOUND { // 存在就覆盖
		l.list.Set(index, element)
	} else { // 不存在就添加
		l.list.Add(element)
	}
}

func (l *listSet) Remove(element any) {
	index := l.list.IndexOf(element)
	if index != utils.ELEMENT_NOT_FOUND {
		l.list.Remove(index)
	}
}

func (l *listSet) Traversal(visitors ...utils.Visitor) {
	visitor := utils.NewVisitor()
	if len(visitors) == 1 {
		visitor = visitors[0]
	}
	size := l.list.Size()
	for i := 0; i < size; i++ {
		if visitor.Visit(l.list.Get(i)) {
			return
		}
	}
}
