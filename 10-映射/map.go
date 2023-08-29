package gomap

import (
	"DataStruct_Go/utils"
	"fmt"
	"reflect"
	"strings"
)

type treeMap struct {
	size       int
	root       *treeNode
	comparator utils.Comparable
}

const (
	RED   = false
	BLACK = true
)

func NewTreeMap() utils.Map {
	return &treeMap{}
}

func (t *treeMap) Size() int {
	return t.size
}

func (t *treeMap) IsEmpty() bool {
	return t.size == 0
}

func (t *treeMap) Clear() {
	t.size = 0
	t.root = nil
}

func (t *treeMap) Put(key any, value any) any {
	keyNotNullCheck(key)
	if t.root == nil {
		t.root = createNode(key, value, nil)
		t.size++
		// 新添加节点之后的处理
		t.afterPut(t.root)
		return nil
	}
	// 添加只能是叶子节点，所以必须找到其父节点
	node := t.root
	parent := t.root
	cmp := 0
	for node != nil {
		cmp = t.compare(value, node.value)
		parent = node
		if cmp > 0 {
			node = node.right
		} else if cmp < 0 {
			node = node.left
		} else { // 相等 直接覆盖
			node.key = key
			oldValue := node.value
			node.value = value
			return oldValue
		}
	}
	// 看看插入到父节点的哪个位置
	newNode := createNode(key, value, parent)
	if cmp > 0 {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
	t.size++
	// 新添加节点之后的处理
	t.afterPut(t.root)
	return nil
}

func (t *treeMap) afterPut(node *treeNode) { // 修复性质 4
	parent := node.parent
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
	uncle := parent.sibling()
	grand := red(parent.parent)
	if isRed(uncle) { // 如果uncle节点是红色 【B树节点上溢】
		// 四种uncle点是红色的情况(null<--红-->null<--黑-->null<--红-->null)
		// red(grand);
		black(parent)
		black(uncle)
		// 把祖父节点当做是新添加的节点
		t.afterPut(grand)
		return
	}
	// 如果uncle节点不是红色
	// 四种uncle点是不红色的情况(null<--红-->null<--黑-->(uncle)null、null(uncle)<--黑-->null<--红-->null)
	// 处理方式先染色、然后再进行LL、LR、RR、RL四种旋转
	if parent.isLeftChild() { // L
		if node.isLeftChild() { // LL
			black(parent)
		} else { // LR
			black(node)
			t.rotateLeft(parent)
		}
		t.rotateRight(grand)
	} else { // R
		if node.isLeftChild() { // RL
			black(node)
			t.rotateRight(parent)
		} else { // RR
			black(parent)
		}
		t.rotateLeft(grand)
	}
}

func (t *treeMap) Get(key any) any {
	node := t.node(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *treeMap) Remove(key any) any {
	return t.remove(t.node(key))
}

func (t *treeMap) remove(node *treeNode) any {
	if node == nil {
		return nil
	}
	t.size--
	oldValue := node.value
	if node.hasTwoChildren() { //删除度为2的节点时，找到他的前驱或者后继节点 先交换两值 删除前驱或者后继节点
		// 找到前驱/后继节点
		predecessor := predecessor(node)
		// 用后继节点的值覆盖度为2的节点的值
		predecessor = successor(node)
		node.key = predecessor.key
		node.value = predecessor.value //用前驱或者后继节点的值覆盖度为2的节点的值
		node = predecessor             //使用前驱或者后继节点替代node
	}
	// 删除node节点（node的度必然是1或者0）
	replacement := utils.If(node.left != nil, node.left, node.right).(*treeNode)
	if replacement != nil { // node是度为1的节点并且是根节点
		// 更改parent
		replacement.parent = node.parent
		if node.parent == nil { //删除度为0,直接删除
			t.root = replacement
		} else if node.parent.left == node {
			node.parent.left = replacement
		} else {
			node.parent.right = replacement
		}
		// 删除节点之后的处理
		t.afterRemove(replacement) //注意区分avl树，此处传用以取代被删除节点的子节点，而不是node，因为此时replacement必然为红色，只需将其染黑即可直接返回
	} else if node.parent == nil { // node是叶子节点并且是根节点
		t.root = nil
	} else { //删除度为0
		if node == node.parent.left {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		// 删除节点之后的处理
		t.afterRemove(node)
	}
	return oldValue
}

func (t *treeMap) afterRemove(node *treeNode) {
	// 1.如果删除的时候红色 或者 用以取代删除节点的子节点是红色 直接返回
	if isRed(node) {
		black(node)
		return // 当删除的节点度为2时，如果用来取代的红色节点直接返回
	}

	// 2.如果删除的黑色叶子节点 会导致B树节点下溢
	// 如果删除的黑色叶子节点是根节点 直接return
	parent := node.parent
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
	left := parent.left == nil || node.isLeftChild()                 //parent.left == null说明当初删除的叶子节点是在左边
	sibling := utils.If(left, parent.right, parent.left).(*treeNode) //不能使用node.subling() 因为parent的left和right在删除的时候被清空了
	if left {                                                        // 被删除的节点在左边，兄弟节点在右边 (左右是对称的)
		// 兄弟节点是红色
		if isRed(sibling) {
			black(sibling)
			red(parent)
			t.rotateLeft(parent)
			// 更换兄弟
			sibling = parent.right
		}
		// 兄弟节点必然是黑色
		if isBlack(sibling.left) && isBlack(sibling.right) { // 兄弟节点没有1个红色子节点，父节点要向下跟兄弟节点合并
			parentBlack := isBlack(parent)
			black(parent)
			red(sibling)
			if parentBlack {
				t.afterRemove(parent)
			}
		} else { // 兄弟节点至少有1个红色子节点，向兄弟节点借元素
			// 兄弟节点的左边是黑色，兄弟要先旋转
			if isBlack(sibling.right) {
				t.rotateRight(sibling)
				sibling = parent.right
			}
			color(sibling, colorOf(parent))
			black(sibling.right)
			black(parent)
			t.rotateLeft(parent)
		}
	} else { // 被删除的节点在右边，兄弟节点在左边
		// 兄弟节点是红色
		if isRed(sibling) {
			black(sibling)
			red(parent)
			t.rotateRight(parent)
			// 更换兄弟
			sibling = parent.left
		}
		// 兄弟节点必然是黑色
		if isBlack(sibling.left) && isBlack(sibling.right) { // 兄弟节点没有1个红色子节点，父节点要向下跟兄弟节点合并
			parentBlack := isBlack(parent)
			black(parent)
			red(sibling)
			if parentBlack {
				t.afterRemove(parent)
			}
		} else { // 兄弟节点至少有1个红色子节点，向兄弟节点借元素
			// 兄弟节点的左边是黑色，兄弟要先旋转
			if isBlack(sibling.left) {
				t.rotateLeft(sibling)
				sibling = parent.left
			}
			color(sibling, colorOf(parent))
			black(sibling.left)
			black(parent)
			t.rotateRight(parent)
		}
	}
}

func (t *treeMap) ContainsKey(key any) bool {
	return t.node(key) != nil
}

func (t *treeMap) ContainsValue(value any) bool {
	if t.root == nil {
		return false
	}
	queue := []*treeNode{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if utils.Equal(value, node.value) {
			return true
		}
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return false
}

func (t *treeMap) Traversal(visitors ...utils.Visitor) {
	visitor := utils.NewVisitor()
	if len(visitors) == 1 {
		visitor = visitors[0]
	}
	traversal(t.root, visitor)
}

func traversal(node *treeNode, visitor utils.Visitor) {
	if node == nil || visitor.Stop {
		return
	}
	traversal(node.left, visitor)
	if visitor.Stop {
		return
	}
	visitor.Visit(fmt.Sprintf("[%v:%v]", node.key, node.value))
	traversal(node.right, visitor)
}

// predecessor 前驱节点
func predecessor(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}

	// 前驱节点在左子树当中（left.right.right.right....）
	p := node.left
	if p != nil {
		for p.right != nil {
			p = p.right
		}
		return p
	}

	// 从父节点、祖父节点中寻找前驱节点
	for node.parent != nil && node == node.parent.left {
		node = node.parent
	}

	// node.parent == null
	// node == node.parent.right
	return node.parent
}

// 后继节点
func successor(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}

	// 后继节点在左子树当中（right.left.left.left....）
	p := node.right
	if p != nil {
		for p.left != nil {
			p = p.left
		}
		return p
	}

	// 从父节点、祖父节点中寻找前驱节点
	for node.parent != nil && node == node.parent.right {
		node = node.parent
	}

	return node.parent
}

// 根据key值找Node
func (t *treeMap) node(key any) *treeNode {
	node := t.root
	for node != nil {
		cmp := t.compare(key, node.key)
		if cmp == 0 {
			return node
		}
		if cmp > 0 {
			node = node.right
		} else {
			node = node.left
		}
	}
	return nil
}

// 返回值等于0，代表e1和e2相等；返回值大于0，代表e1大于e2；返回值小于于0，代表e1小于e2
func (t *treeMap) compare(e1, e2 any) int {
	//if t.Comparator != nil {
	//	return t.Comparator(e1, e2)
	//} //todo 待优化
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

// key 不能为空
func keyNotNullCheck(key any) {
	if key == nil {
		panic("key must not be null")
	}
}

// 左旋转
func (t *treeMap) rotateLeft(grand *treeNode) {
	parent := grand.right
	child := parent.left
	grand.right = child
	parent.left = grand
	t.afterRotate(grand, parent, child)
}

// 右旋转
func (t *treeMap) rotateRight(grand *treeNode) {
	parent := grand.left
	child := parent.right
	grand.left = child
	parent.right = grand
	t.afterRotate(grand, parent, child)
}

// 旋转之后更新节点的parent和高度
func (t *treeMap) afterRotate(grand, parent, child *treeNode) {
	// 让parent称为子树的根节点
	parent.parent = grand.parent
	//更新grand.parent的左右子树
	if grand.isLeftChild() {
		grand.parent.left = parent
	} else if grand.isRightChild() {
		grand.parent.right = parent
	} else { // grand是root节点
		t.root = parent
	}

	// 更新child的parent
	if child != nil {
		child.parent = grand
	}
	// 更新grand的parent
	grand.parent = parent
}

// 染色
func color(node *treeNode, color bool) *treeNode {
	if node == nil {
		return node
	}
	node.color = color
	return node
}

// 染成红色
func red(node *treeNode) *treeNode {
	return color(node, RED)
}

// 染成黑色
func black(node *treeNode) *treeNode {
	return color(node, BLACK)
}

// 判断节点颜色
func colorOf(node *treeNode) bool {
	// 空值节点 ：叶子节点默认是黑色(红黑树性质3)
	if node == nil {
		return BLACK
	}
	return node.color
}

// 判断节点是否是黑色
func isBlack(node *treeNode) bool {
	return colorOf(node) == BLACK
}

// 判断节点是否是红色
func isRed(node *treeNode) bool {
	return colorOf(node) == RED
}

type treeNode struct {
	color bool // 默认是红色(建议新添加的节点默认为 RED，这样能够让红黑树的性质尽快满足（性质 1、2、3、5
	// 都满足，性质 4 不一定）)
	key    any
	value  any
	parent *treeNode
	left   *treeNode
	right  *treeNode
}

func createNode(key any, value any, parent *treeNode) *treeNode {
	return &treeNode{
		color:  RED,
		key:    key,
		value:  value,
		parent: parent,
	}
}

func (n *treeNode) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *treeNode) hasTwoChildren() bool {
	return n.left != nil && n.right != nil
}

func (n *treeNode) isLeftChild() bool {
	return n.parent != nil && n.parent.left == n
}

func (n *treeNode) isRightChild() bool {
	return n.parent != nil && n.parent.right == n
}

// 返回兄弟节点
func (n *treeNode) sibling() *treeNode {
	if n.isLeftChild() {
		return n.parent.right
	}
	if n.isRightChild() {
		return n.parent.left
	}
	return nil
}
