package printer

import (
	"bytes"
	container "container/list"
	"fmt"
	"math"
	"reflect"
	"strings"
)

/*
*

	┌───381────┐
	│          │

┌─12─┐     ┌─410─┐
│    │     │     │
9  ┌─40─┐ 394 ┌─540─┐

	  │    │     │     │
	 35 ┌─190 ┌─476 ┌─760─┐
	    │     │     │     │
	   146   445   600   800

	*
*/
const (
	minSpace     = 1 //节点之间允许的最小间距（最小只能填1）
	topLineSpace = 1 //顶部符号距离父节点的最小距离（最小能填0）
)

type levelOrderPrinter struct {
	Tree     BinaryTreeInfo
	root     *node
	minX     int
	maxWidth int
}

func initLevelOrderPrinter(tree BinaryTreeInfo) basePrinter {
	t := new(levelOrderPrinter)
	t.Tree = tree
	t.root = initNodeOpetaion(tree.RootNode())
	t.maxWidth = t.root.width
	return t
}

type list []*node // 类型别名

// 打印
func (t *levelOrderPrinter) Print() {
	//print(t.PrintString())
	fmt.Print(t.PrintString())
}

// 打印后换行
func (t *levelOrderPrinter) Println() {
	t.Print()
	//println()
	fmt.Println()
}

// 打印字符串
func (t *levelOrderPrinter) PrintString() string {
	// nodes用来存放所有的节点
	nodes := make([]list, 0)
	nodes = t.fillNodes(nodes)
	t.cleanNodes(nodes)
	t.compressNodes(nodes)
	nodes = t.addLineNodes(nodes)

	rowCount := len(nodes)

	// 构建字符串
	var buffer bytes.Buffer
	for i := 0; i < rowCount; i++ {
		if i != 0 {
			buffer.WriteString("\n")
		}

		rowNodes := nodes[i]
		rowSb := ""
		count := 0
		for _, node := range rowNodes {
			leftSpace := node.x - len([]rune(rowSb)) - t.minX + 9*count
			if strings.HasSuffix(node.str, "[0m") {
				count++
			}
			rowSb += blank(leftSpace)
			rowSb += node.str
		}
		buffer.WriteString(rowSb)
	}
	return buffer.String()
}

// 以满二叉树的形式填充节点
func (t *levelOrderPrinter) fillNodes(nodes []list) []list {
	if nodes == nil {
		return nil
	}
	// 第一行
	firstRowNodes := make(list, 0)
	firstRowNodes = append(firstRowNodes, t.root)
	nodes = append(nodes, firstRowNodes)
	// 其他行
	for {
		preRowNodes := nodes[len(nodes)-1]
		rowNodes := make(list, 0)
		notNull := false
		for _, node := range preRowNodes {
			if node == nil {
				rowNodes = append(rowNodes, nil)
				rowNodes = append(rowNodes, nil)
			} else {
				//left := t.addNode(rowNodes, t.Tree.Left(node.btNode))
				left := typeOf(node.btNode, "LNode")
				if left == nil || reflect.ValueOf(left).IsNil() {
					rowNodes = append(rowNodes, nil)
				} else {
					leftNode := t.addNode(left)
					rowNodes = append(rowNodes, leftNode)
					if leftNode != nil {
						node.left = leftNode
						leftNode.parent = node
						notNull = true
					}
				}
				//right := t.addNode(rowNodes, t.Tree.Right(node.btNode))\
				right := typeOf(node.btNode, "RNode")
				if right == nil || reflect.ValueOf(right).IsNil() {
					rowNodes = append(rowNodes, nil)
				} else {
					rightNode := t.addNode(right)
					rowNodes = append(rowNodes, rightNode)
					if rightNode != nil {
						node.right = rightNode
						rightNode.parent = node
						notNull = true
					}
				}
			}
		}
		// 全是null，就退出
		if !notNull {
			break
		}
		nodes = append(nodes, rowNodes)
	}
	return nodes
}

// 添加一个元素节点
func (t *levelOrderPrinter) addNode(btNode interface{}) *node {
	node := new(node)
	if btNode != nil {
		node = initNodeOpetaion(btNode)
		if t.maxWidth < node.width {
			t.maxWidth = node.width
		}
	} else {
		return nil
	}
	return node
}

