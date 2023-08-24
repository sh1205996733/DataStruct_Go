package avl

import (
	"DataStruct_Go/binarytree"
	"DataStruct_Go/binarytree/bbst"
	"DataStruct_Go/binarytree/util"
	"math"
)

// AVL树
type AVLTree struct {
	bbst.BBST
}

func NewAVLTree() *AVLTree {
	avl := new(AVLTree)
	avl.Init()
	return avl
}

func (avl *AVLTree) Init() {
	avl.BSTClass = util.ValueOf(avl)
}

func (avl *AVLTree) CreateNode(value interface{}, parent *binarytree.Node) *binarytree.Node {
	avlNode := new(avlNode)
	node := binarytree.Node{Value: value, Parent: parent, Class: avlNode}
	avlNode.height = 1
	avlNode.Node = &node
	return avlNode.Node
}

func (avl *AVLTree) AfterAdd(node *binarytree.Node) {
	node = node.Parent
	for node != nil {
		if isBalanced(node) { //父节点是否平衡
			// 更新高度
			updateHeight(node)
		} else { //父节点、祖父节点必然不为空
			// 恢复平衡
			avl.rebalance(node) //只需要恢复一次就可以了
			// 整棵树恢复平衡
			break
		}
		node = node.Parent
	}
}

func (avl *AVLTree) AfterRemove(node *binarytree.Node) {
	node = node.Parent
	for node != nil {
		if isBalanced(node) {
			// 更新高度
			updateHeight(node)
		} else {
			// 恢复平衡
			avl.rebalance(node)
		}
		node = node.Parent
	}
}

/**
 * 恢复平衡
 * @param grand 高度最低的那个不平衡节点
 */
func (avl *AVLTree) rebalance(grand *binarytree.Node) {
	parent := tallerChild(grand) //最高的child 必然不为空
	node := tallerChild(parent)  //最高的child，可能为空
	if parent.IsLeftChild() {    //L
		if node.IsLeftChild() { //LL
			avl.RotateRight(grand, avl)
		} else { //LR
			avl.RotateLeft(parent, avl)
			avl.RotateRight(grand, avl)
		}
	} else { //R
		if node.IsLeftChild() { //RL
			avl.RotateRight(parent, avl)
			avl.RotateLeft(grand, avl)
		} else { //RR
			avl.RotateLeft(grand, avl)
		}
	}
}

/**
 * 恢复平衡(统一操作)
 * @param grand 高度最低的那个不平衡节点
 */
func (avl *AVLTree) rebalance01(grand *binarytree.Node) {
	parent := tallerChild(grand) //最高的child 必然不为空
	node := tallerChild(parent)  //最高的child，可能为空
	if parent.IsLeftChild() {    // L
		if node.IsLeftChild() { // LL
			avl.Rotate(grand, node, node.Right, parent, parent.Right, grand, avl)
		} else { // LR
			avl.Rotate(grand, parent, node.Left, node, node.Right, grand, avl)
		}
	} else { // R
		if node.IsLeftChild() { // RL
			avl.Rotate(grand, grand, node.Left, node, node.Right, parent, avl)
		} else { // RR
			avl.Rotate(grand, grand, parent.Left, parent, node.Left, node, avl)
		}
	}
}

func (avl *AVLTree) CallFunc(grand, parent *binarytree.Node, param ...*binarytree.Node) {
	updateHeight(parent)
	updateHeight(grand)
	if len(param) == 1 {
		updateHeight(param[0])
	}
}

// 更新高度
func updateHeight(node *binarytree.Node) {
	node.Class.(*avlNode).updateHeight()
}

// 是否平衡：左右高度差(平衡因子)绝对值小于等于1
func isBalanced(node *binarytree.Node) bool {
	return node.Class.(*avlNode).balanceFactor() <= 1
}

type avlNode struct {
	Node   *binarytree.Node
	height int
}

func (t *avlNode) leftHeight() int {
	if t.Node.Left == nil || t.Node.Left.Class == nil {
		return 0
	} else {
		return t.Node.Left.Class.(*avlNode).height
	}
}
func (t *avlNode) rightHeight() int {
	if t.Node.Right == nil || t.Node.Right.Class == nil {
		return 0
	} else {
		return t.Node.Right.Class.(*avlNode).height
	}
}

// 更新高度height (最高的子树高度加1)
func (t *avlNode) updateHeight() {
	t.height = util.Max(t.leftHeight(), t.rightHeight()) + 1
}

// 计算平衡因子(左右高度差的绝对值)
func (t *avlNode) balanceFactor() int {
	leftHeight := t.leftHeight()
	rightHeight := t.rightHeight()
	balance := math.Abs(float64(leftHeight - rightHeight))
	return int(balance)
}

// 返回node 最高的子节点
func tallerChild(node *binarytree.Node) *binarytree.Node {
	avl := node.Class.(*avlNode)
	leftHeight := avl.leftHeight()
	rightHeight := avl.rightHeight()
	if leftHeight > rightHeight {
		return node.Left
	}
	if leftHeight < rightHeight {
		return node.Right
	}
	return util.If(node.IsLeftChild(), node.Left, node.Right).(*binarytree.Node) //相等看一下父节点是左子树还是右子树
}

func (t *avlNode) ToString() string {
	parentString := "nil"
	if t.Node.Parent != nil {
		parentString = t.Node.Parent.String()
	}
	return t.Node.String() + "_p(" + parentString + ")_h(" + util.ToString(t.height) + ")"
}
