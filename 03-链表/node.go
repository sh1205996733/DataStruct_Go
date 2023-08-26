package linked_list

import (
	"fmt"
	"strings"
)

type node struct {
	element any
	prev    *node // 前一个节点
	next    *node // 下一个节点
}

func NewNode(prev *node, element any, next *node) *node {
	return &node{
		element: element,
		prev:    prev,
		next:    next,
	}
}

func (n *node) String() string {
	sb := strings.Builder{}
	if n.prev != nil {
		sb.WriteString(fmt.Sprintf("%v", n.prev.element))
	} else {
		sb.WriteString("nil")
	}

	sb.WriteString(fmt.Sprintf("_%v_", n.element))

	if n.next != nil {
		sb.WriteString(fmt.Sprintf("%v", n.next.element))
	} else {
		sb.WriteString("nil")
	}

	return sb.String()
}
