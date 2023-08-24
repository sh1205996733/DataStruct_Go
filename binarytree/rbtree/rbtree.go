package rbtree

import (
	"leetcode/binarytree"
	"leetcode/binarytree/bbst"
	"leetcode/binarytree/util"
)

const (
	RED   = false
	BLACK = true
)

// 红黑树
type RBTree struct {
	bbst.BBST
}

func NewRBTree() *RBTree {
	rbTree := new(RBTree)
	rbTree.Init()
	return rbTree
}

func (rb *RBTree) Init() {
	rb.BSTClass = util.ValueOf(rb)
}

func (rb *RBTree) CreateNode(value interface{}, parent *binarytree.Node) *binarytree.Node {
	rbNode := new(rbNode)
	node := binarytree.Node{Value: value, Parent: parent, Class: rbNode}
	rbNode.color = false
	rbNode.Node = &node
	return rbNode.Node
}

func (rb *RBTree) CallFunc(grand, parent *binarytree.Node, param ...*binarytree.Node) {

}
func (rb *RBTree) AfterAdd(node *binarytree.Node) { // 修复性质 4
	parent := node.Parent
	// 添加的是根节点 或者 上溢到达了根节点,染成黑色,直接返回
	if parent == nil {
		black(node)
		return
	}
	// 一共12中情况
	// 先判断父节点是否是黑色：四种父节点是黑色(红<--黑-->null、null<--黑-->红、null<--黑-->null)的不用处理，直接返回
	if isBlack(parent) {
		return
	}

	// 再判断uncle节点是否是红色
	uncle := parent.Subling()
	grand := red(parent.Parent)
	if isRed(uncle) { // 如果uncle节点是红色 【B树节点上溢】
		// 四种uncle点是红色的情况(null<--红-->null<--黑-->null<--红-->null)
		// red(grand)
		black(parent)
		black(uncle)
		// 把祖父节点当做是新添加的节点
		rb.AfterAdd(grand)
		return
	}
	// 如果uncle节点不是红色
	// 四种uncle点是不红色的情况(null<--红-->null<--黑-->(uncle)null、null(uncle)<--黑-->null<--红-->null)
	// 处理方式先染色、然后再进行LL、LR、RR、RL四种旋转

	// red(grand)
	if parent.IsLeftChild() { // L
		if node.IsLeftChild() { // LL
			black(parent)
		} else { // LR
			black(node)
			rb.RotateLeft(parent, rb)
		}
		rb.RotateRight(grand, rb)
	} else { // R
		if node.IsLeftChild() { // RL
			black(node)
			rb.RotateRight(parent, rb)
		} else { // RR
			black(parent)
		}
		rb.RotateLeft(grand, rb)
	}
}