// 删除全部null、更新节点的坐标
func (t *levelOrderPrinter) cleanNodes(nodes []list) {
	if nodes == nil {
		return
	}

	//总行数
	rowCount := len(nodes)
	if rowCount < 2 {
		return
	}

	// 最后一行的节点数量
	lastRowNodeCount := len(nodes[rowCount-1])

	// 每个节点之间的间距
	nodeSpace := t.maxWidth + 2

	// 最后一行的长度
	lastRowLength := lastRowNodeCount*t.maxWidth + nodeSpace*(lastRowNodeCount-1)

	for i := 0; i < rowCount; i++ {
		rowNodes := nodes[i]

		rowNodeCount := len(rowNodes)
		// 节点左右两边的间距
		allSpace := lastRowLength - (rowNodeCount-1)*nodeSpace
		cornerSpace := allSpace/rowNodeCount - t.maxWidth
		cornerSpace >>= 1

		var rowLength = 0
		for j := 0; j < rowNodeCount; j++ {
			if j != 0 {
				// 每个节点之间的间距
				rowLength += nodeSpace
			}
			rowLength += cornerSpace
			node := rowNodes[j]
			if node != nil {
				// 居中（由于奇偶数的问题，可能有1个符号的误差）
				deltaX := (t.maxWidth - node.width) >> 1
				node.x = rowLength + deltaX
				node.y = i
			}
			rowLength += t.maxWidth
			rowLength += cornerSpace
		}

		// 删除所有的null
		for i := 0; i < len(rowNodes); i++ {
			if rowNodes[i] == nil {
				rowNodes = append(rowNodes[:i], rowNodes[i+1:]...)
				i--
			}
		}
		nodes[i] = rowNodes
	}
}

// 压缩空格
func (t *levelOrderPrinter) compressNodes(nodes []list) {
	if nodes == nil {
		return
	}

	//总行数
	rowCount := len(nodes)
	if rowCount < 2 {
		return
	}

	for i := rowCount - 2; i >= 0; i-- {
		rowNodes := nodes[i]
		for _, node := range rowNodes {
			left := node.left
			right := node.right
			if left == nil && right == nil {
				continue
			}
			if left != nil && right != nil {
				// 让左右节点对称
				node.balance(left, right)

				// left和right之间可以挪动的最小间距
				leftEmpty := node.leftBoundEmptyLength()
				rightEmpty := node.rightBoundEmptyLength()
				empty := leftEmpty
				if leftEmpty > rightEmpty {
					empty = rightEmpty
				}
				if empty > (right.x-left.rightX())>>1 {
					empty = (right.x - left.rightX()) >> 1
				}

				// left、right的子节点之间可以挪动的最小间距
				space := left.minLevelSpaceToRight(right) - minSpace
				if empty > (space >> 1) {
					space = space >> 1
				} else {
					space = empty
				}
				// left、right往中间挪动
				if space > 0 {
					left.translateX(space)
					right.translateX(-space)
				}

				// 继续挪动
				space = left.minLevelSpaceToRight(right) - minSpace
				if space < 1 {
					continue
				}

				// 可以继续挪动的间距
				leftEmpty = node.leftBoundEmptyLength()
				rightEmpty = node.rightBoundEmptyLength()
				if leftEmpty < 1 && rightEmpty < 1 {
					continue
				}

				if leftEmpty > rightEmpty {
					deltaX := leftEmpty
					if leftEmpty > space {
						deltaX = space
					}
					left.translateX(deltaX)
				} else {
					deltaX := rightEmpty
					if rightEmpty > space {
						deltaX = space
					}
					right.translateX(-deltaX)
				}
			} else if left != nil {
				left.translateX(node.leftBoundEmptyLength())
			} else if right != nil { // right != nil
				right.translateX(-node.rightBoundEmptyLength())
			}
		}
	}
}

// 添加行
func (t *levelOrderPrinter) addLineNodes(nodes []list) []list {

	rowCount := len(nodes)
	if rowCount < 2 {
		return nodes
	}
	newNodes := make([]list, 0)

	t.minX = t.root.x

	for i := 0; i < rowCount; i++ {
		rowNodes := nodes[i]
		if i == rowCount-1 {
			newNodes = append(newNodes, rowNodes)
			continue
		}

		newRowNodes := make(list, 0)
		lineNodes := make(list, 0)
		for _, node := range rowNodes {
			newRowNodes, lineNodes, _ = t.addLineNode(newRowNodes, lineNodes, node, node.left)
			newRowNodes = append(newRowNodes, node)
			newRowNodes, lineNodes, _ = t.addLineNode(newRowNodes, lineNodes, node, node.right)
		}
		newNodes = append(newNodes, newRowNodes)
		newNodes = append(newNodes, lineNodes)
	}
	nodes = nil
	return newNodes
}

func (t *levelOrderPrinter) addLineNode(curRow, nextRow list, parent, child *node) (list, list, *node) {
	if child == nil {
		return curRow, nextRow, nil
	}

	var top *node
	var topX = child.topLineX()
	if child == parent.left {
		top = initNode("┌")
		curRow = append(curRow, top)
		for x := topX + 1; x < parent.x; x++ {
			curRow = addXLineNode(curRow, parent, x)
		}
	} else {
		for x := parent.rightX(); x < topX; x++ {
			curRow = addXLineNode(curRow, parent, x)
		}
		top = initNode("┐")
		curRow = append(curRow, top)
	}

	// 坐标
	top.x = topX
	top.y = parent.y
	child.y = parent.y + 2
	if t.minX > child.x {
		t.minX = child.x
	}

	// 竖线
	bottom := initNode("│")
	bottom.x = topX
	bottom.y = parent.y + 1
	nextRow = append(nextRow, bottom)

	return curRow, nextRow, top
}

