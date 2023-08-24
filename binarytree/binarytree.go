package binarytree

import (
	"container/list"
	"fmt"
	"reflect"
)

// 二叉树
type BinaryTree struct {
	Size int
	Root *Node
}
type Node struct {
	Value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
	Class  interface{} //子类对象，继承者需要给此字段赋值
}

func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}
func (n *Node) HasTwoChildren() bool {
	return n.Left != nil && n.Right != nil
}
func (n *Node) IsLeftChild() bool {
	return n.Parent != nil && n.Parent.Left == n
}
func (n *Node) IsRightChild() bool {
	return n.Parent != nil && n.Parent.Right == n
}

// 返回兄弟节点
func (n *Node) Subling() *Node {
	if n.IsLeftChild() {
		return n.Parent.Right
	}
	if n.IsRightChild() {
		return n.Parent.Left
	}
	return nil
}
func (n *Node) String() string {
	value := reflect.ValueOf(n.Value).MethodByName("Print")
	if value.IsValid() {
		n.Value = value.Call(nil)[0].Interface()
	}
	return fmt.Sprintf("%v", n.Value)
}

func (n *Node) RootNode() interface{} {
	return n
}
func (n *Node) LNode() interface{} {
	return n.Left
}
func (n *Node) RNode() interface{} {
	return n.Right
}
func (n *Node) ToString() interface{} {
	if n.Class != nil {
		method := reflect.ValueOf(n.Class).MethodByName("ToString")
		if method.IsValid() {
			return method.Call(nil)[0].String()
		}
	} else {
		parentString := "nil"
		if n.Parent != nil {
			parentString = n.Parent.String()
		}
		return n.String() + "_p(" + parentString + ")"
	}
	return ""
}

func (n *Node) ColorOf() bool {
	if n.Class != nil {
		method := reflect.ValueOf(n.Class).MethodByName("ColorOf")
		if method.IsValid() {
			return method.Call(nil)[0].Bool()
		}
	}
	return false
}

type Visitor struct {
	Stop  bool
	Visit func(interface{}) bool
}

func (n *BinaryTree) CreateNode(value interface{}, parent *Node) *Node {
	node := new(Node)
	node.Value = value
	node.Parent = parent
	return node
}

func (n *BinaryTree) RootNode() *Node {
	return n.Root
}

// 元素的数量
func (n *BinaryTree) Count() int {
	return n.Size
}

// 是否为空
func (n *BinaryTree) IsEmpty() bool {
	return n.Size == 0
}

// 清空所有元素
func (n *BinaryTree) Clear() {
	n.Size = 0
	n.Root = nil
}

// 前序遍历
func (n *BinaryTree) PreorderTraversal(visitor ...Visitor) {
	if visitor != nil && len(visitor) == 1 {
		preorderTraversal(n.Root, &visitor[0])
	}
}
func preorderTraversal(root *Node, visitor *Visitor) {
	if root == nil || visitor.Stop {
		return
	}
	visitor.Stop = visitor.Visit(root.Value)
	preorderTraversal(root.Left, visitor)
	preorderTraversal(root.Right, visitor)
}

// 中序遍历
func (n *BinaryTree) InorderTraversal() {
	inorderTraversal(n.Root)
}
func inorderTraversal(node *Node) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left)
	fmt.Print(node.String() + " ")
	inorderTraversal(node.Right)
}

// 后序遍历
func (n *BinaryTree) PostorderTraversal() {
	postorderTraversal(n.Root)
}
func postorderTraversal(node *Node) {
	if node == nil {
		return
	}
	postorderTraversal(node.Left)
	postorderTraversal(node.Right)
	fmt.Print(node.String() + " ")
}

// 层序遍历
func (n *BinaryTree) LevelOrderTraversal() {
	if n.Root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(n.Root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*Node)
		fmt.Println(node.String() + " ")
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

// 前驱节点(中序遍历前一个)
func (n *BinaryTree) Predecessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	//前驱节点在左子树当中（left.right.right.right....）
	p := node.Left
	if p != nil { //node.left != null
		for p.Right != nil {
			p = p.Right
		}
		return p
	}
	// 从父节点、祖父节点中寻找前驱节点（node.parent.parent.parent....）
	for node.Parent != nil && node == node.Parent.Left { //node.left == null && node.parent != null
		node = node.Parent
	}
	// node.left == null && node.parent == null
	// node == node.parent.right
	return node.Parent
}

// 后继节点(中序遍历后一个)
func (n *BinaryTree) Successor(node *Node) *Node {
	if node == nil {
		return nil
	}
	//后继节点在右子树当中（right.left.left.left....）
	p := node.Right
	if p != nil { //node.left != null
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	// 从父节点、祖父节点中寻找前驱节点（node.parent.parent.parent....）
	for node.Parent != nil && node == node.Parent.Right { //node.left == null && node.parent != null
		node = node.Parent
	}
	return node.Parent
}

// 判断是否是完全二叉树
func (n *BinaryTree) IsComplete0() bool {
	if n.Root == nil {
		return false
	}
	queue := list.New()
	queue.PushBack(n.Root)
	isLeaf := false
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*Node)
		if isLeaf && !node.IsLeaf() { // 往后再遍历的节点都必须是叶子节点，但当前节点不是叶子节点 直接返回false
			return false
		}
		if node.HasTwoChildren() { // node.left != null && node.right !=null
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		} else if node.IsLeaf() { // node.left == null && node.right ==null
			isLeaf = true // 遇到第一个叶子节点设置为true，往后再遍历的节点都必须是叶子节点
		} else if node.Left != nil { // node.left != null && node.right ==null
			queue.PushBack(node.Left)
		} else { // node.left == null && node.right != null
			return false
		}
	}
	return true
}

func (n *BinaryTree) IsComplete() bool {
	if n.Root == nil {
		return false
	}
	queue := list.New()
	queue.PushBack(n.Root)
	isLeaf := false
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*Node)
		if isLeaf && !node.IsLeaf() {
			return false
		}
		if node.Left != nil {
			queue.PushBack(node.Left)
		} else if node.Right != nil { //node.left == null && node.right != null
			return false
		}
		//node.left != null
		if node.Right != nil { //node.left != null && node.right != null
			queue.PushBack(node.Right)
		} else { //node.left != null && node.right == null
			isLeaf = true
		}
	}
	return true
}

// 树的高度
func (n *BinaryTree) Height() int {
	fmt.Println(height1(n.Root))
	return height(n.Root)
}

// 递归
func height1(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if leftHeight < rightHeight {
		leftHeight = rightHeight
	}
	return 1 + rightHeight
}

// 非递归
func height(root *Node) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	height := 0
	levelCount := 1
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*Node)
		levelCount--
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		//node.left != null
		if node.Right != nil { //node.left != null && node.right != null
			queue.PushBack(node.Right)
		}
		if levelCount == 0 {
			height++
			levelCount = queue.Len()
		}
	}
	return height
}
