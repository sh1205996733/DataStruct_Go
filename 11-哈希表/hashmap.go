package hashmap

import (
	"DataStruct_Go/utils"
	"fmt"
	"time"
)

type hashmap struct {
	size  int
	table []*hashNode
}

const (
	RED   = false
	BLACK = true
)
const DEFAULT_CAPACITY = 1 << 4
const DEFAULT_LOAD_FACTOR = 0.75

func NewHashmap() utils.Map {
	return &hashmap{
		table: make([]*hashNode, DEFAULT_CAPACITY),
	}
}

func (h *hashmap) Size() int {
	return h.size
}

func (h *hashmap) IsEmpty() bool {
	return h.size == 0
}

func (h *hashmap) Clear() {
	h.size = 0
	for i := 0; i < len(h.table); i++ {
		h.table[i] = nil
	}
}

func (h *hashmap) Put(key any, value any) any {
	h.resize()
	index := h.index(key)
	// 取出index位置的红黑树根节点
	root := h.table[index]
	// 添加第一个节点
	if root == nil {
		root = createNode(key, value, nil)
		h.table[index] = root
		h.size++
		// 新添加节点之后的处理
		h.fixAfterPut(root)
		return nil
	}

	// 添加的不是第一个节点
	// 找到父节点
	parent, node := root, root
	cmp := 0
	k1 := key
	h1 := h.hash(k1)
	var result *hashNode
	var searched bool // 是否已经搜索过这个key
	for node != nil {
		parent = node
		k2 := node.key
		h2 := node.hash
		if h1 > h2 {
			cmp = 1
		} else if h1 < h2 {
			cmp = -1
		} else if k1 == k2 { //  Objects.equals(k1, k2)
			cmp = 0
		} else if k1 != nil && k2 != nil {
			//cmp = ((Comparable) k1).compareTo(k2)) != 0
		} else if searched { // 已经扫描了
			//cmp = System.identityHashCode(k1) - System.identityHashCode(k2);
			cmp = time.Now().Second()
		} else { // searched == false; 还没有扫描，然后再根据内存地址大小决定左右
			l := getNode(node.left, k1)
			r := getNode(node.right, k1)
			if (node.left != nil && l != nil) || (node.right != nil && r != nil) {
				if node.left != nil {
					result = l
				}
				if node.right != nil {
					result = r
				}
				// 已经存在这个key
				node = result
				cmp = 0
			} else { // 不存在这个key
				searched = true
				cmp = time.Now().Second()
				// cmp = System.identityHashCode(k1) - System.identityHashCode(k2);
			}
		}

		if cmp > 0 {
			node = node.right
		} else if cmp < 0 {
			node = node.left
		} else { // 相等
			oldValue := node.value
			node.key = key
			node.value = value
			node.hash = h1
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
	h.size++

	// 新添加节点之后的处理
	h.fixAfterPut(newNode)
	return nil
}

func (h *hashmap) resize() {
	//和containsValue一样遍历每一个节点
	// 装填因子 <= 0.75
	if float64(h.size/len(h.table)) <= DEFAULT_LOAD_FACTOR {
		return
	}
	oldTable := h.table
	h.table = make([]*hashNode, len(oldTable)<<1) //新数组扩大二倍

	var queue []*hashNode
	for i := 0; i < len(oldTable); i++ {
		if oldTable[i] == nil {
			continue
		}
		queue = append(queue, oldTable[i])
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			// 挪动代码得放到最后面
			h.moveNode(node)
		}
	}
}

// 移动节点 和put处理逻辑一样（旧的节点当新的节点添加到新的数组中）,但是不用维护size
func (h *hashmap) moveNode(newNode *hashNode) {
	// 重置
	newNode.parent = nil
	newNode.left = nil
	newNode.right = nil
	newNode.color = RED

	index := h.index(newNode)
	// 取出index位置的红黑树根节点
	root := h.table[index]
	// 添加第一个节点
	if root == nil {
		root = newNode
		h.table[index] = root
		// 新添加节点之后的处理
		h.fixAfterPut(root)
		return
	}

	// 添加的不是第一个节点
	// 找到父节点
	parent, node := root, root
	var cmp = 0
	k1 := newNode.key
	h1 := newNode.hash
	for node != nil {
		parent = node
		k2 := node.key
		h2 := node.hash
		if h1 > h2 {
			cmp = 1
		} else if h1 < h2 {
			cmp = -1
		} else if k1 != nil && k2 != nil {
			//cmp >0
		} else {
			//cmp = System.identityHashCode(k1) - System.identityHashCode(k2);
			cmp = time.Now().Second()
		}

		if cmp > 0 {
			node = node.right
		} else if cmp < 0 {
			node = node.left
		}
	}

	// 看看插入到父节点的哪个位置
	newNode.parent = parent
	if cmp > 0 {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
	// 新添加节点之后的处理
	h.fixAfterPut(newNode)
}
func (h *hashmap) fixAfterPut(node *hashNode) {
	parent := node.parent

	// 添加的是根节点 或者 上溢到达了根节点
	if parent == nil {
		black(node)
		return
	}

	// 如果父节点是黑色，直接返回
	if isBlack(parent) {
		return
	}

	// 叔父节点
	uncle := parent.sibling()
	// 祖父节点
	grand := red(parent.parent)
	if isRed(uncle) { // 叔父节点是红色【B树节点上溢】
		black(parent)
		black(uncle)
		// 把祖父节点当做是新添加的节点
		h.fixAfterPut(grand)
		return
	}

	// 叔父节点不是红色
	if parent.isLeftChild() { // L
		if node.isLeftChild() { // LL
			black(parent)
		} else { // LR
			black(node)
			h.rotateLeft(parent)
		}
		h.rotateRight(grand)
	} else { // R
		if node.isLeftChild() { // RL
			black(node)
			h.rotateRight(parent)
		} else { // RR
			black(parent)
		}
		h.rotateLeft(grand)
	}
}

func (h *hashmap) Get(key any) any {
	node := h.node(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (h *hashmap) Remove(key any) any {
	return h.removeNode(h.node(key))
}

func (h *hashmap) removeNode(node *hashNode) any {
	if node == nil {
		return nil
	}
	willNode := node
	h.size--

	oldValue := node.value

	if node.hasTwoChildren() { // 度为2的节点
		// 找到后继节点
		s := successor(node)
		// 用后继节点的值覆盖度为2的节点的值
		node.key = s.key
		node.value = s.value
		node.hash = s.hash
		// 删除后继节点
		node = s
	}

	// 删除node节点（node的度必然是1或者0）
	replacement := node.left
	if node.left == nil {
		replacement = node.right
	}
	index := h.index(node)

	if replacement != nil { // node是度为1的节点
		// 更改parent
		replacement.parent = node.parent
		// 更改parent的left、right的指向
		if node.parent == nil { // node是度为1的节点并且是根节点
			h.table[index] = replacement
		} else if node == node.parent.left {
			node.parent.left = replacement
		} else { // node == node.parent.right
			node.parent.right = replacement
		}

		// 删除节点之后的处理
		h.fixAfterRemove(replacement)
	} else if node.parent == nil { // node是叶子节点并且是根节点
		h.table[index] = nil
	} else { // node是叶子节点，但不是根节点
		if node == node.parent.left {
			node.parent.left = nil
		} else { // node == node.parent.right
			node.parent.right = nil
		}

		// 删除节点之后的处理
		h.fixAfterRemove(node)
	}

	// 交给子类去处理
	h.afterRemove(willNode, node)

	return oldValue
}

func (h *hashmap) fixAfterRemove(node *hashNode) {
	// 如果删除的节点是红色
	// 或者 用以取代删除节点的子节点是红色
	if isRed(node) {
		black(node)
		return
	}

	parent := node.parent
	if parent == nil {
		return
	}

	// 删除的是黑色叶子节点【下溢】
	// 判断被删除的node是左还是右
	left := parent.left == nil || node.isLeftChild()
	sibling := parent.right
	if left {
		sibling = parent.left
	}
	if left { // 被删除的节点在左边，兄弟节点在右边
		if isRed(sibling) { // 兄弟节点是红色
			black(sibling)
			red(parent)
			h.rotateLeft(parent)
			// 更换兄弟
			sibling = parent.right
		}

		// 兄弟节点必然是黑色
		if isBlack(sibling.left) && isBlack(sibling.right) {
			// 兄弟节点没有1个红色子节点，父节点要向下跟兄弟节点合并
			parentBlack := isBlack(parent)
			black(parent)
			red(sibling)
			if parentBlack {
				h.fixAfterRemove(parent)
			}
		} else { // 兄弟节点至少有1个红色子节点，向兄弟节点借元素
			// 兄弟节点的左边是黑色，兄弟要先旋转
			if isBlack(sibling.right) {
				h.rotateRight(sibling)
				sibling = parent.right
			}

			color(sibling, colorOf(parent))
			black(sibling.right)
			black(parent)
			h.rotateLeft(parent)
		}
	} else { // 被删除的节点在右边，兄弟节点在左边
		if isRed(sibling) { // 兄弟节点是红色
			black(sibling)
			red(parent)
			h.rotateRight(parent)
			// 更换兄弟
			sibling = parent.left
		}

		// 兄弟节点必然是黑色
		if isBlack(sibling.left) && isBlack(sibling.right) {
			// 兄弟节点没有1个红色子节点，父节点要向下跟兄弟节点合并
			parentBlack := isBlack(parent)
			black(parent)
			red(sibling)
			if parentBlack {
				h.fixAfterRemove(parent)
			}
		} else { // 兄弟节点至少有1个红色子节点，向兄弟节点借元素
			// 兄弟节点的左边是黑色，兄弟要先旋转
			if isBlack(sibling.left) {
				h.rotateLeft(sibling)
				sibling = parent.left
			}

			color(sibling, colorOf(parent))
			black(sibling.left)
			black(parent)
			h.rotateRight(parent)
		}
	}
}
func (h *hashmap) afterRemove(willNode *hashNode, node *hashNode) {

}

func (h *hashmap) ContainsKey(key any) bool {
	return h.node(key) == nil
}

func (h *hashmap) ContainsValue(value any) bool {
	if h.size == 0 {
		return false
	}
	var queue []*hashNode
	for i := range h.table {
		if h.table[i] == nil {
			continue
		}
		queue = append(queue, h.table[i])
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
	}
	return false
}

func (h *hashmap) Traversal(visitors ...utils.Visitor) {
	if h.size == 0 {
		return
	}
	visitor := utils.NewVisitor()
	if len(visitors) == 1 {
		visitor = visitors[0]
	}
	var queue []*hashNode
	for i := range h.table {
		if h.table[i] == nil {
			continue
		}
		queue = append(queue, h.table[i])
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if visitor.Visit(fmt.Sprintf("[%v:%v]", node.key, node.value)) {
				return
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
}

func (h *hashmap) node(key any) *hashNode {
	root := h.table[h.index(key)]
	if root == nil {
		return nil
	}
	return getNode(root, key)
}

// 后继节点
func successor(node *hashNode) *hashNode {
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

func getNode(node *hashNode, key any) *hashNode {
	h1 := utils.Hash(key)
	// 存储查找结果
	var result *hashNode
	cmp := 0
	for node != nil {
		k2 := node.key
		h2 := node.hash
		// 先比较哈希值
		if h1 > h2 {
			node = node.right
		} else if h1 < h2 {
			node = node.left
		} else if key == k2 {
			return node
		} else if key != nil && k2 != nil {
			// todo
			//cmp = ((Comparable) k1).compareTo(k2)) != 0
			if cmp != 0 {
				node = node.left
				if cmp > 0 {
					node = node.right
				}
			}
		} else if node.right != nil {
			result = getNode(node.right, key)
			if result != nil {
				return result
			}
		} else { // 只能往左边找
			node = node.left
		}
	}
	return nil
}

// 根据key生成对应的索引（在桶数组中的位置）
func (h *hashmap) index(key any) int {
	return h.hash(key) & (len(h.table) - 1)
}

// 根据key生成对应的hash
func (h *hashmap) hash(key any) int {
	if key == nil {
		return 0
	}
	hashCode := utils.Hash(key)
	return hashCode ^ (hashCode >> 16)
}

// 左旋转
func (h *hashmap) rotateLeft(grand *hashNode) {
	parent := grand.right
	child := parent.left
	grand.right = child
	parent.left = grand
	h.afterRotate(grand, parent, child)
}

// 右旋转
func (h *hashmap) rotateRight(grand *hashNode) {
	parent := grand.left
	child := parent.right
	grand.left = child
	parent.right = grand
	h.afterRotate(grand, parent, child)
}

// 旋转之后更新节点的parent和高度
func (h *hashmap) afterRotate(grand, parent, child *hashNode) {
	// 让parent称为子树的根节点
	parent.parent = grand.parent
	//更新grand.parent的左右子树
	if grand.isLeftChild() {
		grand.parent.left = parent
	} else if grand.isRightChild() {
		grand.parent.right = parent
	} else { // grand是root节点
		h.table[h.index(grand)] = parent
	}

	// 更新child的parent
	if child != nil {
		child.parent = grand
	}
	// 更新grand的parent
	grand.parent = parent
}

// 染色
func color(node *hashNode, color bool) *hashNode {
	if node == nil {
		return node
	}
	node.color = color
	return node
}

// 染成红色
func red(node *hashNode) *hashNode {
	return color(node, RED)
}

// 染成黑色
func black(node *hashNode) *hashNode {
	return color(node, BLACK)
}

// 判断节点颜色
func colorOf(node *hashNode) bool {
	// 空值节点 ：叶子节点默认是黑色(红黑树性质3)
	if node == nil {
		return BLACK
	}
	return node.color
}

// 判断节点是否是黑色
func isBlack(node *hashNode) bool {
	return colorOf(node) == BLACK
}

// 判断节点是否是红色
func isRed(node *hashNode) bool {
	return colorOf(node) == RED
}

type hashNode struct {
	hash  int  // hash值
	color bool // 默认是红色(建议新添加的节点默认为 RED，这样能够让红黑树的性质尽快满足（性质 1、2、3、5
	// 都满足，性质 4 不一定）)
	key    any
	value  any
	parent *hashNode
	left   *hashNode
	right  *hashNode

	prev *hashNode
	next *hashNode
}

func createNode(key any, value any, parent *hashNode) *hashNode {
	var hash int
	if key != nil {
		hash = utils.Hash(key)
	}
	return &hashNode{
		color:  RED,
		hash:   hash ^ (hash >> 16),
		key:    key,
		value:  value,
		parent: parent,
	}
}
func (n *hashNode) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *hashNode) hasTwoChildren() bool {
	return n.left != nil && n.right != nil
}

func (n *hashNode) isLeftChild() bool {
	return n.parent != nil && n.parent.left == n
}

func (n *hashNode) isRightChild() bool {
	return n.parent != nil && n.parent.right == n
}

// 返回兄弟节点
func (n *hashNode) sibling() *hashNode {
	if n.isLeftChild() {
		return n.parent.right
	}
	if n.isRightChild() {
		return n.parent.left
	}
	return nil
}

func (n *hashNode) String() string {
	return fmt.Sprintf("Node_%v_%v", n.key, n.value)
}