func addXLineNode(curRow list, parent *node, x int) list {
	line := initNode("─")
	line.x = x
	line.y = parent.y
	curRow = append(curRow, line)
	return curRow
}

type node struct {
	btNode     interface{}
	left       *node
	right      *node
	parent     *node
	x          int //首字符的位置
	y          int
	treeHeight int
	str        string
	width      int
	color      bool
}

func initNode(str string) *node {
	n := new(node)
	n.init(str)
	return n
}

func initNodeOpetaion(btNode interface{}) *node {
	n := new(node)
	n.color = typeOf(btNode, "ColorOf").(bool)
	n.btNode = btNode
	n.init(fmt.Sprintf("%v", typeOf(btNode, "ToString")))
	return n
}

func (n *node) init(str string) {
	if len(str) == 0 {
		str = " "
	}
	n.width = len([]rune(str))
	if strings.HasSuffix(str, "[0m") {
		n.width = n.width - 9
	}
	if n.color {
		str = fmt.Sprintf("%c[%dm%s%c[0m", 0x1B, 31, str, 0x1B)
	}
	n.str = str
}

// 让left和right基于this对称
func (n *node) balance(left, right *node) {
	if left == nil || right == nil {
		return
	}
	// 【left的尾字符】与【this的首字符】之间的间距
	deltaLeft := n.x - left.rightX()
	// 【this的尾字符】与【this的首字符】之间的间距
	deltaRight := right.x - n.rightX()

	delta := deltaLeft
	if deltaLeft < deltaRight {
		delta = deltaRight
	}

	newRightX := n.rightX() + delta

	right.translateX(newRightX - right.x)

	newLeftX := n.x - delta - left.width
	left.translateX(newLeftX - left.x)
}

// 尾字符的下一个位置
func (n *node) rightX() int {
	return n.x + n.width
}

func (n *node) translateX(deltaX int) {
	if deltaX == 0 {
		return
	}
	n.x += deltaX

	// 如果是LineNode
	if n.btNode == nil {
		return
	}

	if n.left != nil {
		n.left.translateX(deltaX)
	}
	if n.right != nil {
		n.right.translateX(deltaX)
	}
}

// 左边界可以清空的长度
func (n *node) leftBoundEmptyLength() int {
	leftBound := n.x
	if n.left != nil {
		leftBound = n.left.topLineX()
	}
	return n.x - leftBound - 1 - topLineSpace
}

// 右边界可以清空的长度
func (n *node) rightBoundEmptyLength() int {
	if n.right != nil {
		return n.right.topLineX() - n.rightX() - topLineSpace
	}
	return -1 - topLineSpace
}

// 顶部方向字符的X（极其重要）
func (n *node) topLineX() int {
	// 宽度的一半
	delta := n.width
	if delta%2 == 0 {
		delta--
	}
	delta >>= 1

	if n.parent != nil && n == n.parent.left {
		return n.rightX() - 1 - delta
	} else {
		return n.x + delta
	}
}

// 和右节点之间的最小层级距离
func (n *node) minLevelSpaceToRight(right *node) int {
	leftHeight := treeHeight(n)
	rightHeight := treeHeight(right)
	minSpace := math.MaxInt32
	for i := 0; i < leftHeight && i < rightHeight; i++ {
		space := right.levelInfo(i, 1) - n.levelInfo(i, 2)
		if minSpace > space {
			minSpace = space
		}
	}
	return minSpace
}

func (n *node) levelInfo(level, derection int) int {
	if level < 0 {
		return 0
	}
	levelY := n.y + level
	if level >= treeHeight(n) {
		return 0
	}
	list := make(list, 0)
	queue := container.New()
	queue.PushBack(n)
	// 层序遍历找出第level行的所有节点
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*node)
		if levelY == node.y {
			list = append(list, node)
		} else if node.y > levelY {
			break
		}

		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
	if derection == 1 {
		leftX := list[0].x
		if list[0].left != nil {
			leftX = list[0].left.topLineX()
		}
		return leftX
	} else {
		rightX := list[len(list)-1].rightX()
		if list[len(list)-1].right != nil {
			rightX = list[len(list)-1].right.topLineX() + 1
		}
		return rightX
	}
}

// 计算树的高度
func treeHeight(node *node) int {
	if node == nil {
		return 0
	}
	if node.treeHeight != 0 {
		return node.treeHeight
	}
	height := treeHeight(node.left)
	rightHeight := treeHeight(node.right)
	if height < rightHeight {
		height = rightHeight
	}
	node.treeHeight = height + 1
	return node.treeHeight
}
