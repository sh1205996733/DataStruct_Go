package binarytree

import (
	"DataStruct_Go/utils"
	"reflect"
	"strings"
)

type Comparator func(e1, e2 any) int

// BST 二叉搜索树
type BST struct {
	BinaryTree
	Comparator Comparator
	BSTClass   reflect.Value //子类对象，继承者需要给此字段赋值
}

// Add 添加元素
func (bst *BST) Add(value any) { // 添加元素
	valueNotNullCheck(value)
	if bst.Root == nil {
		bst.Root = bst.createNode(value, nil)
		bst.Size++
		// 新添加节点之后的处理
		bst.AfterAdd(bst.Root)
		return
	}
	// 添加只能是叶子节点，所以必须找到其父节点
	current := bst.Root
	parent := bst.Root
	cmp := 0
	for current != nil {
		cmp = bst.compare(value, current.Value)
		parent = current
		if cmp > 0 {
			current = current.Right
		} else if cmp < 0 {
			current = current.Left
		} else { // 相等 直接覆盖
			current.Value = value
			return
		}
	}
	// 看看插入到父节点的哪个位置
	newNode := bst.createNode(value, parent)
	if cmp > 0 {
		parent.Right = newNode
	} else {
		parent.Left = newNode
	}
	bst.Size++
	// 新添加节点之后的处理
	bst.AfterAdd(newNode)
}

// 创建节点
func (bst *BST) createNode(value any, parent *Node) *Node {
	if bst.BSTClass.IsValid() {
		method := bst.BSTClass.MethodByName("CreateNode")
		if method.IsValid() {
			return method.Call([]reflect.Value{reflect.ValueOf(value), reflect.ValueOf(parent)})[0].Interface().(*Node)
		}
	} else {
		return bst.CreateNode(value, parent)
	}
	return nil
}

// AfterAdd 新添加节点之后的处理
func (bst *BST) AfterAdd(newNode *Node) {
	// 新添加节点之后的处理
	if bst.BSTClass.IsValid() {
		method := bst.BSTClass.MethodByName("AfterAdd")
		if method.IsValid() {
			method.Call([]reflect.Value{reflect.ValueOf(newNode)})
		}
	}
}

// Remove 删除元素
func (bst *BST) Remove(value any) { // 删除元素
	remove(bst.node(value), bst)
}

// 删除指定节点
func remove(node *Node, bst *BST) {
	if node == nil {
		return
	}
	bst.Size--
	if node.HasTwoChildren() { //删除度为2的节点时，找到他的前驱或者后继节点 先交换两值 删除前驱或者后继节点
		// 找到前驱/后继节点
		predecessor := bst.Successor(node)
		//predecessor := Predecessor(node)
		node.Value = predecessor.Value //用前驱或者后继节点的值覆盖度为2的节点的值
		node = predecessor             //使用前驱或者后继节点替代node
	}
	// 删除node节点（node的度必然是1或者0）

	replacement := node.Left
	if node.Left == nil {
		replacement = node.Right
	}
	if replacement != nil { // node是度为1的节点并且是根节点
		// 更改parent
		replacement.Parent = node.Parent
		if node.Parent == nil { //删除度为0,直接删除
			bst.Root = replacement
		} else if node.Parent.Left == node {
			node.Parent.Left = replacement
		} else {
			node.Parent.Right = replacement
		}
		// 删除节点之后的处理
		bst.AfterRemove(replacement)
	} else if node.Parent == nil { // node是叶子节点并且是根节点
		bst.Root = nil
		// 删除节点之后的处理
		bst.AfterRemove(node)
	} else { //删除度为0
		if node == node.Parent.Left {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
		// 删除节点之后的处理
		bst.AfterRemove(node)
	}
}

// AfterRemove 新添加节点之后的处理
func (bst *BST) AfterRemove(newNode *Node) {
	// 新添加节点之后的处理
	if bst.BSTClass.IsValid() {
		method := bst.BSTClass.MethodByName("AfterRemove")
		if method.IsValid() {
			method.Call([]reflect.Value{reflect.ValueOf(newNode)})
		}
	}
}

// Contains 是否包含某元素
func (bst *BST) Contains(value any) bool {
	return bst.node(value) != nil
}

// 根据元素内容找节点
func (bst *BST) node(value any) *Node {
	node := bst.Root
	for node != nil {
		cmp := bst.compare(value, node.Value)
		if cmp == 0 {
			return node
		}
		if cmp > 0 {
			node = node.Right
		} else { // cmp < 0
			node = node.Left
		}
	}
	return nil
}

// 返回值等于0，代表e1和e2相等；返回值大于0，代表e1大于e2；返回值小于于0，代表e1小于e2
func (bst *BST) compare(e1, e2 any) int {
	if bst.Comparator != nil {
		return bst.Comparator(e1, e2)
	} //todo 待优化
	if reflect.TypeOf(e1) != reflect.TypeOf(e2) {
		panic("类型不一致！")
	} else {
		var cmp int
		var err error
		switch v := e1.(type) {
		case utils.Comparable:
			cmp, err = (e1.(utils.Comparable)).CompareTo(e2)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			cmp = v.(int) - e2.(int)
		case string:
			cmp = strings.Compare(v, e2.(string))
		case float64, float32:
			cmp = int(v.(float64) - e2.(float64))
		default:
			panic("类型不能比较")
		}
		if err != nil {
			panic("missing method CompareTo")
		} else {
			return cmp
		}
	}
	return 0
}

func valueNotNullCheck(value any) {
	if value == nil {
		panic("value must not be null")
	}
}
