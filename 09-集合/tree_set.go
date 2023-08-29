package set

import (
	rbtree "DataStruct_Go/08-红黑树"
	"DataStruct_Go/utils"
)

// treeSet 树形集合
type treeSet struct {
	tree *rbtree.RBTree
}

func NewTreeSet() utils.Set {
	return &treeSet{
		tree: new(rbtree.RBTree),
	}
}

func (t *treeSet) Size() int {
	return t.tree.Size
}

func (t *treeSet) IsEmpty() bool {
	return t.tree.IsEmpty()
}

func (t *treeSet) Clear() {
	t.tree.Clear()
}

func (t *treeSet) Contains(element any) bool {
	return t.tree.Contains(element)
}

func (t *treeSet) Add(element any) {
	t.tree.Add(element)
}

func (t *treeSet) Remove(element any) {
	t.tree.Remove(element)
}

func (t *treeSet) Traversal(visitors ...utils.Visitor) {
	t.tree.InorderTraversal()
}
