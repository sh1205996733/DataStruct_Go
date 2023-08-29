package set

import (
	gomap "DataStruct_Go/10-映射"
	"DataStruct_Go/utils"
)

// treeMapSet treeMap集合
type treeMapSet struct {
	m utils.Map
}

func NewTreeMapSet() utils.Set {
	return &treeMapSet{
		m: gomap.NewTreeMap(),
	}
}

func (t *treeMapSet) Size() int {
	return t.m.Size()
}

func (t *treeMapSet) IsEmpty() bool {
	return t.m.IsEmpty()
}

func (t *treeMapSet) Clear() {
	t.m.Clear()
}

func (t *treeMapSet) Contains(element any) bool {
	return t.m.ContainsValue(element)
}

func (t *treeMapSet) Add(element any) {
	t.m.Put(element, nil)
}

func (t *treeMapSet) Remove(element any) {
	t.m.Remove(element)
}

func (t *treeMapSet) Traversal(visitors ...utils.Visitor) {
	t.m.Traversal(visitors...)
}
