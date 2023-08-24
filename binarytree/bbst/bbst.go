package bbst

import (
	"DataStruct_Go/binarytree"
	"DataStruct_Go/binarytree/bst"
)

// 平衡二叉搜索树
type BBST struct {
	bst.BST
}

// 回调函数
type CallBack interface {
	CallFunc(*binarytree.Node, *binarytree.Node, ...*binarytree.Node)
}

func (bbst *BBST) Rotate(r, b, c, d, e, f *binarytree.Node, callBack CallBack) { //r 子树的根节点
	// 让d成为这棵子树的根节点
	d.Parent = r.Parent
	if r.IsLeftChild() {
		r.Parent.Left = d
	} else if r.IsRightChild() {
		r.Parent.Right = d
	} else {
		bbst.Root = d
	}

	//b-c
	b.Right = c
	if c != nil {
		c.Parent = b
	}

	// e-f
	f.Left = e
	if e != nil {
		e.Parent = f
	}

	// b-d-f
	d.Left = b
	d.Right = f
	b.Parent = d
	f.Parent = d

	// 节点旋转之后的处理
	callBack.CallFunc(b, f, d)
}

// 左旋转
func (bbst *BBST) RotateLeft(grand *binarytree.Node, callBack ...CallBack) {
	parent := grand.Right
	child := parent.Left
	grand.Right = child
	parent.Left = grand
	bbst.AfterRotate(grand, parent, child, callBack...)
}

// 右旋转
func (bbst *BBST) RotateRight(grand *binarytree.Node, callBack ...CallBack) {
	parent := grand.Left
	child := parent.Right
	grand.Left = child
	parent.Right = grand
	bbst.AfterRotate(grand, parent, child, callBack...)
}

// 旋转之后更新节点的parent和高度
func (bbst *BBST) AfterRotate(grand, parent, child *binarytree.Node, callBack ...CallBack) {
	// 让parent称为子树的根节点
	parent.Parent = grand.Parent
	//更新grand.parent的左右子树
	if grand.IsLeftChild() {
		grand.Parent.Left = parent
	} else if grand.IsRightChild() {
		grand.Parent.Right = parent
	} else { // grand是root节点
		bbst.Root = parent
	}

	// 更新child的parent
	if child != nil {
		child.Parent = grand
	}
	// 更新grand的parent
	grand.Parent = parent

	// 节点旋转之后的处理
	if len(callBack) == 1 {
		callBack[0].CallFunc(grand, parent)
	}
}