func (rb *RBTree) AfterRemove(node *binarytree.Node) {
	// 1.如果删除的时候红色 或者 用以取代删除节点的子节点是红色 直接返回
	if isRed(node) {
		black(node)
		return // 当删除的节点度为2时，如果用来取代的红色节点直接返回
	}

	// 2.如果删除的黑色叶子节点 会导致B树节点下溢
	// 如果删除的黑色叶子节点是根节点 直接return
	parent := node.Parent
	if parent == nil {
		return
	}

	// 3.如果删除的黑色叶子节点,兄弟节点是黑色并且兄弟节点时有红色节点
	// 进行旋转操作
	// 旋转之后的中心节点继承 parent 的颜色
	// 旋转之后的左右节点染为 BLAC
	// 4.如果删除的黑色叶子节点,兄弟节点是黑色并且兄弟节点没有红色节点(兄弟节点也是叶子节点)
	// 将 sibling 染成 RED、parent 染成 BLACK 即可修复红黑树性质
	// 如果 parent 是 BLACK
	// 会导致 parent 也下溢
	// 这时只需要把 parent 当做被删除的节点处理即可

	// 5.如果删除的黑色叶子节点,兄弟节点是红色
	// sibling 染成 BLACK，parent 染成 RED，进行旋转
	// 于是又回到 sibling 是 BLACK 的情况
	// // 判断被删除的node是左还是右
	left := parent.Left == nil || node.IsLeftChild()                       //parent.left == null说明当初删除的叶子节点是在左边
	sibling := util.If(left, parent.Right, parent.Left).(*binarytree.Node) //不能使用node.subling() 因为parent的left和right在删除的时候被清空了
	if left {                                                              // 被删除的节点在左边，兄弟节点在右边 (左右是对称的)
		// 兄弟节点是红色
		if isRed(sibling) {
			black(sibling)
			red(parent)
			rb.RotateLeft(parent)
			// 更换兄弟
			sibling = parent.Right
		}
		// 兄弟节点必然是黑色
		if isBlack(sibling.Left) && isBlack(sibling.Right) { // 兄弟节点没有1个红色子节点，父节点要向下跟兄弟节点合并
			parentBlack := isBlack(parent)
			black(parent)
			red(sibling)
			if parentBlack {
				rb.AfterRemove(parent)
			}
		} else { // 兄弟节点至少有1个红色子节点，向兄弟节点借元素
			// 兄弟节点的左边是黑色，兄弟要先旋转
			if isBlack(sibling.Right) {
				rb.RotateRight(sibling)
				sibling = parent.Right
			}
			color(sibling, colorOf(parent))
			black(sibling.Right)
			black(parent)
			rb.RotateLeft(parent)
		}
	} else { // 被删除的节点在右边，兄弟节点在左边
		// 兄弟节点是红色
		if isRed(sibling) {
			black(sibling)
			red(parent)
			rb.RotateRight(parent)
			// 更换兄弟
			sibling = parent.Left
		}
		// 兄弟节点必然是黑色
		if isBlack(sibling.Left) && isBlack(sibling.Right) { // 兄弟节点没有1个红色子节点，父节点要向下跟兄弟节点合并
			parentBlack := isBlack(parent)
			black(parent)
			red(sibling)
			if parentBlack {
				rb.AfterRemove(parent)
			}
		} else { // 兄弟节点至少有1个红色子节点，向兄弟节点借元素
			// 兄弟节点的左边是黑色，兄弟要先旋转
			if isBlack(sibling.Left) {
				rb.RotateLeft(sibling)
				sibling = parent.Left
			}
			color(sibling, colorOf(parent))
			black(sibling.Left)
			black(parent)
			rb.RotateRight(parent)
		}
	}
}

// 染色
func color(node *binarytree.Node, color bool) *binarytree.Node {
	if node == nil {
		return node
	}
	node.Class.(*rbNode).color = color
	return node
}

// 染成红色
func red(node *binarytree.Node) *binarytree.Node {
	return color(node, RED)
}

// 染成黑色
func black(node *binarytree.Node) *binarytree.Node {
	return color(node, BLACK)
}

// 判断节点颜色
func colorOf(node *binarytree.Node) bool {
	// 空值节点 ：叶子节点默认是黑色(红黑树性质3)
	if node == nil {
		return BLACK
	}
	return node.Class.(*rbNode).color
}

// 判断节点是否是黑色
func isBlack(node *binarytree.Node) bool {
	return colorOf(node) == BLACK
}

// 判断节点是否是红色
func isRed(node *binarytree.Node) bool {
	return colorOf(node) == RED
}

type rbNode struct {
	Node  *binarytree.Node
	color bool //默认是红色(建议新添加的节点默认为 RED，这样能够让红黑树的性质尽快满足（性质 1、2、3、5 // 都满足，性质 4 不一定）)
}

func (t *rbNode) ToString() string {
	str := ""
	if t.color == RED {
		str = "R_"
	}
	return str + t.Node.String()
}
func (t *rbNode) ColorOf() bool {
	return !t.color
}
