package printer

import (
	"fmt"
	"reflect"
)

/*
*

	            ┌──800
	        ┌──760
	        │   └──600
	    ┌──540
	    │   └──476
	    │       └──445
	┌──410
	│   └──394

381

	│     ┌──190
	│     │   └──146
	│  ┌──40
	│  │  └──35
	└──12
	   └──9
	*
*/
type inorderPrinter struct {
	Tree        BinaryTreeInfo
	length      int
	rightAppend string
	leftAppend  string
	blankAppend string
	lineAppend  string
}

func initInorderPrinter(tree BinaryTreeInfo) basePrinter {
	t := new(inorderPrinter)
	t.Tree = tree
	t.length = 2
	t.rightAppend = "┌" + repeat("─", t.length)
	t.leftAppend = "└" + repeat("─", t.length)
	t.blankAppend = blank(t.length + 1)
	t.lineAppend = "│" + blank(t.length)
	return t
}

// 打印
func (t *inorderPrinter) Print() {
	print(t.PrintString())
}

// 打印后换行
func (t *inorderPrinter) Println() {
	t.Print()
	println()
}

// 打印字符串
func (t *inorderPrinter) PrintString() string {
	str := t.printTree(t.Tree.RootNode(), "", "", "")
	return str
}

/**
 * 生成node节点的字符串
 * @param nodePrefix node那一行的前缀字符串
 * @param leftPrefix node整棵左子树的前缀字符串
 * @param rightPrefix node整棵右子树的前缀字符串
 * @return
 */
func (t *inorderPrinter) printTree(root interface{}, nodePrefix, leftPrefix, rightPrefix string) string {
	left := typeOf(root, "LNode")
	right := typeOf(root, "RNode")
	str := fmt.Sprintf("%v", typeOf(root, "ToString"))
	length := len(str)
	if length%2 == 0 {
		length--
	}
	length >>= 1

	nodeString := ""

	if left != nil && !reflect.ValueOf(left).IsNil() {
		rightPrefix += blank(length)
		nodeString += t.printTree(right,
			rightPrefix+t.rightAppend,
			rightPrefix+t.lineAppend,
			rightPrefix+t.blankAppend)
	}
	nodeString += nodePrefix + str + "\n"
	if right != nil && !reflect.ValueOf(right).IsNil() {
		leftPrefix += blank(length)
		nodeString += t.printTree(left,
			leftPrefix+t.leftAppend,
			leftPrefix+t.blankAppend,
			leftPrefix+t.lineAppend)
	}
	return nodeString
}
