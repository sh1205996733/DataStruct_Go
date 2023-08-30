package trie

// 前缀树
type trie struct {
	size int
	root *trieNode
}

func NewTrie() *trie {
	return &trie{}
}

func (t *trie) Size() int {
	return t.size
}

func (t *trie) IsEmpty() bool {
	return t.size == 0
}

func (t *trie) Clear() {
	t.size = 0
	t.root = nil
}

func (t *trie) StartsWith(prefix string) bool {
	return t.node(prefix) != nil
}

func (t *trie) Get(key string) any {
	node := t.node(key)
	if node != nil && node.word {
		return node.value
	}
	return nil
}

func (t *trie) Contains(key string) bool {
	node := t.node(key)
	return node != nil && node.word
}

func (t *trie) Remove(key string) any {
	// 找到最后一个节点
	node := t.node(key)
	// 如果不是单词结尾，不用作任何处理
	if node == nil || !node.word {
		return nil
	}
	oldValue := node.value
	t.size--
	// 如果还有子节点,word置为false，value置空 然后直接返回
	if node.children != nil && len(node.children) > 0 {
		node.word = false
		node.value = nil
		return oldValue
	}
	// 如果没有子节点，从下往上遍历，看其父节点是否还有其他子节点
	for parent := node.parent; parent != nil; parent = node.parent {
		delete(parent.children, node.character)
		if parent.word || len(parent.children) > 0 { //如果parent.children仍不为空或者parent.word是一个单词就直接break
			break
		}
		node = parent
	}
	return oldValue
}

func (t *trie) Add(key string, value any) any {
	keyCheck(key)
	// 创建根节点
	if t.root == nil {
		t.root = createNode(nil)
	}
	node := t.root
	/**
	 * 遍历key的每一个字节，如果的node的children==null或者node.children.get(c)不存在 就新增一个childNode，
	 * 存在就node = childNode
	 */
	for _, k := range key {
		c := string(k)
		emptyChildren := node.children == nil
		var childNode *trieNode
		if !emptyChildren {
			childNode = node.children[c]
		}
		if childNode == nil { //node.children == null 或者node.children.get(c)不存在
			childNode = createNode(node)
			childNode.character = c
			if emptyChildren {
				node.children = make(map[string]*trieNode)
			}
			node.children[c] = childNode
		}
		node = childNode
	}
	if node.word { // 已经存在这个单词,覆盖
		old := node.value
		node.value = value
		return old
	}
	// 新增一个单词
	node.word = true
	node.value = value
	t.size++
	return nil
}

func (t *trie) node(key string) *trieNode {
	keyCheck(key)
	node := t.root
	for _, c := range key {
		if node == nil || node.children == nil || len(node.children) == 0 {
			return nil
		}
		node = node.children[string(c)]
	}
	return node
}

func keyCheck(key string) {
	if key == "" || len(key) == 0 {
		panic("key must not be empty")
	}
}

type trieNode struct {
	parent    *trieNode            // 节点的父节点
	children  map[string]*trieNode // 当前节点上所有节点
	character string               // 节点上的每个字符
	value     any                  // 如果是单词的结尾（存储一个完整的单词）
	word      bool                 // 是否为单词的结尾（是否为一个完整的单词）
}

func createNode(parent *trieNode) *trieNode {
	return &trieNode{parent: parent}
}
