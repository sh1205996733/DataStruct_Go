package skiplist

import (
	"DataStruct_Go/utils"
	"fmt"
	"math/rand"
	"strings"
)

// 跳表
type skiplist struct {
	size  int
	level int       // 有效层数
	first *skipNode // 不存放任何K-V
}

const (
	MAX_LEVEL = 32   // 最大层数
	P         = 0.25 //
)

func NewSkipList() *skiplist {
	return &skiplist{
		first: createNode(nil, nil, MAX_LEVEL),
	}
}

func (s *skiplist) Size() int {
	return s.size
}

func (s *skiplist) IsEmpty() bool {
	return s.size == 0
}

func (s *skiplist) Get(key any) any {
	keyCheck(key)
	node := s.first
	for i := s.level - 1; i >= 0; i-- {
		for node.nexts[i] != nil && utils.Compare(key, node.nexts[i].key) > 0 {
			node = node.nexts[i]
		}
		// node.nexts[i].key >= key
		if node.nexts[i] != nil && utils.Compare(key, node.nexts[i].key) == 0 { // 节点是存在的,直接覆盖
			return node.nexts[i].value
		}
	}
	return nil
}

func (s *skiplist) Put(key any, value any) any {
	keyCheck(key)
	node := s.first
	prevs := make([]*skipNode, s.level)
	for i := s.level - 1; i >= 0; i-- {
		for node.nexts[i] != nil && utils.Compare(key, node.nexts[i].key) > 0 {
			node = node.nexts[i]
		}
		if node.nexts[i] != nil && utils.Compare(key, node.nexts[i].key) == 0 { // 节点是存在的,直接覆盖
			oldV := node.nexts[i].value
			node.nexts[i].value = value
			return oldV
		}
		prevs[i] = node
	}
	// 新节点的层数
	newLevel := randomLevel()
	newNode := createNode(key, value, newLevel)
	// 设置前驱和后继
	for i := 0; i < newLevel; i++ {
		if i >= s.level {
			s.first.nexts[i] = newNode
		} else {
			newNode.nexts[i] = prevs[i].nexts[i]
			prevs[i].nexts[i] = newNode
		}
	}
	// 节点数量增加
	s.size++
	// 计算跳表的最终层数
	s.level = utils.Max(s.level, newLevel)
	return value
}

func (s *skiplist) Remove(key any) any {
	keyCheck(key)
	node := s.first
	prevs := make([]*skipNode, s.level)
	var exist = false
	for i := s.level - 1; i >= 0; i-- {
		for node.nexts[i] != nil && utils.Compare(key, node.nexts[i].key) > 0 {
			node = node.nexts[i]
		}
		prevs[i] = node
		if node.nexts[i] != nil && utils.Compare(key, node.nexts[i].key) == 0 {
			exist = true
		}
	}
	if !exist {
		return nil
	}
	// 需要被删除的节点
	removedNode := node.nexts[0]
	s.size--
	// 设置后继
	for i := 0; i < len(removedNode.nexts); i++ {
		prevs[i].nexts[i] = removedNode.nexts[i]
	}
	// 更新跳表的层数
	for newLevel := s.level - 1; newLevel >= 0 && s.first.nexts[newLevel] == nil; newLevel-- {
		s.level = newLevel
	}
	return removedNode.value
}

func (s *skiplist) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("一共%d层\n", s.level))
	for i := s.level - 1; i >= 0; i-- {
		node := s.first
		for node.nexts[i] != nil {
			sb.WriteString(fmt.Sprintf("%v ", node.nexts[i]))
			node = node.nexts[i]
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func randomLevel() int {
	var level int
	for p := 0.00; p < P && level < MAX_LEVEL; p = float64(rand.Intn(100)) * 0.01 {
		level++
	}
	return level
}

func keyCheck(key any) {
	if key == nil {
		panic("key must not be null.")
	}
}

type skipNode struct {
	key   any
	value any
	nexts []*skipNode
}

func createNode(key, value any, level int) *skipNode {
	return &skipNode{
		key:   key,
		value: value,
		nexts: make([]*skipNode, level),
	}
}
func (s *skipNode) String() string {
	return fmt.Sprintf("%s:%s_%d", s.key, s.value, len(s.nexts))
}
